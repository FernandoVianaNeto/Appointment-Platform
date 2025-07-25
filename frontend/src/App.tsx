import { Routes, Route } from 'react-router-dom';
import Login from './pages/Login/index';
import Appointments from './pages/Appointments';
import Patients from './pages/Patients';
import ForgotPassword from './pages/ForgotPassword';

export default function App() {
  return (
    <Routes>
      <Route path="/login" element={<Login />} />
      <Route path="/appointments" element={<Appointments />} />
      <Route path="/patients" element={<Patients />} />
      <Route path="/forgot-password" element={<ForgotPassword />} />
    </Routes>
  );
}