import { AppointmentRow, IconWrapper, Column, Status, PhoneIcon, EditIcon } from './styles';
import { CiCalendar } from "react-icons/ci";
import { MdOutlineCancel } from "react-icons/md";
import { GiConfirmed } from "react-icons/gi";
import { useEffect, useState } from 'react';

type Props = {
  uuid: string,
  patientName: string,
  insurance: string,
  procedure: string,
  technician: string,
  location: string,
  status: string,
  startDate: string,
  endDate: string,
  rowSelected?: boolean,
  onRowSelected?: (uuid: string) => void;
  onEdit?: () => void;
};

function ListCard({ uuid, patientName, insurance, location, procedure, status, technician, startDate, endDate, rowSelected, onRowSelected, onEdit }: Props) {
  const [rowSelect, setRowSelected] = useState(rowSelected);

  useEffect(() => {
    setRowSelected(rowSelected);
  }, [rowSelected]);

  function handleRowSelected() {
    setRowSelected(!rowSelect);
    onRowSelected?.(uuid);
  }

  return (
    <AppointmentRow rowSelected={rowSelect}>
      <input type="checkbox" onChange={handleRowSelected}  checked={rowSelect}/>
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
      <IconWrapper>
        <button type="button" onClick={onEdit}>
          <EditIcon /> 
        </button>
      </IconWrapper>
    </AppointmentRow>
  );
}

export default ListCard;