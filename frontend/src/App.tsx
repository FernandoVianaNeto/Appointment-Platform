import { Routes, Route } from 'react-router-dom';
import Login from './pages/Login/index';
import Appointments from './pages/Appointments';

export default function App() {
  return (
    <Routes>
      <Route path="/login" element={<Login />} />
      <Route path="/appointments" element={<Appointments />} />
    </Routes>
  );
}