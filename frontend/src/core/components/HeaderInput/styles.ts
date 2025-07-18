import styled from "styled-components";


export const Container = styled.div`
  display: flex;
  align-items: center;
  gap: 8px;
`;

export const Input = styled.input`
  border: 1px solid #ccc;
  border-radius: 0px 4px 4px 0px;
  padding: 6px 8px;
  font-size: 1rem;
  outline: none;
  height: 40px;
  
  &:focus {
    outline: none;
    box-shadow: none;
    border-color: #ccc;
  }
`;