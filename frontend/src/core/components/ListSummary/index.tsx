import { useState } from 'react';
import { AppointmentRow, Column, IconWrapper, PhoneIcon } from './styles';

interface ListSummaryProps {
    fields: string[];
    onChange?: (value: boolean) => void;
};

function ListSummary({ fields, onChange }: ListSummaryProps) {
    const [rowsSelected, setRowsSelected] = useState(false);

    const handleRowsSelected = () => {
        setRowsSelected(true);
        onChange?.(rowsSelected);
      };

    return (
        <AppointmentRow>
            <input type="checkbox" onChange={() => handleRowsSelected()}/>
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