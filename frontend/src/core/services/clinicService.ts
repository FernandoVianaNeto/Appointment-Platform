import type { AxiosError } from "axios";
import api from "./api";

type TCreateClinicFormData = {
    email: string;
    password: string;
    name: string;
}

export async function createClinic(formData: TCreateClinicFormData): Promise<void> {
    try {
        const data = new FormData();
        for (const key in formData) {
            const typedKey = key as keyof TCreateClinicFormData;
            const value = formData[typedKey];
    
            if (value !== undefined && value !== null) {
                data.append(typedKey, String(value));
            }
        }

        await api.post('/user/create', data, { headers: { 'Content-Type': 'multipart/form-data' }});
    
        return;
    } catch (error) {
        const axiosError = error as AxiosError;
    
        if (axiosError.response?.status === 401) {
            throw new Error('Invalid email or password');
        }
    
        throw new Error('An unexpected error occurred during login');
    }
  }