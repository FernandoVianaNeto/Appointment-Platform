import { Container } from './styles';

type Props = {
  patient: string,
  insurance: string,
  procedure: string,
  technician: string,
  location: string,
  status: string,
};

function ListCard({ patient, insurance, location, procedure, status, technician }: Props) {
  return (
    <Container>
        
    </Container>
  );
}

export default ListCard;