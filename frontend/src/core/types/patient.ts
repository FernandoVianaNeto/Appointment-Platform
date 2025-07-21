import type { TPaginationMetadata } from "./paginationMetadata"

export type TPatientData = {
    uuid: string,
    name: string,
    insurance: string,
    phone: string
}

export type TGetPatientListResponse = {
    data: TPatientData[],
    metadata: TPaginationMetadata,
}