import type { TPaginationMetadata } from "./paginationMetadata";

export type TAppointmentItem = {
    uuid: string;
    end_date: string;
    location: string;
    patient: {
        name: string;
        insurance: string;
        phone: string;
    };
    procedure: string;
    start_date: string;
    status: string;
};

export type TAppointmentResponse = {
    data: TAppointmentItem[];
    metadata: TPaginationMetadata
}

