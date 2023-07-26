import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import './App.css';
import NavBar from './components/NavBar';
import RegisterForm from './components/registerForm';

const queryClient = new QueryClient();

function App() {
  return (
    <>
      <QueryClientProvider client={queryClient}>
        <NavBar />
        <RegisterForm />
      </QueryClientProvider>
    </>
  );
}

export default App;
