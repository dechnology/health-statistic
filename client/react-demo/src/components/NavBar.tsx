import { NavLink } from 'react-router-dom';
import { useAuth0 } from '@auth0/auth0-react';
import AuthButton from './authButton';

const NavBar = () => {
  const { user, isAuthenticated } = useAuth0();

  return (
    <>
      <div className="flex justify-between bg-sky-300 px-10 font-semibold text-lg">
        <div className="flex py-5">
          {isAuthenticated ? <p>{user && user.name}</p> : 'Not Logged In'}
          <nav>
            {/* <ul className="flex">
              <li className="mx-5">
                <NavLink to="/">Registration</NavLink>
              </li>
              <li className="mx-5">
                <NavLink to="/profile">Profile</NavLink>
              </li>
            </ul> */}
          </nav>
        </div>
        <div className="flex items-center gap-2">
          {isAuthenticated ? <p>Hi! {user && user.name}.</p> : 'Not Logged In'}
          <AuthButton />
        </div>
      </div>
    </>
  );
};

export default NavBar;
