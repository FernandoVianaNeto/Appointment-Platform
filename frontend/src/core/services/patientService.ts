import api from './api';


export async function listPatients(name?: string): Promise<any> {
  try {
    const token = localStorage.getItem('token');

    let endpoint = '/patient/list'

    if (name != undefined) {
        endpoint = `/patient/list?name=${name}`
    }

    const res = await api.get(endpoint, { headers: { 'Authorization': token }});
    console.log('RESPONSE', res)
    return res.data;
  } catch (error: any) {
    if (error.response?.status === 401) {
      throw new Error('unauthorized');
    }
    throw new Error('Login failed');
  }
}

// export async function createPatient(formData: {
//   patientName: string;
//   insurance: string;
//   procedure: string;
//   technician: string;
//   location: string;
//   start_date: string;
//   end_date: string;
// }) : Promise<any> {
//   try {
//     const token = localStorage.getItem('token');
//     const data: any = new FormData();
//     for (const key in formData) {
//       data.append(key, formData[key]);
//     }
//     const res = await api.post('/appointment/create', data, { headers: { 'Authorization': token, 'Content-Type': 'multipart/form-data' }});
//     return res.data;
//   } catch (error: any) {
//     if (error.response?.status === 401) {
//       throw new Error('unauthorized');
//     }
//     throw new Error(error);
//   }
// }