import { useEffect, useState } from "react";
import { DateWrapper, Input, InputWrapper, ModalContainer, Overlay } from "./styles";
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
  const [isMissingRequiredFields, setIsMissingRequiredFields] = useState<boolean>(false);
  const [isMissingName, setIsMissingName] = useState<boolean>(false);
  const [isMissingInsurance, setIsMissingInsurance] = useState<boolean>(false);
  const [isMissingProcedure, setIsMissingProcedure] = useState<boolean>(false);
  const [isMissingTechnician, setIsMissingTechnician] = useState<boolean>(false);
  const [isMissingStartDate, setIsMissingStartDate] = useState<boolean>(false);
  const [isMissingLocation, setIsMissingLocation] = useState<boolean>(false);
  const [isMissingStartTime, setIsMissingStartTime] = useState<boolean>(false);
  const [isMissingProcedureDuration, setIsMissingProcedureDuration] = useState<boolean>(false);

  const [recommendedPatients, setRecommendedPatients] = useState<TPatientData[]>([])
  const [selectedPatient, setSelectedPatient] = useState<TPatientData>()


  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const form = e.currentTarget;

    console.log(form.insurance.value === '', isMissingInsurance, isMissingLocation, isMissingProcedure, isMissingProcedureDuration, isMissingStartDate, isMissingStartTime);

    if (selectedPatient === undefined) {
      setIsMissingRequiredFields(true);
      setIsMissingName(true);
    }

    if (form.insurance.value === '') {
      setIsMissingInsurance(true); 
      setIsMissingRequiredFields(true);
    }

    if (form.procedure.value === '') {
      setIsMissingProcedure(true); 
      setIsMissingRequiredFields(true);
    }

    if (form.technician.value === '') {
      setIsMissingTechnician(true); 
      setIsMissingRequiredFields(true);
    }

    if (form.startDate.value === '') {
      setIsMissingStartDate(true); 
      setIsMissingRequiredFields(true);
    }

    if (form.startTime.value === '') {
      setIsMissingStartTime(true); 
      setIsMissingRequiredFields(true);
    }

    if (form.procedureDuration.value === '') {
      setIsMissingProcedureDuration(true); 
      setIsMissingRequiredFields(true);
    }

    if (form.location.value === '') {
      setIsMissingLocation(true); 
      setIsMissingRequiredFields(true);
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
        const result = await listPatients({ searchTerm: name, page: 1 });
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
              <Input
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
                missingField={isMissingName}
              />
            </label>
            {show && (
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
            <Input className="insurance-input" type="text" name="insurance" missingField={isMissingInsurance}/>
          </label>

          <label>
            Procedure*:
            <Input className="procedure-input" type="text" name="procedure" missingField={isMissingProcedure}/>
          </label>

          <label>
            Technician*:
            <Input className="technician-input" type="text" name="technician" missingField={isMissingTechnician}/>
          </label>

          <label>
            Location*:
            <Input className="location-input" type="text" name="location" missingField={isMissingLocation}/>
          </label>

          <DateWrapper>
            <label>
              Start Date*:
              <InputWrapper>
                <Input className="start-date-input" type="date" name="startDate" missingField={isMissingStartDate}/>
                <Input className="start-date-time-input" type="time" name="startTime" missingField={isMissingStartTime}/>
                <Input className="procedure-duration-input" type="number" name="procedureDuration" placeholder="Duration in minutes" missingField={isMissingProcedureDuration}/>
              </InputWrapper>
            </label>
          </DateWrapper>
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

export default CreateAppointmentModal;