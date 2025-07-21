import { Container, Input } from "./styles";

interface Props {
  onChange: (searchTerm: string) => void
}

export const HeaderInput = ({ onChange }: Props) => {
  return (
    <Container>
      <Input
        id="search-input"
        type="text"
        placeholder="Type to search..."
        onChange={(e) => onChange(e.target.value)}
      />
    </Container>
  );
};


export default HeaderInput;