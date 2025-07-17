import { Container } from './styles';

interface SideBarButtonProps {
    text: string;
    highlight?: boolean,
}

function SideBarButton({ text, highlight }: SideBarButtonProps) {
  return (
    <Container highlight={highlight}>
       {text}
    </Container>
  );
}

export default SideBarButton;