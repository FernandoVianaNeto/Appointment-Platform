import styled from 'styled-components';

export const Container = styled.div`
  display: flex;
`;

export const H1 = styled.h1`
  font-family: ${({ theme }) => theme.font.family};
  color: ${({ theme }) => theme.colors.primary};
  margin-top: 30px;
`;

export const Wrapper = styled.div`
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
`;

export const Div = styled.div`
`;