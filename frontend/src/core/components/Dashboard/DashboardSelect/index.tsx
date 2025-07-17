import { Container } from './styles';

type Props = {
  children: React.ReactNode;
};

function DashboardSelect({ children }: Props) {
  return (
    <Container>
      {children}
    </Container>
  );
}

export default DashboardSelect;