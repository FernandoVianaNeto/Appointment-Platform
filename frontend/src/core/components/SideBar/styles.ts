import styled from 'styled-components';

export const Container = styled.section`
    display: flex;
    flex-direction: column;
    align-items: center;
    height: 100vh;
    background-color: ${({ theme }) => theme.colors.primary};
    width: 10vw;
    color: ${({theme}) => theme.colors.white};
    border: none;
`;

export const Logo = styled.h1`
  display: flex;
  align-items: center;
  justify-content: center;
  padding-bottom: 20px;
  border-bottom: 0.5px solid ${({ theme }) => theme.colors.highlighBackground};
  padding: 30px;
`;

export const ButtonWrapper = styled.div`
  margin-top: 50px;
`;
