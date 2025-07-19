import api from './api';


export async function listAppointments(): Promise<any> {
  try {
    const token = localStorage.getItem('token');

    const res = await api.get('/appointment/list', { headers: { 'Authorization': token }});
    console.log('RESPONSE', res)
    return res.data;
  } catch (error: any) {
    if (error.response?.status === 401) {
      throw new Error('unauthorized');
    }
    throw new Error('Login failed');
  }
}

export async function createAppointment(formData: {
  patientName: string;
  insurance: string;
  procedure: string;
  technician: string;
  location: string;
  startDate: string;
  endDate: string;
}) : Promise<any> {
  try {
    const token = localStorage.getItem('token');

    const res = await api.post('/appointment/create', formData, { headers: { 'Authorization': token }});
    return res.data;
  } catch (error: any) {
    if (error.response?.status === 401) {
      throw new Error('unauthorized');
    }
    throw new Error('Login failed');
  }
}