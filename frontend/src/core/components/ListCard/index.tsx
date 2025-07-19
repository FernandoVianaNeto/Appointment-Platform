import { AppointmentRow, IconWrapper, Column, Status, PhoneIcon } from './styles';
import { CiCalendar } from "react-icons/ci";
import { MdOutlineCancel } from "react-icons/md";
import { GiConfirmed } from "react-icons/gi";
import { useState } from 'react';

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
  const [rowSelect, setRowSelected] = useState(false);

  return (
    <AppointmentRow rowSelected={rowSelect}>
      <input type="checkbox" onChange={() => setRowSelected(!rowSelect)}/>
      <Column>{startDate} - {endDate}</Column>
      <Column bold>{patientName}</Column>
      <Column>{insurance}</Column>
      <Column>{procedure}</Column>
      <Column>{technician}</Column>
      <Column>{location}</Column>
      <Status canceled={status === 'canceled'} confirmed={status === 'confirmed'}>
        {status === 'canceled' ? <MdOutlineCancel /> : (status === 'confirmed' ? <GiConfirmed /> : <CiCalendar />)} 
        {status}
      </Status>
      <IconWrapper>
        <PhoneIcon />
      </IconWrapper>
    </AppointmentRow>
  );
}

export default ListCard;