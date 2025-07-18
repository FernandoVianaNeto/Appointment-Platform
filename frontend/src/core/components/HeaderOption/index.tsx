import { Option } from "./styles";

type Props = {
  selected: boolean;
  children: React.ReactNode;
};

function HeaderOption({ children, selected }: Props) {
    return (
      <Option selected={selected}>
        {children}
      </Option>
    );
  }
  

export default HeaderOption;