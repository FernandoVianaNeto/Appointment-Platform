import api from './api';


export async function listAppointments(): Promise<any> {
  try {
    const res = await api.get('/appointment/list');

    const { token, user } = res.data;

    localStorage.setItem('token', token);

    localStorage.setItem('user', JSON.stringify(user));

    return res.data;
  } catch (error: any) {
    if (error.response?.status === 401) {
      throw new Error('unauthorized');
    }
    throw new Error('Login failed');
  }
}