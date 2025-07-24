import { useState } from "react";
import { Input, ModalContainer, Overlay } from "./styles";

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSave: (data: any) => void;
}

function SignUpModal({ isOpen, onClose, onSave }: ModalProps) {
  if (!isOpen) return null;

  const [isMissingRequiredFields, setIsMissingRequiredFields] = useState<boolean>(false);
  const [isMissingEmail, setIsMissingEmail] = useState<boolean>(false);
  const [isMissingClinicName, setIsMissingClinicName] = useState<boolean>(false);
  const [isMissingPassword, setIsMissingPassword] = useState<boolean>(false);
  const [isMissingPasswordConfirmation, setIsMissingPasswordConfirmation] = useState<boolean>(false);
  const [isDifferentPasswords, setIsDifferentPasswords] = useState<boolean>(false);

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const form = e.currentTarget;

    if (form.email.value === '') {
      setIsMissingRequiredFields(true);
      setIsMissingEmail(true);
    }

    if (form.clinicName.value === '') {
      setIsMissingRequiredFields(true);
      setIsMissingClinicName(true);
    }

    if (form.password.value === '') {
      setIsMissingRequiredFields(true);
      setIsMissingPassword(true);
    }

    if (form.passwordConfirmation.value === '') {
      setIsMissingRequiredFields(true);
      setIsMissingPasswordConfirmation(true);
    }

    if (form.password.value !== form.passwordConfirmation.value) {
      setIsDifferentPasswords(true);
    }
  
    const data = {
      email: form.email.value,
      clinicName: form.clinicName.value,
      password: form.password.value,
    };

    onSave(data);
  };

  return (
    <Overlay>
      <ModalContainer>
        <h2>Create New Clinic</h2>
        <form method="post" onSubmit={handleSubmit}>
          <label>
            Email*:
            <Input className="email-input" type="text" name="email" missingField={isMissingEmail}/>
          </label>

          <label>
            Clinic Name*:
            <Input className="clinic-name-input" type="text" name="clinicName" missingField={isMissingClinicName}/>
          </label>

          <label>
            Password*:
            <Input className="password-input" type="password" name="password" missingField={isMissingPassword}/>
          </label>

          <label>
            Confirm Password*:
            <Input className="confirm-password-input" type="password" name="passwordConfirmation" missingField={isMissingPasswordConfirmation}/>
          </label>
         {
          isMissingRequiredFields && <p>Missing required fields</p>
         }
         {
          isDifferentPasswords && <p>Passwords need to be equal</p>
         }
          <div className="actions">
            <button type="submit">Save</button>
            <button type="button" onClick={onClose}>Cancel</button>
          </div>

        </form>
      </ModalContainer>
    </Overlay>
  );
};

export default SignUpModal;