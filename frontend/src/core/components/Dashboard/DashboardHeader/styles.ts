import styled from 'styled-components';

export const Container = styled.header`
    display: flex;
    justify-content: space-between;
    width: 100%;
    padding: 50px;
    color: ${({theme}) => theme.colors.white};
    border-bottom: 2px solid ${({theme}) => theme.colors.highlighBackground};
`;

export const Wrapper = styled.div`
    display: flex;
`;

export const ButtonText = styled.p`
    font-size: 12px;
    margin-left: 5px;
`;

export const IconButton = styled.button`
    border: none;
    background-color: transparent;
    font-size: 20px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    cursor: pointer;
    margin-left: 20px;
`;

export const SearchButton = styled.button`
    border: none;
    background-color: transparent;
    font-size: 20px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    cursor: pointer;
    margin-left: 20px;
`;

export const Input = styled.input`
`;