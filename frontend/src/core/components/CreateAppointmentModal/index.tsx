import { useEffect, useState } from "react";
import { DateWrapper, InputWrapper, ModalContainer, Overlay } from "./styles";
import { listPatients } from "../../services/patientService";
import type { TPatientData } from "../../types/patient";

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSave: (data: any) => void;
}

function CreateAppointmentModal ({ isOpen, onClose, onSave }: ModalProps) {
  if (!isOpen) return null;
  const [name, setName] = useState("");
  const [show, setShow] = useState(false);
  const [recommendedPatients, setRecommendedPatients] = useState<TPatientData[]>([])

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const form = e.currentTarget;
    const data = {
      patientName: form.patientName.value,
      insurance: form.insurance.value,
      procedure: form.procedure.value,
      technician: form.technician.value,
      location: form.location.value,
      startDate: `${form.startDate.value}T${form.startTime.value}`,
      endDate: `${form.endDate.value}T${form.endTime.value}`,
    };

    onSave(data);
  };

  useEffect(() => {
    const timeoutId = setTimeout(() => {
      async function getPatientsByName() {
        if (name.trim() === "") {
          return;
        }
        const result = await listPatients(name);
        console.log("RESULT", result);
        setRecommendedPatients(result.data);
      }
  
      getPatientsByName();
    }, 300);
  
    return () => clearTimeout(timeoutId);
  }, [name]);

  return (
    <Overlay>
      <ModalContainer>
        <h2>Create New Appointment</h2>
        <form method="post" onSubmit={handleSubmit}>
          <div style={{ position: 'relative' }}>
            <label>
              Patient Name:
              <input
                type="text"
                name="patientName"
                value={name}
                onChange={(e) => {
                  setName(e.target.value);
                  setShow(true);
                }}
                onBlur={(e) => {
                  if (!e.currentTarget.contains(e.relatedTarget)) {
                    setTimeout(() => setShow(false), 150);
                  }
                }}
                onFocus={() => name && setShow(true)}
              />
              {show && recommendedPatients?.length > 0 && (
                <ul style={{
                  marginTop: 4,
                  padding: 0,
                  listStyle: 'none',
                  border: '1px solid #ccc',
                  borderRadius: 8,
                  maxHeight: 150,
                  overflowY: 'auto',
                  boxShadow: '0 4px 10px rgba(0, 0, 0, 0.1)',
                  position: 'absolute',
                  backgroundColor: 'white',
                  zIndex: 1000,
                  width: '100%'
                }}>
                  {recommendedPatients.map((patient, index) => (
                    <li
                      key={patient.uuid || index}
                      style={{ padding: '10px', cursor: 'pointer' }}
                      onMouseDown={() => {
                        setName(patient.name);
                        setShow(false);
                      }}
                    >
                      {patient.name}
                    </li>
                  ))}
                </ul>
              )}
            </label>
          </div>
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
                <input type="time" name="startTime"/>
              </InputWrapper>
            </label>
            <label>
            End Date:
              <InputWrapper>
                <input type="date" name="endDate" />
                <input type="time" name="endTime"/>
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

export default CreateAppointmentModal;