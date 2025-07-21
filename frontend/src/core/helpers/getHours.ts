export const getHours = (stringDate: string) => {
    const date = new Date(stringDate);

    const hours = date.getHours().toString().padStart(2, '0');
    const minutes = date.getMinutes().toString().padStart(2, '0');

    return `${hours}:${minutes}`
}