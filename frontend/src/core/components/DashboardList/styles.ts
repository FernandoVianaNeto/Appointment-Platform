import styled from 'styled-components';

export const ScrollableList = styled.ul`
  max-height: 100%;
  overflow-y: auto;
  padding: 0;
  margin: 0;
  list-style: none;
  width: 100%;

  p {
    display: flex;
    align-items: center;
    justify-content: center;
    font-weight: 400;
  }

  &::-webkit-scrollbar {
    width: 8px;
  }

  &::-webkit-scrollbar-thumb {
    background: #ccc;
    border-radius: 4px;
  }

  &::-webkit-scrollbar-track {
    background: #f0f0f0;
  }
`;
