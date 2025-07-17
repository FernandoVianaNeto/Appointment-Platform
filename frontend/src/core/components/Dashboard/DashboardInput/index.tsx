import { Container } from './styles';

type Props = {
  children: React.ReactNode;
};

function Dashboard({ children }: Props) {
  return (
    <Container>
      {children}
    </Container>
  );
}

export default Dashboard;