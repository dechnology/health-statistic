// create and export a react-router Routes in object

import { createBrowserRouter } from 'react-router-dom';
import RegisterView from './views/registrationView';
import ProfileView from './views/profileView';
import Root from './routes/root';

export const router = createBrowserRouter([
  {
    path: '/',
    element: <Root />,
    children: [
      {
        path: '/',
        element: <RegisterView />,
      },
      {
        path: '/profile',
        element: <ProfileView />,
      },
    ],
  },
]);
