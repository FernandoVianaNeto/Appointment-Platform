import api from './api';

interface ListAppointmentsFilters {
  searchTerm?: string, 
  filterType?: string,
}

export async function listAppointments(input?: ListAppointmentsFilters): Promise<any> {
  try {
    const token = localStorage.getItem('token');

    let endpoint = '/appointment/list';

    if (input?.searchTerm || input?.filterType) {
      const query = [];
    
      if (input.searchTerm) query.push(`searchTerm=${encodeURIComponent(input.searchTerm)}`);
      if (input.filterType) query.push(`filterType=${encodeURIComponent(input.filterType)}`);
    
      endpoint += `?${query.join('&')}`;
    }

    console.log(input, endpoint)

    const res = await api.get(endpoint, { headers: { 'Authorization': token }});
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
  start_date: string;
  end_date: string;
}) : Promise<any> {
  try {
    const token = localStorage.getItem('token');
    const data: any = new FormData();
    for (const key in formData) {
      data.append(key, formData[key]);
    }
    const res = await api.post('/appointment/create', data, { headers: { 'Authorization': token, 'Content-Type': 'multipart/form-data' }});
    return res.data;
  } catch (error: any) {
    if (error.response?.status === 401) {
      throw new Error('unauthorized');
    }
    throw new Error(error);
  }
}