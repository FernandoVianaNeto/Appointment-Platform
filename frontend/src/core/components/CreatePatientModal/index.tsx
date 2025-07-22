import { useState } from "react";
import { ModalContainer, Overlay } from "./styles";

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSave: (data: any) => void;
}

function CreatePatientModal ({ isOpen, onClose, onSave }: ModalProps) {
  if (!isOpen) return null;
  const [name, setName] = useState<string>("");

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const form = e.currentTarget;

    if ( 
      form.patientName.value === undefined ||
      form.insurance.value === undefined || 
      form.phone.value === undefined || 
      form.address.value === undefined || 
      form.email.value === undefined
    ) {
      throw new Error('Could not save patient. Missing Required fields')
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
              <input
                type="text"
                name="patientName"
                value={name}
                onChange={(e) => {
                  setName(e.target.value);
                }}
                className="patient-input"
              />
            </label>
          </div>

          <label>
            Insurance:
            <input className="insurance-input" type="text" name="insurance" />
          </label>

          <label>
            Phone*:
            <input className="phone-input" type="phone" name="phone" />
          </label>

          <label>
            Email:
            <input className="email-input" type="text" name="email" />
          </label>

          <label>
            Address:
            <input className="address-input" type="text" name="address" />
          </label>

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