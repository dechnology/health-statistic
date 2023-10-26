import './index.css';
import ReactDOM from 'react-dom/client';
import { Auth0Provider } from '@auth0/auth0-react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { RouterProvider } from 'react-router-dom';
import { router } from './router.tsx';

const queryClient = new QueryClient();

const domain = 'itri-dechnology.jp.auth0.com';

const clientId = 'holxN6SuSQtRV5oOSwOIXWYwJnvioObh';

const audience = 'https://health-statistic.dechnology.com.tw/api/v1/';

const scopes = [
  'profile',
  'email',
  'read:all',
  'write:all',
  'read:current_user',
  'write:current_user',
  'read:public',
];

ReactDOM.createRoot(document.getElementById('root')!).render(
  <QueryClientProvider client={queryClient}>
    <Auth0Provider
      domain={domain}
      clientId={clientId}
      authorizationParams={{
        redirect_uri: window.location.origin,
        audience: audience,
        scope: scopes.join(' '),
      }}
    >
      <RouterProvider router={router} />
    </Auth0Provider>
  </QueryClientProvider>,
);
