import { useEffect, useState } from 'react';
import Dashboard from '../../core/components/Dashboard';
import Header from '../../core/components/Header';
import HeaderInput from '../../core/components/HeaderInput';
import HeaderSelect from '../../core/components/HeaderSelect';
import SideBar from '../../core/components/SideBar';
import SideBarButton from '../../core/components/SideBarButton';
import { Container, DashboardWrapper, Div, H1, ListOptionsWrapper, Wrapper } from './styles';
import DateSelector from '../../core/components/DateSelector';
import CreationEditButton from '../../core/components/CreationEditButton';
import DashboardList from '../../core/components/DashboardList';
import ListCard from '../../core/components/ListCard';
import { createAppointment, listAppointments } from '../../core/services/appointmentsService';
import { useNavigate } from 'react-router-dom';
import type { TAppointmentItem, TAppointmentResponse } from '../../core/types/appointments';
import ListSummary from '../../core/components/ListSummary';
import { MdDeleteOutline } from "react-icons/md";
import CreateEditModal from '../../core/components/CreateEditModal';

function Appointments() {
  const navigate = useNavigate();
  const [_, setSelected] = useState('');
  const [appointments, setAppointments] = useState<TAppointmentResponse>();
  const [totalItems, setTotalItems] = useState(0)
  const [loading, setLoading] = useState(true);
  const [rowSelection, setRowSelection] = useState(false);
  const [createEditModalOpen, setCreateEditModalOpen] = useState(false);

  const handleSelectChange = (value: string) => {
    setSelected(value)
  };

  const getHours = (stringDate: string) => {
    const date = new Date(stringDate);

    const hours = date.getHours().toString().padStart(2, '0');
    const minutes = date.getMinutes().toString().padStart(2, '0');

    return `${hours}:${minutes}`
  }

  const handleCreateAppointment = async (formData: any) => {
    try {
      await createAppointment(formData);
      setCreateEditModalOpen(false);
      window.location.reload();
    } catch (err) {
      console.error("Error on appointment creation", err);
    }
  };


  useEffect(() => {
    try {
      async function fetchAppointmentsList() {
        const appointments = await listAppointments();
        setAppointments(appointments)
        setTotalItems(appointments?.metadata.totalItems ?? 0)
        setLoading(false)
      }
    
      fetchAppointmentsList();
    } catch (error: any) {
      if (error === "unauthorized")
      navigate('/')
    }
  }, [])

  return (
    <Container>
      <CreateEditModal onSave={(e) => handleCreateAppointment(e)} isOpen={createEditModalOpen} onClose={() => setCreateEditModalOpen(false)} />
      <SideBar>
          <SideBarButton text="Appointments" highlight/>
          <SideBarButton text="Patients" />
          <SideBarButton text="Settings"/>
      </SideBar>

      <Dashboard>

        <Header>
          <HeaderSelect options={['All', 'Patient', 'Procedure']} onChange={handleSelectChange}/>
          <HeaderInput />  
        </ Header>

        <Wrapper>
          <H1>Appointments</H1>
            <Div>
              <DateSelector />
              <CreationEditButton text="New Appointment" highlight onClick={() => setCreateEditModalOpen(true)}/>
            </Div>
        </Wrapper>
        
        <DashboardWrapper>

          <ListOptionsWrapper deleteSelection={rowSelection}>
            <span>Showing: <p>{loading ? 0 : totalItems} appointments</p></span>
            <button>
              <MdDeleteOutline />
            </button>
          </ListOptionsWrapper>
          
          <ListSummary fields={["Time", "Patiet Name", "Insurance", "Procedure", "Technician", "Location", "Status"]} onChange={() => setRowSelection(!rowSelection)}/>
          {loading ? <p>Loading...</p> : (
            <DashboardList>
              {
                appointments?.data.length === 0 ?
                <p>No appointments available</p> : 
                  appointments?.data.map((appointment: TAppointmentItem) => (
                    <ListCard 
                      key={appointment.uuid}
                      endDate={getHours(appointment.end_date)}
                      startDate={getHours(appointment.start_date)}
                      insurance={appointment.patient.insurance}
                      location={appointment.location}
                      patientName={appointment.patient.name}
                      procedure={appointment.procedure}
                      status={appointment.status}
                      technician="Fernando"
                      rowSelected={rowSelection}
                    />
                  ))
              }
            </DashboardList>
          )}

        </DashboardWrapper>
      </Dashboard>
    </Container>
  );
}

export default Appointments;