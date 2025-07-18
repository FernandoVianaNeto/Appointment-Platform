import { useState } from "react";
import { DropdownIcon, Options, SelectWrapper, StyledSelect } from "./styles";

type Props = {
    children: React.ReactNode;
  };

function HeaderSelect({children}: Props) {
    const [open, setOpen] = useState(false);
  
    function toggleOpen() {
      setOpen(!open);
    }
  
    return (
      <SelectWrapper>
        <StyledSelect onClick={toggleOpen}>
          <DropdownIcon open={open} />
        </StyledSelect>
        {open && (
          <Options>
            {children}
          </Options>
        )}
      </SelectWrapper>
    );
  }
  

export default HeaderSelect;