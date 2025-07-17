import styled from 'styled-components';

export const Container = styled.div`
  display: flex;
  height: 100vh;
  width: 90vw;
  padding: 50px;
`;

export const Button = styled.button`
  width: 100%;
  padding: ${({ theme }) => theme.spacing.sm};
  background-color: ${({ theme }) => theme.colors.primary};
  color: white;
  border: none;
  border-radius: ${({ theme }) => theme.borderRadius};
  cursor: pointer;
  font-weight: bold;
`;
