import styled from 'styled-components';

export const Wrapper = styled.div`
    display: flex;
    align-items: center;
    gap: 16px;
    font-family: 'sans-serif';
    min-width: 250px;
`;

export const CircleButton = styled.button`
    width: 32px;
    height: 32px;
    border-radius: 40px;
    border: none;
    background: transparent;
    cursor: pointer;
    font-size: 20px;
    color: ${({ theme }) => theme.colors.highlightButton};
`;

export const DateContainer = styled.div`
    display: flex;
    align-items: center;
    font-size: 12px;
    color: ${({ theme }) => theme.colors.primary};
`;

export const TodayText = styled.span`
    display: flex;
    align-items: center;
    font-size: 12px;
    cursor: pointer;
    padding: 5px;
`;

