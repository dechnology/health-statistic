import axios from 'axios';
import { useState } from 'react';
import { useAuth0 } from '@auth0/auth0-react';
import { baseUrl } from '../utils/const';

const FcmForm = () => {
  const { getAccessTokenSilently } = useAuth0();
  const [fcmToken, setFcmToken] = useState<string>('');

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const token = await getAccessTokenSilently();

    const response = await axios.put(
      `${baseUrl}/user/fcm`,
      {
        fcm_token: fcmToken,
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
      <h2 className="text-2xl font-semibold text-center">FCM Token Update</h2>
      <label className="flex flex-col gap-2">
        <div className="uppercase font-medium">Token</div>
        <textarea
          name="stepCount"
          value={fcmToken}
          className="border resize-none h-fit w-full border-black border-opacity-40 p-2 rounded "
          onChange={(e) => setFcmToken(e.target.value)}
        ></textarea>
      </label>
      <button
        type="submit"
        className="border border-solid rounded-md py-4 px-8 mx-auto bg-sky-500 text-white font-bold text-xl hover:bg-sky-600 transition-all"
      >
        Update Token
      </button>
    </form>
  );
};

export default FcmForm;
