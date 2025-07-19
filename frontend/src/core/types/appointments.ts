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

export type TAppointmentMetadata = {
    currentPage: number;
    next: number;
    total: number;
    totalItems: number;
}

export type TAppointmentResponse = {
    data: TAppointmentItem[];
    metadata: TAppointmentMetadata
}

