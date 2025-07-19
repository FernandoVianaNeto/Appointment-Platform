import { DateWrapper, InputWrapper, ModalContainer, Overlay } from "./styles";

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSave: (data: any) => void;
}

function CreateEditModal ({ isOpen, onClose, onSave }: ModalProps) {
  if (!isOpen) return null;

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const form = e.currentTarget;
    const data = {
      patientName: form.patientName.value,
      insurance: form.insurance.value,
      procedure: form.procedure.value,
      technician: form.technician.value,
      location: form.location.value,
      startDate: form.startDate.value,
      startTime: form.startTime.value,
      endDate: form.endDate.value,
      endTime: form.endTime.value,
    };

    onSave(data);
  };

  return (
    <Overlay>
      <ModalContainer>
        <h2>Create New Appointment</h2>
        <form method="post" onSubmit={handleSubmit}>
          <label>
            Patient Name:
            <input type="text" name="patientName" />
          </label>
          <label>
            Insurance:
            <input type="text" name="insurance" />
          </label>
          <label>
            Procedure:
            <input type="text" name="procedure" />
          </label>
          <label>
            Technician:
            <input type="text" name="technician" />
          </label>
          <label>
            Location:
            <input type="text" name="location" />
          </label>

          <DateWrapper>
            <label>
              Start Date:
              <InputWrapper>
                <input type="date" name="startDate" />
                <input type="time" name="Time"/>
              </InputWrapper>
            </label>
            <label>
            End Date:
              <InputWrapper>
                <input type="date" name="endDate" />
                <input type="time" name="Time"/>
              </InputWrapper>
            </label>
          </DateWrapper>
         
          <div className="actions">
            <button type="submit">Save</button>
            <button type="button" onClick={onClose}>Cancel</button>
          </div>
        </form>
      </ModalContainer>
    </Overlay>
  );
};

export default CreateEditModal;