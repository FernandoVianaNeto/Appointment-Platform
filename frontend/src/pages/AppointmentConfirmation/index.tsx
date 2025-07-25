import { ModalContainer, ModalWrapper, Overlay } from "./styles";
import { updateAppointmentStatus } from "../../core/services/appointmentsService";
import { useState } from "react";

function AppointmentConfirmation() {
    const [operationFinished, setOperationFinished] = useState<boolean>(false);

  const queryParams = new URLSearchParams(window.location.search);
  const appointmentUuid = queryParams.get("uuid");

  if (appointmentUuid === '' || appointmentUuid === undefined || appointmentUuid === null) {
    return (
        <Overlay>
            <ModalContainer>
                <ModalWrapper>
                    <h2>No appointment found</h2>
                </ModalWrapper>
            </ModalContainer>
        </Overlay>
    )
  }

  async function setAppointmentStatus(status: string) {
    try {
      await updateAppointmentStatus(appointmentUuid as string, status);
      setOperationFinished(true);
    } catch {
       console.log("Could not confirm or cancel the appointment") 
    }
  }

  return (
    <Overlay>
        {
            operationFinished ? 
            <ModalContainer>
                <ModalWrapper>
                    <h2>Thank you for the confirmation</h2>
                </ModalWrapper>
            </ModalContainer> :
            <ModalContainer>
                <ModalWrapper>
                    <h2>Would you like to confirm your appointment?</h2>
                    <div className="actions">
                        <button type="button" className="confirmed" onClick={() => setAppointmentStatus('confirmed')}>Confirm</button>
                        <button type="button" className="canceled" onClick={() => setAppointmentStatus('canceled')}>Cancel</button>
                    </div>
                </ModalWrapper>
            </ModalContainer>
        }
        
    </Overlay>
  );
}

export default AppointmentConfirmation;
