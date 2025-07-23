import styled from "styled-components";

export const Overlay = styled.div`
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 999;
`;

export const ModalContainer = styled.div`
    display: flex;
    flex-direction: column;
    background: white;
    padding: 32px;
    border-radius: 16px;
    width: 100%;
    max-width: 350px;
    box-shadow: 0 8px 30px rgba(0, 0, 0, 0.15);
    gap: 20px;

    p {
        display: flex;
        align-items: center;
        justify-content: center;
        text-align: center;
        font-size: 16px;
    }

    .actions {
        display: flex;
        justify-content: flex-end;
        gap: 12px;
        margin-top: 12px;

        button {
        padding: 10px 18px;
        border: none;
        border-radius: 6px;
        font-size: 14px;
        font-weight: 500;
        cursor: pointer;
        transition: all 0.2s ease;

        &:first-child {
            background-color: #5a67d8;
            color: white;

            &:hover {
            background-color: #434190;
            }
        }

        &:last-child {
            background-color: #e2e8f0;

            &:hover {
            background-color: #cbd5e0;
            }
        }
    }
  }
`;