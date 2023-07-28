import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { FormEvent, useEffect, useRef, useState } from 'react';
import QuestionnaireForm from '../components/questionnaireForm';
import { Answer, Questionnaire } from '../types/questionnaire';
import InfoQuestionnaire from '../utils/infoQuestionnaire';
import { useAuth0 } from '@auth0/auth0-react';
import * as jose from 'jose';

const baseUrl = 'https://health-statistic.dechnology.com.tw/api/v1';

const fetchRegistrationQuestionnaire = async (): Promise<Questionnaire> => {
  const res = await axios.get(`${baseUrl}/questionnaires/registration`);
  return res.data;
};

const RegistrationView = () => {
  const { user, isAuthenticated, getAccessTokenSilently } = useAuth0();
  const token = useRef<string | null>(null);

  const {
    isLoading,
    error,
    data: registrationQuestionnaire,
  } = useQuery<Questionnaire, Error>(
    ['registrationQuestionnaire'],
    fetchRegistrationQuestionnaire,
  );

  const [infoAnswers, setInfoAnswers] = useState<Answer[]>(
    InfoQuestionnaire.questions.map((q) => ({
      question_id: q.id,
      type: q.type,
      body: q.type === 'single_choice' && q.choices ? q.choices[0].id : '',
    })),
  );
  const [registrationAnswers, setRegistrationAnswers] = useState<
    Answer[] | undefined
  >([]);
  const requestBody = useRef<any>();

  useEffect(() => {
    const getToken = async () => {
      token.current = await getAccessTokenSilently();
      console.log('token', token.current);
      console.log('header', jose.decodeProtectedHeader(token.current));
      console.log('claims', jose.decodeJwt(token.current));
    };

    getToken();

    if (!registrationQuestionnaire) return;

    const newRegitrationAnswers = registrationQuestionnaire.questions.map(
      (q) => ({
        question_id: q.id,
        type: q.type,
        body: q.type === 'single_choice' && q.choices ? q.choices[0].id : '',
      }),
    );

    setRegistrationAnswers(newRegitrationAnswers);
  }, [registrationQuestionnaire]);

  const handleInfoAnswersChange = (
    questionIndex: number,
    choice_id: string,
  ) => {
    const newAnswers = infoAnswers?.map((answer, answerIndex) =>
      questionIndex === answerIndex ? { ...answer, body: choice_id } : answer,
    );
    setInfoAnswers(newAnswers);
  };

  const handleRegistrationAnswersChange = (
    questionIndex: number,
    choice_id: string,
  ) => {
    const newAnswers = registrationAnswers?.map((answer, answerIndex) =>
      questionIndex === answerIndex ? { ...answer, body: choice_id } : answer,
    );
    setRegistrationAnswers(newAnswers);
  };

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();

    if (
      !(
        isAuthenticated &&
        user &&
        registrationQuestionnaire &&
        registrationAnswers
      )
    )
      return;

    const token = await getAccessTokenSilently();

    const userInfo = {
      id: user.sub,
      ...infoAnswers.reduce((acc, answer) => {
        let value: string | number | boolean = answer.body;

        if (answer.question_id === 'birth_year') {
          value = parseInt(value) + 1911;
        } else if (
          answer.question_id === 'weight' ||
          answer.question_id === 'height'
        ) {
          value = parseFloat(value);
        } else if (
          answer.question_id === 'demented_among_direct_relatives' ||
          answer.question_id === 'head_injury_experience'
        ) {
          value = value === 'true';
        }
        return { ...acc, [answer.question_id]: value };
      }, {}),
    };

    const response = {
      questionnaire_id: registrationQuestionnaire.id,
      answers: registrationAnswers.map((a) =>
        a.type === 'single_choice'
          ? { question_id: a.question_id, choice_ids: [a.body] }
          : { question_id: a.question_id, body: a.body },
      ),
    };

    requestBody.current = {
      response: response,
      user: userInfo,
    };

    console.dir(requestBody.current);

    await axios.post(`${baseUrl}/register`, requestBody.current, {
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    });
  };

  if (isLoading || !registrationAnswers)
    return <div>Loading Registration Questionnaire...</div>;
  if (error) return <div>Error: {error.message}</div>;

  return (
    <form onSubmit={handleSubmit}>
      <h1 className="text-xl">Registration</h1>
      <QuestionnaireForm
        quesionnaire={InfoQuestionnaire}
        answers={infoAnswers}
        handleAnswersChange={handleInfoAnswersChange}
      />
      <QuestionnaireForm
        quesionnaire={registrationQuestionnaire}
        answers={registrationAnswers}
        handleAnswersChange={handleRegistrationAnswersChange}
      />
      <button type="submit">Submit</button>
    </form>
  );
};

export default RegistrationView;
