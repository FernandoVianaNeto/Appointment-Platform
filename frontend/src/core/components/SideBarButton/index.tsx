import { Container } from './styles';

interface SideBarButtonProps {
    text: string;
    highlight?: boolean,
    onClick?: () => void
}

function SideBarButton({ text, highlight, onClick }: SideBarButtonProps) {
  return (
    <Container highlight={highlight} onClick={onClick}>
       {text}
    </Container>
  );
}

export default SideBarButton;