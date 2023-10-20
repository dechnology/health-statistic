import { useEffect, useState } from 'react';
import NavBar from './components/NavBar';
import { RouterProvider } from 'react-router-dom';
import { router } from './router';

function App() {
  return (
    <>
      <NavBar />
      <div className=" bg-slate-50">
        <RouterProvider router={router} />
      </div>
    </>
  );
}

export default App;
