import { useAuth0 } from '@auth0/auth0-react';
import axios from 'axios';
import { baseUrl } from '../utils/const';

const AuthButton = () => {
  const { logout, isAuthenticated, loginWithRedirect, getAccessTokenSilently } =
    useAuth0();

  const handleDelete = async (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();

    const token = await getAccessTokenSilently();

    const response = await axios.delete(`${baseUrl}/user`, {
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    });

    console.log(response.data);
  };

  if (isAuthenticated) {
    return (
      <>
        <button
          type="button"
          className="hover:bg-sky-500 py-5 px-2 transition-all"
          onClick={() =>
            logout({ logoutParams: { returnTo: window.location.origin } })
          }
        >
          Log Out
        </button>
        <button
          type="button"
          className="hover:bg-sky-500 py-5 px-2 transition-all"
          onClick={handleDelete}
        >
          Delete User
        </button>
      </>
    );
  }

  return (
    <button
      type="button"
      className="hover:bg-sky-500 py-5 px-2 transition-all"
      onClick={() => loginWithRedirect()}
    >
      Log In
    </button>
  );
};

export default AuthButton;
