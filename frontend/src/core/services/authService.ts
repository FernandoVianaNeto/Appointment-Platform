import api from './api';
import type { AxiosError } from 'axios';

type TUser = {
  id: string;
  name: string;
  email: string;
};

type TLoginResponse = {
  token: string;
  user: TUser;
};

export async function login(email: string, password: string): Promise<void> {
  try {
    const response = await api.post<TLoginResponse>('/auth/login', { email, password });

    const { token, user } = response.data;

    localStorage.setItem('token', token);
    localStorage.setItem('user', JSON.stringify(user));

    return;
  } catch (error) {
    const axiosError = error as AxiosError;

    if (axiosError.response?.status === 401) {
      throw new Error('Invalid email or password');
    }

    throw new Error('An unexpected error occurred during login');
  }
}

export function logout() {
  localStorage.removeItem('token');
  localStorage.removeItem('user');
  window.location.href = '/login';
}

export async function generateResetPasswordCode(email: string): Promise<void> {
  try {
    await api.post('/auth/generate-reset-code', { email });

    return;
  } catch (error) {
    throw new Error('An unexpected error occurred during generating the reset password code');
  }
}

export async function validateResetPasswordCode(email: string, code: number): Promise<void> {
  try {
    await api.post('/auth/validate-code', { email, code });
    return;
  } catch (error) {
    throw new Error('An unexpected error occurred during validation of reset password code');
  }
}

export async function resetPasswordCall(email: string, code: number, newPassword: string): Promise<void> {
  try {
    await api.post('/auth/reset-password', { email, code, new_password: newPassword });
    return;
  } catch (error) {
    throw new Error('An unexpected error occurred during validation of reset password code');
  }
}