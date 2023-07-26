import { useAuth0 } from '@auth0/auth0-react';

const AuthButton = () => {
  const { logout, isAuthenticated, loginWithRedirect } = useAuth0();

  if (isAuthenticated) {
    return (
      <button
        onClick={() =>
          logout({ logoutParams: { returnTo: window.location.origin } })
        }
      >
        Log Out
      </button>
    );
  }

  return <button onClick={() => loginWithRedirect()}>Log In</button>;
};

export default AuthButton;
