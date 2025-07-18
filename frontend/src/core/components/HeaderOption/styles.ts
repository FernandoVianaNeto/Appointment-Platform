import styled from "styled-components";

export const Option = styled.li<{
    selected?: boolean
}>`
  padding: 10px 15px;
  color: ${({ selected }) => (selected ? '#3b49df' : '#7a8cbc')};
  cursor: pointer;
  display: flex;
  align-items: center;

  &:hover {
    background: #f0f3ff;
  }
`;

export const CheckMark = styled.span<{
    visible?: boolean
}>`
  width: 16px;
  height: 16px;
  border: 2px solid #3b49df;
  border-radius: 3px;
  margin-left: auto;
  display: ${({ visible }) => (visible ? 'block' : 'none')};

  &::after {
    content: '';
    display: block;
    width: 6px;
    height: 10px;
    border: solid #3b49df;
    border-width: 0 2px 2px 0;
    transform: rotate(45deg);
    margin: 1px 0 0 4px;
  }
`;