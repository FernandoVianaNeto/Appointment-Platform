import { Container } from './styles';
import { FaPlus } from "react-icons/fa6"
import { MdEdit } from "react-icons/md";
interface CreationButtonProps {
    text: string;
    highlight?: boolean,
    edit?: boolean
}

function CreationButton({ text, highlight, edit }: CreationButtonProps) {
  return (
    <Container highlight={highlight}>
      {
        edit ? <MdEdit /> : <FaPlus />
      }
      {text}
    </Container>
  );
}

export default CreationButton;