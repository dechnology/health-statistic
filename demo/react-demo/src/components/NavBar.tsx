import { useAuth0 } from '@auth0/auth0-react';
import AuthButton from './authButton';

const NavBar = () => {
  const { user, isAuthenticated } = useAuth0();

  return (
    <>
      <nav className="flex justify-between">
        <div className="flex">
          {isAuthenticated ? <p>{user && user.name}</p> : 'Not Logged In'}
        </div>
        <AuthButton />
      </nav>
    </>
  );
};

export default NavBar;
