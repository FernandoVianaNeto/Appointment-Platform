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
import CreateAppointmentModal from '../../core/components/CreateAppointmentModal';
import LoadingSpinner from '../../core/components/Loading';
import { getHours } from '../../core/helpers/getHours';

function Appointments() {
  const navigate = useNavigate();
  const [filterTypeSelected, setFilterTypeSelected] = useState<string>();
  const [searchTerm, setSearchTerm] = useState<string>();
  const [appointments, setAppointments] = useState<TAppointmentResponse>();
  const [totalItems, setTotalItems] = useState(0)
  const [loading, setLoading] = useState(true);
  const [rowSelection, setRowSelection] = useState(false);
  const [createEditModalOpen, setCreateEditModalOpen] = useState(false);

  const handleFilterTypeSelected = (value: string) => {
    setFilterTypeSelected(value)
  };

  const handleCreateAppointment = async (formData: any) => {
    try {
      await createAppointment(formData);
      setCreateEditModalOpen(false);
      window.location.reload();
    } catch (err) {
      console.error("Error on appointment creation", err);
    }
  };

  const handleSubmitFilter = async () => {
    try {
      async function fetchAppointmentsList() {
        const appointments = await listAppointments({ searchTerm, filterType: filterTypeSelected });
        console.log("APPOINTMENTS", appointments);
        setAppointments(appointments)
        setTotalItems(appointments?.metadata.totalItems ?? 0)
        setLoading(false)
      }
    
      fetchAppointmentsList();
    } catch (error: any) {
      if (error === "unauthorized")
      navigate('/')
    }
  }

  useEffect(() => {
    console.log('RELOADED')
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
      <CreateAppointmentModal onSave={(e) => handleCreateAppointment(e)} isOpen={createEditModalOpen} onClose={() => setCreateEditModalOpen(false)} />
      <SideBar>
          <SideBarButton text="Appointments" highlight/>
          <SideBarButton text="Patients" />
          <SideBarButton text="Settings"/>
      </SideBar>

      <Dashboard>
        <Header onSubmit={handleSubmitFilter}>
          <HeaderSelect options={['All', 'Patient', 'Procedure']} onChange={handleFilterTypeSelected}/>
          <HeaderInput onChange={(searchTerm) => setSearchTerm(searchTerm)}/>  
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
          
          <ListSummary fields={["Time", "Patient Name", "Insurance", "Procedure", "Technician", "Location", "Status"]} onChange={() => setRowSelection(!rowSelection)}/>
          {loading ? <LoadingSpinner /> : (
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