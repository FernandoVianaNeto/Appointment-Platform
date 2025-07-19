import styled from 'styled-components';
import { FiPhone } from 'react-icons/fi';

export const AppointmentRow = styled.div`
  display: flex;
  align-items: center;
  gap: 16px;
  background-color: #f5f6ff;
  padding: 12px 16px;
  border-radius: 8px;
  margin-bottom: 8px;
  font-size: 14px;
  color: #333;
`;

export const Column = styled.div<{ bold?: boolean }>`
  flex: 1;
  font-weight: ${({ bold }) => (bold ? '600' : '400')};
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  width: 300px;
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
    width: 200px;
    display: flex;
    align-items: center;
    justify-content: center;
  `;
  
  export const IconWrapper = styled.div`
    background-color: white;
    border-radius: 50%;
    padding: 8px;
    display: flex;
    align-items: center;
    justify-content: center;
  `;
  
  export const PhoneIcon = styled(FiPhone)`
    color: #6a5acd;
    font-size: 16px;
  `;
  