import styled from 'styled-components';

export const Container = styled.button<{
    highlight?: boolean
}>`
    width: 100%;
    cursor: pointer;
    background-color: ${({ theme, highlight }) => highlight ? theme.colors.highlightButton : theme.colors.primary };
    border: none;
    padding: 20px 10px;
    color: ${({ theme }) => theme.colors.white };
    font-weight: 500;
`;

export const Button = styled.button`
  width: 100%;
  padding: ${({ theme }) => theme.spacing.sm};
  background-color: ${({ theme }) => theme.colors.primary};
  color: white;
  border: none;
  border-radius: ${({ theme }) => theme.borderRadius};
  font-weight: bold;
`;
