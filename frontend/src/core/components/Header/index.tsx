import { ButtonText, Container, IconButton, SearchButton, Wrapper } from './styles';
import { RiAccountCircleFill } from "react-icons/ri";
import { IoMdSettings } from "react-icons/io";

type Props = {
  children: React.ReactNode;
  onSubmit: () => void;
};

function Header({children, onSubmit }: Props) {
  return (
    <Container>
        <Wrapper>
            {children}
            <SearchButton onClick={() => onSubmit()}> 
              Search 
            </SearchButton>
        </Wrapper>
        <Wrapper>
            <IconButton>
                <IoMdSettings />
            </IconButton>
            <IconButton>
                <RiAccountCircleFill />
                <ButtonText>My Account</ButtonText>
            </IconButton>
        </Wrapper>
    </Container>
  );
}

export default Header;