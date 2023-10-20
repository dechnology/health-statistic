import axios from 'axios';
import { useState } from 'react';
import { useAuth0 } from '@auth0/auth0-react';
import { Deegoo } from '../types/deegoo';
import { baseUrl } from '../utils/const';

const DeegooForm = () => {
  const { user, getAccessTokenSilently } = useAuth0();
  const [deegooData, setDeegooData] = useState<Omit<Deegoo, 'id'>>({
    user_id: user?.sub,
    perception: 0,
    focus: 0,
    execution: 0,
    memory: 0,
    language: 0,
  });

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const token = await getAccessTokenSilently();

    const response = await axios.post(
      `${baseUrl}/deegoo`,
      { ...deegooData },
      {
        headers: {
          Authorization: `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
      },
    );

    console.log(response.data);
  };

  return (
    <form
      className="border border-solid border-slate-300 bg-white p-4 rounded-2xl flex flex-col gap-4"
      onSubmit={handleSubmit}
    >
      <h2 className="text-2xl font-semibold text-center">Deegoo Data Upload</h2>
      <div className="flex flex-col gap-2">
        <label className="flex">
          <div>Perception</div>
          <input
            type="text"
            name="stepCount"
            onChange={(e) =>
              setDeegooData((prev) => ({
                ...prev,
                perception: parseFloat(e.target.value),
              }))
            }
          />
        </label>
        <label className="flex">
          <div>Focus</div>
          <input
            type="text"
            name="stepCount"
            onChange={(e) =>
              setDeegooData((prev) => ({
                ...prev,
                focus: parseFloat(e.target.value),
              }))
            }
          />
        </label>
      </div>
      <button
        type="submit"
        className="border border-solid rounded-md p-4 bg-sky-500 text-white font-bold text-xl hover:bg-sky-600 transition-all"
      >
        Submit
      </button>
    </form>
  );
};

export default DeegooForm;
