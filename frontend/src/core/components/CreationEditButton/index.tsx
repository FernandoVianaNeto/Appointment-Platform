import { Button } from './styles';
import { FaPlus } from "react-icons/fa6"
import { MdEdit } from "react-icons/md";
interface CreationEditButtonProps {
    text: string;
    highlight?: boolean,
    edit?: boolean,
    onClick?: () => void,
}

function CreationEditButton({ text, highlight, edit, onClick }: CreationEditButtonProps) {
  function handleClick() {
    onClick?.();
  } 

  return (
    <Button highlight={highlight} onClick={handleClick}>
      {
        edit ? <MdEdit /> : <FaPlus />
      }
      {text}
    </Button>
  );
}

export default CreationEditButton;