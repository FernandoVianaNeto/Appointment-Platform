import styled from 'styled-components';

export const Wrapper = styled.div`
    display: flex;
    align-items: center;
    gap: 24px;
    font-family: 'sans-serif';
`;

export const CircleButton = styled.button`
    width: 32px;
    height: 32px;
    border-radius: 40px;
    border: 1px solid ${({theme}) => theme.colors.primary};
    background: transparent;
    cursor: pointer;
    font-size: 16px 32px;
`;

export const DateContainer = styled.div`
    display: flex;
    align-items: center;
    font-size: 12px;
`;

export const TodayText = styled.span`
    display: flex;
    align-items: center;
    font-size: 12px;
`;

