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
    <div key={question.id} className="flex flex-col text-xl gap-2">
      <h3 className="font-medium text-xl">{question.body}</h3>
      {question.type === 'single_choice' && question.choices ? (
        <ol className="flex flex-col text-xl gap-2">
          {question.choices.map((choice) => (
            <li key={choice.id}>
              <label
                className={`p-2 ${
                  answers[questionIndex] &&
                  answers[questionIndex].body === choice.id
                    ? 'bg-slate-300'
                    : 'hover:bg-slate-200 transition-all'
                }`}
              >
                <input
                  type="radio"
                  hidden
                  name={question.id}
                  checked={
                    answers[questionIndex] &&
                    answers[questionIndex].body === choice.id
                  }
                  onChange={() => handleAnswersChange(questionIndex, choice.id)}
                />
                <span>{choice.body}</span>
              </label>
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
    <div className="border border-solid border-slate-300 bg-white p-4 rounded-2xl flex flex-col gap-4">
      <h2 className="text-2xl font-semibold text-center">
        {quesionnaire?.name}
      </h2>
      <div className="flex flex-col gap-2">{questions}</div>
    </div>
  );
};

export default QuestionnaireForm;
