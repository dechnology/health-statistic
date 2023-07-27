import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import NavBar from './components/NavBar';
import RegisterView from './views/registrationView';
import { useAuth0 } from '@auth0/auth0-react';

const queryClient = new QueryClient();

function App() {
  const { isAuthenticated } = useAuth0();
  return (
    <QueryClientProvider client={queryClient}>
      <NavBar />
      {isAuthenticated ? <RegisterView /> : <div>Log In First</div>}
    </QueryClientProvider>
  );
}

export default App;
