import styled from 'styled-components';
import { FiPhone } from 'react-icons/fi';
import { MdEdit } from "react-icons/md";

export const AppointmentRow = styled.div<{
  rowSelected?: boolean;
}>`
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 12px 16px;
  border-radius: 8px;
  margin-bottom: 8px;
  font-size: 14px;
  color: #333;
  background-color: ${({ theme, rowSelected }) => rowSelected && theme.colors.highlighBackground};
  transition: transform 1s ease;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  td {
    vertical-align: middle;
  }
`;

export const Column = styled.div<{ bold?: boolean }>`
  flex: 1;
  font-weight: ${({ bold }) => (bold ? '600' : '400')};
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 250px;
  display: flex;
  align-items: center;
`;

export const Status = styled.div<{
  confirmed?: boolean,
  canceled?: boolean,
}>`
  display: flex;
  align-items: center;
  gap: 6px;
  color: ${({ confirmed, canceled }) => confirmed ? 'green' : (canceled ? 'red' : '#b0ae31')};
  font-weight: 500;
  flex-shrink: 0;
  width: 150px;
  display: flex;
  align-items: center;
`;

export const IconWrapper = styled.div`
  background-color: white;
  padding: 8px;
  display: flex;
  align-items: center;
  justify-content: center;

  button {
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: transparent;
    border: none;
    cursor: pointer;
  }
`;

export const PhoneIcon = styled(FiPhone)`
  color: #6a5acd;
  font-size: 16px;
`;

export const EditIcon = styled(MdEdit)`
  color: #6a5acd;
  font-size: 16px;
`;