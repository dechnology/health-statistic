import { Answer, Questionnaire } from '../types/questionnaire';

interface QuestionnaireFormProps {
  quesionnaire: Questionnaire;
  answers: Answer[];
  handleAnswersChange: (quesionIndex: number, choice_id: string) => void;
}

const QuestionnaireForm = ({
  quesionnaire,
  answers,
  handleAnswersChange,
}: QuestionnaireFormProps) => {
  const questions = quesionnaire.questions.map((question, questionIndex) => (
    <div key={question.id}>
      <h3>{question.body}</h3>
      {question.type === 'single_choice' && question.choices ? (
        <ol>
          {question.choices.map((choice) => (
            <li key={choice.id}>
              <input
                type="radio"
                name={question.id}
                checked={
                  answers[questionIndex] &&
                  answers[questionIndex].body === choice.id
                }
                onChange={() => handleAnswersChange(questionIndex, choice.id)}
              />
              <span>{choice.body}</span>
            </li>
          ))}
        </ol>
      ) : (
        <input
          type="text"
          value={answers[questionIndex].body}
          onChange={(e) => handleAnswersChange(questionIndex, e.target.value)}
        />
      )}
    </div>
  ));

  return (
    <div>
      <h2 className="text-xl">{quesionnaire?.name}</h2>
      {questions}
    </div>
  );
};

export default QuestionnaireForm;
