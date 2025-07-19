import { AppointmentRow, Column, IconWrapper, PhoneIcon } from './styles';

interface ListSummaryProps {
    fields: string[];
};

function ListSummary({ fields }: ListSummaryProps) {
  return (
    <AppointmentRow>
        <input type="checkbox" />
        {
            fields.map((field) => (
                <Column>
                    {field}
                </Column>
            ))
            
        }
        <IconWrapper>
            <PhoneIcon />
        </IconWrapper>
    </AppointmentRow>
  );
}

export default ListSummary;