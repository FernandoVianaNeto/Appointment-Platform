import { AppointmentRow, IconWrapper, Column, PhoneIcon, EditIcon } from './styles';
import { useEffect, useState } from 'react';

type Props = {
  uuid: string,
  patientName: string,
  insurance: string,
  phone: string,
  email?: string,
  address?: string,
  rowSelected?: boolean,
  onRowSelected?: (uuid: string) => void;
  onEdit?: () => void;
};

function PatientCard({ uuid, patientName, insurance, phone, address, email, rowSelected, onRowSelected, onEdit }: Props) {
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
      <Column bold>{patientName}</Column>
      <Column>{insurance}</Column>
      <Column>{phone }</Column>
      <Column>{address ?? 'Not available'}</Column>
      <Column>{email ?? 'Not available'}</Column>
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

export default PatientCard;