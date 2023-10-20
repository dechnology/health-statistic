import axios from 'axios';
import { useState } from 'react';
import { useAuth0 } from '@auth0/auth0-react';
import { HealthKit } from '../types/healthKit';
import { baseUrl } from '../utils/const';

const HealthKitForm = () => {
  const { getAccessTokenSilently } = useAuth0();
  const [healthKitData, setHealthKitData] = useState<Omit<HealthKit, 'id'>>();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setHealthKitData({
      startDate: new Date(),
      endDate: new Date(),
      stepCount: parseFloat(e.target.value),
    });
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const token = await getAccessTokenSilently();

    const response = await axios.post(
      `${baseUrl}/user/healthkit`,
      {
        start_date: healthKitData?.startDate,
        end_Date: healthKitData?.endDate,
        step_count: healthKitData?.stepCount,
      },
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
      <h2 className="text-2xl font-semibold text-center">
        HealthKit Data Upload
      </h2>
      <div className="flex flex-col gap-2">
        <label>
          <input type="text" name="stepCount" onChange={handleChange} />
          <span>Step Count</span>
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

export default HealthKitForm;
