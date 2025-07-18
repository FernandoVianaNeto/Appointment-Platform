import styled from 'styled-components';

export const Container = styled.button<{
    highlight?: boolean
}>`
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 10px;
    border-radius: 4px;
    padding: 20px 20px;
    cursor: pointer;
    background-color: ${({ theme, highlight }) => highlight ? theme.colors.highlightButton : theme.colors.primary };
    border: none;
    padding: 20px 10px;
    color: ${({ theme }) => theme.colors.white };
    font-weight: 500;
    min-width: 200px;
    height: 40px;
`;