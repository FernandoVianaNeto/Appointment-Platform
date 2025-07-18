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