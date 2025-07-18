import styled from 'styled-components';

export const SelectWrapper = styled.div`
  position: relative;
  width: 200px;
  font-family: Arial, sans-serif;
`;

export const StyledSelect = styled.div`
  background: white;
  border: 1px solid #ccc;
  border-radius: 6px;
  padding: 10px 40px 10px 15px;
  cursor: pointer;
  color: #333;
  user-select: none;
  position: relative;
`;

export const DropdownIcon = styled.span<{
    open?: boolean
}>`
  position: absolute;
  right: 15px;
  top: 50%;
  transform: translateY(-50%);
  border: solid #666;
  border-width: 0 2px 2px 0;
  display: inline-block;
  padding: 4px;
  transform-origin: center;
  transition: transform 0.2s ease;
  transform: ${({ open }) => (open ? 'translateY(-50%) rotate(-135deg)' : 'translateY(-50%) rotate(45deg)')};
`;

export const Options = styled.ul`
  position: absolute;
  width: 100%;
  background: white;
  margin: 5px 0 0 0;
  padding: 0;
  list-style: none;
  border-radius: 6px;
  box-shadow: 0 4px 8px rgb(0 0 0 / 0.1);
  max-height: 150px;
  overflow-y: auto;
  z-index: 10;
`;

