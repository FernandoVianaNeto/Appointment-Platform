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

export const ModalContainer = styled.div<{
}>`
  background: white;
  padding: 32px;
  border-radius: 16px;
  width: 100%;
  max-width: 650px;
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.15);

  h2 {
    text-align: center;
    font-size: 24px;
    font-weight: 600;
    margin-bottom: 24px;
    color: ${({ theme }) => theme.colors.primary};
  }

  form {
    display: flex;
    flex-direction: column;
    gap: 20px;

    label {
      display: flex;
      flex-direction: column;
      font-size: 14px;
      font-weight: 600;
      color: ${({ theme }) => theme.colors.primary};
      gap: 5px;
    }

    select {
      padding: 10px 12px;
      font-size: 14px;
      border: 1px solid #cfd4dc;
      border-radius: 8px;
      outline: none;
      transition: border-color 0.2s;

      &:focus {
        border-color: #5a67d8;
      }
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
  }
`;

export const Input = styled.input<{
  missingField?: boolean,
}>`
  padding: 10px 12px;
  font-size: 14px;
  border: 1px solid ${({ missingField }) => missingField ? 'red' : '#cfd4dc' } ;
  border-radius: 8px;
  outline: none;
  transition: border-color 0.2s;

  &:focus {
    border-color: #5a67d8;
  }
`;

export const DateWrapper = styled.div`
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  justify-content: space-between;

  label {
    flex: 1;
    min-width: 220px;
    gap: 10px;
  }
`;

export const InputWrapper = styled.div`
  display: flex;
  gap: 10px;

  input {
    flex: 1;
  }
`;
