import axios from 'axios';
import { useAuth0 } from '@auth0/auth0-react';
import { HealthKit } from '../types/healthKit';
import { baseUrl, sampleHKDataTuples } from '../utils/const';

const HKDataForm = () => {
  const { getAccessTokenSilently } = useAuth0();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const token = await getAccessTokenSilently();

    const healthkit: Omit<HealthKit, 'id' | 'data'> & {
      data: string[][];
    } = {
      startTime: new Date(),
      endTime: new Date(),
      data: sampleHKDataTuples,
    };

    const response = await axios.post(`${baseUrl}/user/healthkit`, healthkit, {
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    });

    console.log(response.data);
  };

  return (
    <form
      className="border border-solid border-slate-300 bg-white p-4 rounded-2xl flex flex-col gap-4"
      onSubmit={handleSubmit}
    >
      <h2 className="text-2xl font-semibold text-center">
        HealthKit Data Upload
      </h2>
      <button
        type="submit"
        className="border border-solid rounded-md py-4 px-8 mx-auto bg-sky-500 text-white font-bold text-xl hover:bg-sky-600 transition-all"
      >
        Submit HealthKit
      </button>
    </form>
  );
};

export default HKDataForm;
