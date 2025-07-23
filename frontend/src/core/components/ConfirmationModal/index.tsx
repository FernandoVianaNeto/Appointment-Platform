import { ModalContainer, Overlay } from "./styles";

interface ModalProps {
    confirmationText: string;
    isOpen: boolean;
    onClose: () => void;
    onConfirm: (data: any) => void;
}

function ConfirmationModal ({ confirmationText, isOpen, onClose, onConfirm }: ModalProps) {
  if (!isOpen) return null;

  return (
    <Overlay>
      <ModalContainer>
        <p>{confirmationText}</p>
        <div className="actions">
          <button type="button" onClick={onConfirm}>Confirm</button>
          <button type="button" onClick={onClose}>Cancel</button>
        </div>
      </ModalContainer>
    </Overlay>
  );
};

export default ConfirmationModal;