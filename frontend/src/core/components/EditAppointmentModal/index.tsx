import { useEffect, useState } from "react";
import { DateWrapper, InputWrapper, ModalContainer, Overlay } from "./styles";
import { listPatients } from "../../services/patientService";
import type { TPatientData } from "../../types/patient";
import SuggestionDropdown from "../SuggestionDropdown";
import { addMinutesToTime } from "../../helpers/addMinutesToTime";
import type { TAppointmentItem } from "../../types/appointments";

interface ModalProps {
  isOpen: boolean;
  appointment: TAppointmentItem
  onClose: () => void;
  onEdit: (data: any) => void;
}

function EditAppointmentModal ({ isOpen, onClose, onEdit, appointment }: ModalProps) {
  if (!isOpen) return null;
  const [name, setName] = useState<string>(appointment.patient.name);
  const [procedure, setProcedure] = useState<string>(appointment.procedure);
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
        uuid: appointment.uuid,
        patient_uuid: selectedPatient?.uuid,
        insurance: appointment.patient.insurance,
        procedure: procedure,
        technician: "Fernando",
        location: appointment.location,
        start_date: `${form.startDate.value}T${form.startTime.value}`,
        end_date: `${form.startDate.value}T${endTime}`,
    };

    onEdit(data);
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
            <input className="insurance-input" type="text" name="insurance" value={appointment.patient.insurance} disabled />
          </label>

          <label>
            Procedure*:
            <input className="procedure-input" type="text" name="procedure" value={procedure} onChange={(e) => setProcedure(e.target.value)}/>
          </label>

          <label>
            Technician*:
            <input className="technician-input" type="text" name="technician" value="Fernando" disabled/>
          </label>

          <label>
            Location*:
            <input className="location-input" type="text" name="location" value={appointment.location} disabled />
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
            <button type="submit">Edit</button>
            <button type="button" onClick={onClose}>Cancel</button>
          </div>

        </form>
      </ModalContainer>
    </Overlay>
  );
};

export default EditAppointmentModal;