import { useState } from 'react';
import { AppointmentRow, Column, EditIcon, IconWrapper, PhoneIcon } from './styles';

interface ListSummaryProps {
    fields: string[];
    lessColumns?: boolean
    onChange?: (value: boolean) => void;
};

function ListSummary({ fields, lessColumns, onChange }: ListSummaryProps) {
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
                    <Column key={field} lessColumns={lessColumns ?? false}>
                        {field}
                    </Column>
                ))
            }
            <IconWrapper>
                <PhoneIcon />
            </IconWrapper>
            <IconWrapper>
                <EditIcon />
            </IconWrapper>
        </AppointmentRow>
    );
}

export default ListSummary;