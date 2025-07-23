import { useState } from "react";
import { ModalContainer, Overlay } from "./styles";
import type { TPatientData } from "../../types/patient";

interface ModalProps {
  isOpen: boolean;
  patient: TPatientData
  onClose: () => void;
  onEdit: (data: any) => void;
}

function EditPatientModal ({ isOpen, onClose, onEdit, patient }: ModalProps) {
  if (!isOpen) return null;
  const [name, setName] = useState<string>(patient.name);
  const [address, setAddress] = useState<string>(patient.address);
  const [phone, setPhone] = useState<string>(patient.phone);
  const [email, setEmail] = useState<string>(patient.email);
  const [insurance, setInsurance] = useState<string>(patient.insurance);
  const [isMissingRequiredFields, setIsMissingRequiredFields] = useState<boolean>();
  // const [missingRequiredFields, setMissingRequiredFields] = useState<boolean>();

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const form = e.currentTarget;

    if (
      form.patientName.value === "" ||
      form.patientName.value === undefined ||
      form.insurance.value === undefined || 
      form.insurance.value === "" || 
      form.patientName.value === undefined || 
      form.phone.value === undefined ||
      form.phone.value === ""
    ) {
      setIsMissingRequiredFields(true);
    } else {
      setIsMissingRequiredFields(false);
    }
    
    const data: TPatientData = {
        uuid: patient.uuid,
        address: address,
        email: email,
        insurance: insurance,
        name: name,
        phone: phone
    };

    onEdit(data);
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
            <input className="insurance-input" type="text" name="insurance" value={insurance} onChange={(e) => setInsurance(e.target.value)}/>
          </label>

          <label>
            Address:
            <input className="address-input" type="text" name="address" value={address} onChange={(e) => setAddress(e.target.value)}/>
          </label>

          <label>
            Email:
            <input className="email-input" type="text" name="email" value={email} onChange={(e) => setEmail(e.target.value)}/>
          </label>

          <label>
            Phone:
            <input className="phone-input" type="text" name="phone" value={phone} onChange={(e) => setPhone(e.target.value)}/>
          </label>
          {
            isMissingRequiredFields && <p>Missing required fields</p>
          }
          <div className="actions">
            <button type="submit">Edit</button>
            <button type="button" onClick={onClose}>Cancel</button>
          </div>
          
        </form>
      </ModalContainer>
    </Overlay>
  );
};

export default EditPatientModal;