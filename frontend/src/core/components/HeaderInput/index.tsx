import { Container, Input } from "./styles";

export const HeaderInput = () => {
    return (
      <Container>
        <Input id="search-input" type="text" placeholder="Type to search..." />
      </Container>
    );
  };

export default HeaderInput;