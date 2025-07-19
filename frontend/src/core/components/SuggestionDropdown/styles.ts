import styled from "styled-components";

export const Container = styled.button<{
    highlight?: boolean
}>`
    width: 100%;
    cursor: pointer;
    background-color: ${({ theme, highlight }) => highlight ? theme.colors.highlightButton : theme.colors.primary };
    border: none;
    padding: 20px 10px;
    color: ${({ theme }) => theme.colors.white };
    font-weight: 500;
`;