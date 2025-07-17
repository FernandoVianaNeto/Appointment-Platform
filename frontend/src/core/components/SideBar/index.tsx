import { ButtonWrapper, Container, Logo } from './styles';

type Props = {
  children: React.ReactNode;
};

function SideBar({children}: Props) {
  return (
    <Container>
        <Logo>
            Platform
        </Logo>
        <ButtonWrapper>
          {children}
        </ButtonWrapper>
    </Container>
  );
}

export default SideBar;