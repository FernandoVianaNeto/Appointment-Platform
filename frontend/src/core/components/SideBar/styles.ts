import styled from 'styled-components';

export const Container = styled.section`
    display: flex;
    flex-direction: column;
    align-items: center;
    height: 100vh;
    background: linear-gradient(116deg, #344293 -0.38%, rgba(52, 66, 147, 0.00) 66.54%), #344293;
    width: 10vw;
    position: absolute;
    color: ${({theme}) => theme.colors.white};
    border: none;
`;

export const Logo = styled.h1`
  display: flex;
  align-items: center;
  justify-content: center;
  padding-bottom: 20px;
  padding: 30px;
  height: 100px;
`;

export const ButtonWrapper = styled.div`
  margin-top: 50px;
`;
