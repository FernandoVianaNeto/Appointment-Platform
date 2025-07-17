import { ButtonText, Container, IconButton, Input, SearchButton, Wrapper } from './styles';
import { RiAccountCircleFill } from "react-icons/ri";
import { IoMdSettings } from "react-icons/io";

type Props = {
  children: React.ReactNode;
};

function DashboardHeader({children}: Props) {
  return (
    <Container>
        <Wrapper>
            {children}
            <Input value='Search'/>
            <SearchButton />
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

export default DashboardHeader;