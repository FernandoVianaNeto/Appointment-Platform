import { ButtonText, Container, IconButton, SearchButton, Wrapper } from './styles';
import { RiAccountCircleFill } from "react-icons/ri";
import { IoMdSettings } from "react-icons/io";

type Props = {
  children: React.ReactNode;
};

function Header({children}: Props) {
  return (
    <Container>
        <Wrapper>
            {children}
            <SearchButton> 
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