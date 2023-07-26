import { useQuery } from '@tanstack/react-query';
import axios from 'axios';

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

interface Questionnaire {
  id: string;
  name: string;
  created_at: string;
  questions: Question[];
}

const url =
  'https://health-statistic.dechnology.com.tw/api/v1/questionnaires/registration';

const fetchRegistrationQuestionnaire = async (): Promise<Questionnaire> => {
  console.log(`fetching ${url}`);
  const res = await axios.get(url, { mode: 'no-cors' });
  return res.data;
};

const RegisterForm = () => {
  const { isLoading, error, data } = useQuery<Questionnaire, Error>(
    ['registrationQuestionnaireData'],
    fetchRegistrationQuestionnaire,
  );

  if (isLoading) return <div>Loading Registration Questionnaire...</div>;
  if (error) return <div>Error: {error.message}</div>;

  const questions = data?.questions.map((question, idx) => (
    <>
      <h3 key={idx}>{question.body}</h3>
      <ol>
        {question.choices.map((choice, idx) => (
          <li key={idx}>
            <p>{choice.body}</p>
          </li>
        ))}
      </ol>
    </>
  ));

  return (
    <div>
      <h2>{data?.name}</h2>
      <div>{questions}</div>
    </div>
  );
};

export default RegisterForm;
