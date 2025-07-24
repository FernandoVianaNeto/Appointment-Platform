import type { AxiosError } from "axios";
import api from "./api";

type TCreateClinicFormData = {
    email: string;
    password: string;
    name: string;
}

export async function createClinic(formData: TCreateClinicFormData): Promise<void> {
    try {
        const data: any = new FormData();
        for (const key in formData) {
            data.append(key, formData[key]);
        }

        await api.post('/user/create', data, { headers: { 'Content-Type': 'multipart/form-data' }});
    
        return;
    } catch (error) {
        console.log(error)
        const axiosError = error as AxiosError;
    
        if (axiosError.response?.status === 401) {
            throw new Error('Invalid email or password');
        }
    
        throw new Error('An unexpected error occurred during login');
    }
  }