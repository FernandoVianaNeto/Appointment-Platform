import styled from 'styled-components';

export const Container = styled.header`
    display: flex;
    justify-content: space-between;
    width: 100%;
    padding: 40px 0px;
    max-height: 120px;
    color: ${({theme}) => theme.colors.white};
    border-bottom: 2px solid ${({theme}) => theme.colors.highlighBackground};
`;

export const Wrapper = styled.div`
    display: flex;
`;

export const ButtonText = styled.p`
    font-size: 14px;
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
    background-color: ${({theme}) => theme.colors.background};
    border: none;
    padding: 13px 20px;
    font-size: 14px;
    font-weight: 600;

    display: flex;
    justify-content: space-between;
    align-items: center;
    cursor: pointer;
    margin-left: 20px;
`;

export const Select = styled.input`
    border: 1px solid ${({theme}) => theme.colors.grey};
    padding: 10px;
`;