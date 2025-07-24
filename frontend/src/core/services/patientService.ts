import type { TGetPatientListResponse } from '../types/patient';
import api from './api';

type TListPatientsFilters = {
  searchTerm?: string,
  page: number,
}

type TCreatePatientData = {
  patientName: string;
  phone: string;
  address?: string;
  insurance?: string;
  email?: string;
}

type TEditPatientData = {
  uuid: string;
  name: string;
  phone: string;
  email?: string;
  address?: string;
}

export async function listPatients(filters: TListPatientsFilters): Promise<TGetPatientListResponse> {
  try {
    const token = localStorage.getItem('token');

    let endpoint = '/patient/list';
    const query = [];
    query.push(`page=${encodeURIComponent(filters.page)}`);

    if (filters.searchTerm || filters?.page ) {
      if (filters.searchTerm) query.push(`searchTerm=${encodeURIComponent(filters.searchTerm)}`);
      endpoint += `?${query.join('&')}`;
    }
    const res = await api.get(endpoint, { headers: { 'Authorization': token }});
    return res.data;
  } catch (error: any) {
    if (error.response?.status === 401) {
      throw new Error('unauthorized');
    }
    throw new Error('Login failed');
  }
}

export async function createPatient(formData: TCreatePatientData) : Promise<void> {
  try {
    const token = localStorage.getItem('token');

    const data = new FormData();

    for (const key in formData) {
      const typedKey = key as keyof TCreatePatientData;
      const value = formData[typedKey];
    
      if (value !== undefined && value !== null) {
        data.append(typedKey, String(value));
      }
    }

    await api.post('/patient/create', data, { headers: 
      { 
        'Authorization': token, 
        'Content-Type': 'multipart/form-data' 
      }
    });
    return;
  } catch (error: any) {
    if (error.response?.status === 401) {
      throw new Error('unauthorized');
    }
    throw new Error(error);
  }
}

export async function editPatient(formData: TEditPatientData) : Promise<void> {
  try {
    const token = localStorage.getItem('token');

    const data = new FormData();

    for (const key in formData) {
      const typedKey = key as keyof TEditPatientData;
      const value = formData[typedKey];
    
      if (value !== undefined && value !== null) {
        data.append(typedKey, String(value));
      }
    }

    await api.put(`/patient/${formData.uuid}`, data, { headers: { 'Authorization': token, 'Content-Type': 'multipart/form-data' }});
    return;
  } catch (error: any) {
    if (error.response?.status === 401) {
      throw new Error('unauthorized');
    }
    throw new Error(error);
  }
}

export async function deletePatients(input: string[]): Promise<void> {
  try {
    const token = localStorage.getItem('token');
    if (!token) throw new Error('unauthorized');

    const endpoint = '/patient/';

    await api.delete(endpoint, {
      data: {
        uuids: input
      },
      headers: { Authorization: token }
    });

    return;
  } catch (error: any) {
    if (error.response?.status === 401) {
      throw new Error('unauthorized');
    }
    console.error('Error deleting patients:', error);
    throw new Error('Delete failed');
  }
}