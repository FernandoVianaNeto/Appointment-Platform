import styled from 'styled-components';

export const Container = styled.div`
  display: flex;
`;

export const H1 = styled.h1`
  font-family: ${({ theme }) => theme.font.family};
  color: ${({ theme }) => theme.colors.primary};
`;

export const Wrapper = styled.div`
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 30px;
`;

export const Div = styled.div`
  display: flex;
  gap: 20px;
`;