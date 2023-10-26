import { NavLink } from 'react-router-dom';
import { useAuth0 } from '@auth0/auth0-react';
import AuthButton from './authButton';

const NavBar = () => {
  const { user, isAuthenticated } = useAuth0();

  return (
    <>
      <div className="flex justify-between bg-sky-300 px-10 font-semibold text-lg">
        <nav>
          <ul className="flex gap-2">
            <li className="py-5 px-2 hover:bg-sky-500 transition-all">
              <NavLink to="/">Registration</NavLink>
            </li>
            <li className="py-5 px-2 hover:bg-sky-500 transition-all">
              <NavLink to="/profile">Profile</NavLink>
            </li>
          </ul>
        </nav>
        <div className="flex items-center gap-2">
          <p className="italic">
            {isAuthenticated ? `Hi! ${user && user.name}` : 'Not Logged In'}
          </p>
          <AuthButton />
        </div>
      </div>
    </>
  );
};

export default NavBar;
