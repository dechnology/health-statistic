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
      {isAuthenticated ? <RegisterView /> : <div>Log In First</div>}
    </>
  );
}

export default App;
