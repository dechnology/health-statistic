import NavBar from '../components/NavBar';
import { Outlet } from 'react-router-dom';

function Root() {
  return (
    <>
      <NavBar />
      <div className="bg-slate-50 mx-auto px-10 pt-10">
        <Outlet />
      </div>
    </>
  );
}

export default Root;
