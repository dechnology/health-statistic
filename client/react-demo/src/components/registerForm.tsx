import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { useEffect, useRef } from 'react';

interface Choice {
  id: string;
  body: string;
}

interface Question {
  id: string;
  body: string;
  type: 'short_answer' | 'single_choice' | 'multiple_choice';
  choices: Choice[];
}

interface Answer {
  question_id: string;
  choice_ids?: string[];
  body?: string;
}

interface Questionnaire {
  id: string;
  name: string;
  created_at: string;
  questions: Question[];
}

const baseUrl = 'https://health-statistic.dechnology.com.tw/api/v1';

const url = `${baseUrl}/questionnaires/registration`;

const fetchRegistrationQuestionnaire = async (): Promise<Questionnaire> => {
  const res = await axios.get(url);
  return res.data;
};

const RegisterForm = () => {
  const { isLoading, error, data } = useQuery<Questionnaire, Error>(
    ['registrationQuestionnaireData'],
    fetchRegistrationQuestionnaire,
  );
  const answers = useRef<Answer[]>();

  useEffect(() => {
    answers.current = data?.questions.map((question) => {
      const answer = { question_id: question.id };
      return question.type === 'short_answer'
        ? { body: '', ...answer }
        : { choice_ids: [question.choices[0].id], ...answer };
    });
  }, [data]);

  if (isLoading) return <div>Loading Registration Questionnaire...</div>;
  if (error) return <div>Error: {error.message}</div>;

  const questions = data.questions.map((question, idx) => (
    <>
      <h3 key={question.id}>{question.body}</h3>
      <ol>
        {question.choices.map((choice) => (
          <li key={choice.id}>
            <input
              type="radio"
              name={question.id}
              checked={
                answers.current &&
                answers.current[idx] &&
                answers.current[idx].choice_ids?.includes(choice.id)
              }
              onChange={() => {
                if (answers.current && answers.current[idx]) {
                  answers.current[idx].choice_ids = [choice.id];
                }
              }}
            />
            <span>{choice.body}</span>
          </li>
        ))}
      </ol>
    </>
  ));

  return (
    <div>
      <h2 className="text-xl">{data?.name}</h2>
      <div>{questions}</div>
      <button type="submit">Submit</button>
    </div>
  );
};

export default RegisterForm;
