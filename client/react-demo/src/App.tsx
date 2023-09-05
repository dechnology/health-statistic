import { useEffect } from 'react';
import NavBar from './components/NavBar';
import RegisterView from './views/registrationView';
import { useAuth0 } from '@auth0/auth0-react';

function App() {
  const { isAuthenticated } = useAuth0();

  useEffect(() => {}, []);

  return (
    <>
      <NavBar />
      <div className=" bg-slate-50">
        {isAuthenticated ? <RegisterView /> : <div>Log In First</div>}
      </div>
    </>
  );
}

export default App;
