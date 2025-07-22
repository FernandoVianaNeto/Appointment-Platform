import { useEffect, useState } from "react";
import { DateWrapper, InputWrapper, ModalContainer, Overlay } from "./styles";
import { listPatients } from "../../services/patientService";
import type { TPatientData } from "../../types/patient";
import SuggestionDropdown from "../SuggestionDropdown";
import { addMinutesToTime } from "../../helpers/addMinutesToTime";

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSave: (data: any) => void;
}

function CreateAppointmentModal ({ isOpen, onClose, onSave }: ModalProps) {
  if (!isOpen) return null;
  const [name, setName] = useState<string>("");
  const [show, setShow] = useState<boolean>(false);
  const [recommendedPatients, setRecommendedPatients] = useState<TPatientData[]>([])
  const [selectedPatient, setSelectedPatient] = useState<TPatientData>()

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const form = e.currentTarget;

    if (selectedPatient === undefined || 
      form.insurance.value === undefined || 
      form.procedure.value === undefined || 
      form.technician.value === undefined || 
      form.startDate.value === undefined ||
      form.startTime.value == undefined ||
      form.procedureDuration.value == undefined
    ) {
      throw new Error('Could not save appointment. Missing Required fields')
    }

    const endTime = addMinutesToTime(form.startTime.value, form.procedureDuration.value)

    const data = {
      patient_uuid: selectedPatient?.uuid,
      insurance: form.insurance.value,
      procedure: form.procedure.value,
      technician: form.technician.value,
      location: form.location.value,
      start_date: `${form.startDate.value}T${form.startTime.value}`,
      end_date: `${form.startDate.value}T${endTime}`,
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
              Patient Name*:
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
                className="patient-input"
              />
            </label>
            {show && recommendedPatients?.length > 0 && (
              <SuggestionDropdown
                results={recommendedPatients}
                onSelect={(selectedName) => {
                  setSelectedPatient(selectedName);
                  setName(selectedName.name)
                  setShow(false);
                }}
              />
            )}
          </div>

          <label>
            Insurance*:
            <input className="insurance-input" type="text" name="insurance" />
          </label>

          <label>
            Procedure*:
            <input className="procedure-input" type="text" name="procedure" />
          </label>

          <label>
            Technician*:
            <input className="technician-input" type="text" name="technician" />
          </label>

          <label>
            Location*:
            <input className="location-input" type="text" name="location" />
          </label>

          <DateWrapper>
            <label>
              Start Date*:
              <InputWrapper>
                <input className="start-date-input" type="date" name="startDate" />
                <input className="start-date-time-input" type="time" name="startTime" />
                <input className="procedure-duration-input" type="number" name="procedureDuration" placeholder="Duration in minutes" />
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