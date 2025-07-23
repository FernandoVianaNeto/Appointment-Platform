import { useState } from "react";
import { Input, ModalContainer, Overlay } from "./styles";

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSave: (data: any) => void;
}

function CreatePatientModal ({ isOpen, onClose, onSave }: ModalProps) {
  if (!isOpen) return null;
  const [name, setName] = useState<string>("");
  const [isMissingName, setIsMissingName] = useState<boolean>(false);
  const [isMissingInsurance, setIsMissingInsurance] = useState<boolean>(false);
  const [isMissingPhone, setIsMissingPhone] = useState<boolean>(false);
  const [isMissingRequiredFields, setIsMissingRequiredFields] = useState<boolean>(false);

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const form = e.currentTarget;

    if (form.patientName.value === "") {
      setIsMissingRequiredFields(true);
      setIsMissingName(true);
    }

    if (form.insurance.value === '') {
      setIsMissingInsurance(true); 
      setIsMissingRequiredFields(true);
    }

    if (form.phone.value === '') {
      setIsMissingPhone(true); 
      setIsMissingRequiredFields(true);
    }

    const data = {
      name: form.patientName.value,
      insurance: form.insurance.value,
      phone: form.phone.value,
      address: form.address.value ?? '',
      email: form.email.value ?? '',
    };

    onSave(data);
  };

  return (
    <Overlay>
      <ModalContainer>
        <h2>Create New Patient</h2>
        <form method="post" onSubmit={handleSubmit}>
          <div style={{ position: 'relative' }}>
            <label>
              Patient Name*:
              <Input
                type="text"
                name="patientName"
                value={name}
                onChange={(e) => {
                  setName(e.target.value);
                }}
                className="patient-input"
                missingField={isMissingName}
              />
            </label>
          </div>

          <label>
            Insurance*:
            <Input className="insurance-input" type="text" name="insurance" missingField={isMissingInsurance}/>
          </label>

          <label>
            Phone*:
            <Input className="phone-input" type="phone" name="phone" missingField={isMissingPhone}/>
          </label>

          <label>
            Email:
            <Input className="email-input" type="text" name="email" />
          </label>

          <label>
            Address:
            <Input className="address-input" type="text" name="address" />
          </label>
          {
            isMissingRequiredFields && <p>Missing required fields</p>
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

export default CreatePatientModal;