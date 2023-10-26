import axios from 'axios';
import { baseUrl } from '../utils/const';
import { useQuery } from '@tanstack/react-query';
import { useEffect, useRef } from 'react';
import { useAuth0 } from '@auth0/auth0-react';
import HealthKitForm from '../components/healthKitForm';
import DeegooForm from '../components/deegooForm';

const getRegistrationQuestionnaire = async (
  token: string | null,
): Promise<Record<string, string>> => {
  const res = await axios.get(`${baseUrl}/user`, {
    headers: { Authorization: `Bearer ${token}` },
  });
  console.log(res.data);
  return res.data;
};

const ProfileView = () => {
  const { getAccessTokenSilently } = useAuth0();
  const token = useRef<string | null>(null);

  const { data: userData } = useQuery<Record<string, string>, Error>({
    queryKey: ['registrationQuestionnaire'],
    queryFn: async () => getRegistrationQuestionnaire(token.current),
    enabled: token !== null,
  });

  useEffect(() => {
    const getToken = async () => {
      token.current = await getAccessTokenSilently();
    };
    getToken();
    console.log(userData);
  }, []);

  return (
    <>
      <div className="flex flex-col gap-4">
        <h2 className="text-3xl font-bold text-center">Profile</h2>
        <HealthKitForm />
        <DeegooForm />
      </div>
    </>
  );
};

export default ProfileView;
