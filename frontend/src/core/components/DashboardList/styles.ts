import styled from 'styled-components';

export const ScrollableList = styled.ul<{
  noContent?: boolean
}>`
  height: 350px;
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
    height: ${({ noContent }) => noContent && '350px'};
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

  @media (min-width: 768px) {
    max-height: 350px;
  }

  @media (min-height: 900px) {
    max-height: 600px;
    padding: 0;
  }
`;
