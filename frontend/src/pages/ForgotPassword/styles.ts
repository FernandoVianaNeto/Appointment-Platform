import styled from "styled-components";

export const Overlay = styled.div`
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 999;
`;

export const ModalContainer = styled.div`
  background: white;
  padding: 32px;
  border-radius: 12px;
  width: 100%;
  max-width: 480px;
  box-shadow: 0 10px 35px rgba(0, 0, 0, 0.2);

  h2 {
    text-align: center;
    font-size: 22px;
    font-weight: 600;
    color: ${({ theme }) => theme.colors.primary || '#2D3748'};
    margin-bottom: 8px;
  }

  #subtitle {
    text-align: center;
    font-size: 14px;
    color: #4a5568;
    margin-bottom: 24px;
  }
`;

export const ModalWrapper = styled.div`
  display: flex;
  flex-direction: column;
  gap: 16px;

  label {
    font-size: 14px;
    font-weight: 500;
    color: #2d3748;
  }

  .actions {
    display: flex;
    justify-content: flex-end;
    gap: 10px;
    margin-top: 10px;

    button {
      padding: 10px 18px;
      border: none;
      border-radius: 6px;
      font-size: 14px;
      font-weight: 500;
      cursor: pointer;
      transition: background-color 0.2s ease;

      &.primary {
        background-color: #2b6cb0;
        color: white;

        &:hover {
          background-color: #2c5282;
        }
      }

      &.secondary {
        background-color: #e2e8f0;
        color: #1a202c;

        &:hover {
          background-color: #cbd5e0;
        }
      }
    }
  }
`;

export const ErrorText = styled.p`
  color: red;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  font-size: 12px;
`;

export const Input = styled.input<{ missingField?: boolean }>`
  padding: 10px 12px;
  font-size: 14px;
  border: 1px solid ${({ missingField }) => missingField ? 'red' : '#cbd5e0'};
  border-radius: 8px;
  outline: none;
  transition: border-color 0.2s;

  &:focus {
    border-color: #3182ce;
  }

  &::placeholder {
    color: #a0aec0;
  }
`;
