export function addMinutesToTime(time: string, minutesToAdd: number): string {
    const [hours, minutes] = time.split(':').map(Number);
    const date = new Date();
  
    date.setHours(hours);
    date.setMinutes(minutes + minutesToAdd);
  
    const newHours = String(date.getHours()).padStart(2, '0');
    const newMinutes = String(date.getMinutes()).padStart(2, '0');
  
    return `${newHours}:${newMinutes}`;
}