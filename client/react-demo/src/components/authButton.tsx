import { useAuth0 } from '@auth0/auth0-react';

const AuthButton = () => {
  const { logout, isAuthenticated, loginWithRedirect } = useAuth0();

  if (isAuthenticated) {
    return (
      <button
        type="button"
        className="hover:bg-sky-500 py-5 px-2 transition-all"
        onClick={() =>
          logout({ logoutParams: { returnTo: window.location.origin } })
        }
      >
        Log Out
      </button>
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
