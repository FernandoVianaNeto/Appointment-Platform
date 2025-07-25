import type { AxiosError } from 'axios';
import type { TAppointmentResponse } from '../types/appointments';
import api from './api';

type TListAppointmentsFilters = {
  page: string|number,
  searchTerm?: string, 
  filterType?: string,
  date?: string,
};

type TCreateAppointmentFormData = {
  patientName: string;
  insurance: string;
  procedure: string;
  technician: string;
  location: string;
  start_date: string;
  end_date: string;
};

type TEditAppointmentFormData = {
  uuid: string;
  patientName: string;
  insurance: string;
  procedure: string;
  technician: string;
  location: string;
  start_date: string;
  end_date: string;
};

export async function listAppointments(input: TListAppointmentsFilters): Promise<TAppointmentResponse> {
  try {
    const token = localStorage.getItem('token');

    let endpoint = '/appointment/list';

    const query = [];
    query.push(`page=${encodeURIComponent(input.page)}`);

    if (input.searchTerm || input.filterType || input.date || input?.page ) {
      if (input.searchTerm) query.push(`searchTerm=${encodeURIComponent(input.searchTerm)}`);
      if (input.filterType) query.push(`filterType=${encodeURIComponent(input.filterType)}`);
      if (input.date) query.push(`date=${encodeURIComponent(input.date)}`);
      endpoint += `?${query.join('&')}`;
    }

    const res = await api.get(endpoint, { headers: { 'Authorization': token }});

    return res.data;
  } catch (error) {
    const axiosError = error as AxiosError;
        
    if (axiosError.response?.status === 401) {
        throw new Error('Invalid email or password');
    }
    
    throw new Error('Could not list appointments by filters');
  }
}

export async function createAppointment(formData: TCreateAppointmentFormData) : Promise<void> {
  try {
    const token = localStorage.getItem('token');

    const data = new FormData();

    for (const key in formData) {
      const typedKey = key as keyof TCreateAppointmentFormData;
      const value = formData[typedKey];

      if (value !== undefined && value !== null) {
          data.append(typedKey, String(value));
      }
    }

    await api.post('/appointment/create', data, { headers: { 
      'Authorization': token, 
      'Content-Type': 'multipart/form-data' 
    }});

    return;
  } catch (error) {
    const axiosError = error as AxiosError;
        
    if (axiosError.response?.status === 401) {
        throw new Error('Invalid email or password');
    }

    throw new Error('Could not create an appointment');
  }
}

export async function editAppointment(formData: TEditAppointmentFormData) : Promise<any> {
  try {
    const token = localStorage.getItem('token');
    if (!token) throw new Error('unauthorized');

    const data = new FormData();
    for (const key in formData) {
      const typedKey = key as keyof TEditAppointmentFormData;
        const value = formData[typedKey];

        if (value !== undefined && value !== null) {
            data.append(typedKey, String(value));
        }
    }
    
    const res = await api.put(`/appointment/${formData.uuid}`, data, { headers: { 'Authorization': token, 'Content-Type': 'multipart/form-data' }});
    return res.data;
  } catch (error) {
    const axiosError = error as AxiosError;
        
    if (axiosError.response?.status === 401) {
        throw new Error('Invalid email or password');
    }
    throw new Error("Error on edit appointment");
  }
}

export async function deleteAppointments(input: string[]) {
  try {
    const token = localStorage.getItem('token');
    if (!token) throw new Error('unauthorized');

    const endpoint = '/appointment/';

    const res = await api.delete(endpoint, {
      data: {
        uuids: input
      },
      headers: { Authorization: token }
    });

    return res.data;
  } catch (error: any) {
    if (error.response?.status === 401) {
      throw new Error('unauthorized');
    }
    throw new Error('Delete appointment failed');
  }
}


export async function updateAppointmentStatus(uuid: string, status: string) {
  try {
    const endpoint = `/appointment/update-status?uuid=${uuid}&status=${status}`;

    const res = await api.post(endpoint);

    return res.data;
  } catch (error: any) {
    if (error.response?.status === 401) {
      throw new Error('unauthorized');
    }
    throw new Error('Delete appointment failed');
  }
}