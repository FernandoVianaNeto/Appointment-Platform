import styled from 'styled-components';

export const Container = styled.section`
    display: flex;
    flex-direction: column;
    align-items: center;
    height: 100vh;
    background-color: ${({ theme }) => theme.colors.primary};
    width: 10vw;
    position: absolute;
    color: ${({theme}) => theme.colors.white};
    border: none;
`;