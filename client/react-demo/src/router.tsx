// create and export a react-router Routes in object

import { createBrowserRouter } from 'react-router-dom';
import RegisterView from './views/registrationView';
import ProfileView from './views/profileView';

export const router = createBrowserRouter([
  {
    path: '/',
    element: <RegisterView />,
  },
  {
    path: '/profile',
    element: <ProfileView />,
  },
]);
