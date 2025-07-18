import styled, { keyframes } from "styled-components";

const slideIn = keyframes`
  from {
    opacity: 0;
    transform: translateY(4px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
`;

export const Option = styled.li<{
    selected?: boolean
}>`
  padding: 10px 15px;
  color: ${({ selected, theme }) => (selected ? theme.colors.text : theme.colors.primary )};
  cursor: pointer;
  display: flex;
  align-items: center;
  animation: ${slideIn} 0.2s ease-out;

  &:hover {
    background: #f0f3ff;
  }
`;

const fadeIn = keyframes`
  from {
    opacity: 0;
    transform: scale(0.8);
  }
  to {
    opacity: 1;
    transform: scale(1);
  }
`;

export const CheckMark = styled.span<{
  visible?: boolean;
}>`
  width: 16px;
  height: 16px;
  border-radius: 3px;
  margin-left: auto;
  display: ${({ visible }) => (visible ? 'flex' : 'none')};
  align-items: center;
  justify-content: center;
  animation: ${({ visible }) => (visible ? fadeIn : 'none')} 0.2s ease;

  &::after {
    content: '';
    width: 6px;
    height: 10px;
    border: solid ${({theme}) => theme.colors.text};
    border-width: 0 2px 2px 0;
    transform: rotate(45deg);
    display: inline-block;
  }
`;