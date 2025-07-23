import api from './api';

type TListPatientsFilters = {
  searchTerm?: string,
  page: number,
}

export async function listPatients(filters: TListPatientsFilters): Promise<any> {
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
    console.log('PATIENTS RESPONSE', res.data)
    return res.data;
  } catch (error: any) {
    if (error.response?.status === 401) {
      throw new Error('unauthorized');
    }
    throw new Error('Login failed');
  }
}

export async function createPatient(formData: {
  patientName: string;
  phone: string;
  address?: string;
  insurance?: string;
  email?: string;
}) : Promise<any> {
  try {
    const token = localStorage.getItem('token');
    const data: any = new FormData();
    for (const key in formData) {
      data.append(key, formData[key]);
    }
    const res = await api.post('/patient/create', data, { headers: { 'Authorization': token, 'Content-Type': 'multipart/form-data' }});
    return res.data;
  } catch (error: any) {
    if (error.response?.status === 401) {
      throw new Error('unauthorized');
    }
    throw new Error(error);
  }
}

export async function editPatient(formData: {
  uuid: string;
  name: string;
  phone: string;
  email?: string;
  address?: string;
}) : Promise<any> {
  try {
    const token = localStorage.getItem('token');
    const data: any = new FormData();

    for (const key in formData) {
    console.log("DATA EDIT", key)

      data.append(key, formData[key]);
    }

    console.log("DATA EDIT", data)
    const res = await api.put(`/patient/${formData.uuid}`, data, { headers: { 'Authorization': token, 'Content-Type': 'multipart/form-data' }});
    return res.data;
  } catch (error: any) {
    if (error.response?.status === 401) {
      throw new Error('unauthorized');
    }
    throw new Error(error);
  }
}

export async function deletePatients(input: string[]) {
  try {
    const token = localStorage.getItem('token');
    if (!token) throw new Error('unauthorized');

    const endpoint = '/patient/';

    const res = await api.delete(endpoint, {
      data: {
        uuids: input
      },
      headers: { Authorization: token }
    });

    console.log("DELETE RESPONSE",res);
    return res.data;
  } catch (error: any) {
    if (error.response?.status === 401) {
      throw new Error('unauthorized');
    }
    console.error('Error deleting patients:', error);
    throw new Error('Delete failed');
  }
}