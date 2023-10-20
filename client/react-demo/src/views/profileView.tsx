import axios from 'axios';
import { baseUrl } from '../utils/const';
import { useQuery } from '@tanstack/react-query';
import { useEffect, useRef } from 'react';
import { useAuth0 } from '@auth0/auth0-react';
import HealthKitForm from '../components/healthkitForm';

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
  const { user, isAuthenticated, getAccessTokenSilently } = useAuth0();
  const token = useRef<string | null>(null);

  const {
    isLoading,
    error,
    data: userData,
  } = useQuery<Record<string, string>, Error>({
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
  }, [userData, token]);

  return (
    <>
      <div className="min-w-[1080px] px-10 mx-auto">
        <h2 className="text-3xl font-bold text-center">Profile</h2>
        <HealthKitForm />
      </div>
    </>
  );
};

export default ProfileView;
