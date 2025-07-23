import type { TPaginationMetadata } from "./paginationMetadata"

export type TPatientData = {
    uuid: string,
    name: string,
    insurance: string,
    phone: string,
    email: string,
    address: string,
}

export type TGetPatientListResponse = {
    data: TPatientData[],
    metadata: TPaginationMetadata,
}