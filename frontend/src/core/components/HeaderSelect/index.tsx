import { useState } from "react";
import { DropdownIcon, Options, SelectWrapper, StyledSelect } from "./styles";
import HeaderOption from "../HeaderOption";


type Props = {
  options: string[];
  defaultValue?: string;
  onChange?: (value: string) => void;
};

function HeaderSelect({ options, defaultValue = 'All', onChange }: Props) {
  const [open, setOpen] = useState(false);
  const [selected, setSelected] = useState(defaultValue);

  function toggleOpen() {
    setOpen(!open);
  }

  function handleSelect(option: string) {
    setSelected(option);
    onChange?.(option);
    setOpen(false);
  }

  return (
    <SelectWrapper>
      <StyledSelect onClick={toggleOpen}>
        {selected}
        <DropdownIcon open={open} />
      </StyledSelect>
      {open && (
        <Options>
          {options.map((option) => (
            <HeaderOption
              key={option}
              value={option}
              selected={option === selected}
              onClick={handleSelect}
            >
              {option}
            </HeaderOption>
          ))}
        </Options>
      )}
    </SelectWrapper>
  );
}


export default HeaderSelect;