import api from './api';

export async function login(email: string, password: string) {
  try {
    const res = await api.post('/auth/login', { email, password });

    const { token, user } = res.data;

    localStorage.setItem('token', token);

    localStorage.setItem('user', JSON.stringify(user));

    return res.data;
  } catch (error: any) {
    if (error.response?.status === 401) {
      throw new Error('Invalid email or password');
    }
    throw new Error('Login failed');
  }
}

export function logout() {
  localStorage.removeItem('token');
  localStorage.removeItem('user');
  window.location.href = '/login';
}