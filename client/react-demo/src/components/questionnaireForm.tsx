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
      <h3 className="font-medium text-2xl py-2">{question.body}</h3>
      {question.type === 'single_choice' && question.choices ? (
        <ol className="flex text-xl gap-2">
          {question.choices.map((choice) => (
            <li
              key={choice.id}
              className={`p-2 rounded border cursor-pointer border-blue-600 ${
                answers[questionIndex] &&
                answers[questionIndex].body === choice.id
                  ? 'bg-blue-300'
                  : 'hover:bg-blue-200 hover:text-gray-800 transition-all text-gray-500 '
              }`}
            >
              <label>
                <input
                  type="radio"
                  hidden
                  name={question.id}
                  checked={
                    answers[questionIndex]
                      ? answers[questionIndex].body === choice.id
                      : false
                  }
                  onChange={() => handleAnswersChange(questionIndex, choice.id)}
                />
                <span>{choice.body}</span>
              </label>
            </li>
          ))}
        </ol>
      ) : (
        <label>
          <input
            type="text"
            placeholder="Enter your answer here"
            name={question.id}
            value={answers[questionIndex].body}
            className="p-2 w-full border border-opacity-40 border-black rounded"
            onChange={(e) => handleAnswersChange(questionIndex, e.target.value)}
          />
        </label>
      )}
    </div>
  ));

  return (
    <div className="border border-solid border-slate-300 bg-white p-4 rounded-2xl flex flex-col gap-4">
      <h2 className="text-2xl font-semibold text-center">
        {quesionnaire?.name}
      </h2>
      <div className="flex flex-col gap-4">{questions}</div>
    </div>
  );
};

export default QuestionnaireForm;
