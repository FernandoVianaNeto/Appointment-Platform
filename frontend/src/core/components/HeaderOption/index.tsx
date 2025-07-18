import { CheckMark, Option } from "./styles";

type Props = {
  selected: boolean;
  value: string;
  onClick: (value: string) => void;
  children: React.ReactNode;
};

function HeaderOption({ children, selected, value, onClick }: Props) {
    return (
      <Option selected={selected} onClick={() => onClick(value)} value={value}>
        {children}
        <CheckMark visible={selected} />
      </Option>
    );
  }
  

export default HeaderOption;