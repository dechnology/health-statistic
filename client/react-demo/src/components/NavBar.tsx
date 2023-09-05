import { useAuth0 } from '@auth0/auth0-react';
import AuthButton from './authButton';

const NavBar = () => {
  const { user, isAuthenticated } = useAuth0();

  return (
    <>
      <nav className="flex justify-between bg-sky-300 px-10 font-semibold text-lg">
        <div className="flex py-5">
          {isAuthenticated ? <p>{user && user.name}</p> : 'Not Logged In'}
        </div>
        <AuthButton />
      </nav>
    </>
  );
};

export default NavBar;
