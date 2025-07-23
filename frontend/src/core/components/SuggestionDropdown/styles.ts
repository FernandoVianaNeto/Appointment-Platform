import styled from "styled-components";

export const Ul = styled.ul<{
    highlight?: boolean
}>`
    margin-top: 4px;
    padding: 0;
    list-style: none;
    border: 1px solid #ccc;
    border-radius: 8px;
    max-height: 150;
    overflow-y: auto;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
    position: absolute;
    background-color: white;
    z-index: 1000;
    width: 100%;

    li {
        padding: 10px;
        cursor: pointer;
    }
`;

export const NoPatientsWrapper = styled.div<{
    highlight?: boolean
}>`
    margin-top: 4px;
    padding: 30px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 10px;
    max-height: 150;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
    position: absolute;
    background-color: white;
    z-index: 1000;
    width: 100%;
`;