import { AppointmentRow, IconWrapper, Column, Status, PhoneIcon } from './styles';

type Props = {
  patientName: string,
  insurance: string,
  procedure: string,
  technician: string,
  location: string,
  status: string,
  startDate: string,
  endDate: string,
};

function ListCard({ patientName, insurance, location, procedure, status, technician, startDate, endDate }: Props) {
  return (
    <AppointmentRow>
      <input type="checkbox" />
      <Column>{startDate} - {endDate}</Column>
      <Column bold>{patientName}</Column>
      <Column>{insurance}</Column>
      <Column>{procedure}</Column>
      <Column>{technician}</Column>
      <Column>{location}</Column>
      <Status>
        <img src="/calendar-icon.svg" alt="calendar" />
        {status}
      </Status>
      <IconWrapper>
        <PhoneIcon />
      </IconWrapper>
    </AppointmentRow>
  );
}

export default ListCard;