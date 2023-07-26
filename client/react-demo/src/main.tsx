import ReactDOM from 'react-dom/client';
import { Auth0Provider } from '@auth0/auth0-react';
import App from './App.tsx';
import './index.css';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <Auth0Provider
    domain="itri-dechnology.jp.auth0.com"
    clientId="holxN6SuSQtRV5oOSwOIXWYwJnvioObh"
    authorizationParams={{
      redirect_uri: window.location.origin,
    }}
  >
    <App />
  </Auth0Provider>,
);
