import { useEffect, useState } from 'react';
import Dashboard from '../../core/components/Dashboard';
import Header from '../../core/components/Header';
import HeaderInput from '../../core/components/HeaderInput';
import HeaderSelect from '../../core/components/HeaderSelect';
import SideBar from '../../core/components/SideBar';
import SideBarButton from '../../core/components/SideBarButton';
import { Container, DashboardWrapper, Div, H1, Wrapper } from './styles';
import DateSelector from '../../core/components/DateSelector';
import CreationEditButton from '../../core/components/CreationEditButton';
import DashboardList from '../../core/components/DashboardList';
import ListCard from '../../core/components/ListCard';
import { listAppointments } from '../../core/services/appointmentsService';
import { useNavigate } from 'react-router-dom';
import type { TAppointmentItem, TAppointmentResponse } from '../../core/types/appointments';
import ListSummary from '../../core/components/ListSummary';
import dayjs from 'dayjs';

function Appointments() {
  const navigate = useNavigate();
  const [selected, setSelected] = useState('');
  const [appointments, setAppointments] = useState<TAppointmentResponse>();
  const [totalItems, setTotalItems] = useState(0)
  const [loading, setLoading] = useState(true);
  
  const handleSelectChange = (value: string) => {
    setSelected(value)
  };

  const getHours = (stringDate: string) => {
    const date = new Date(stringDate);

    const hours = date.getHours().toString().padStart(2, '0');
    const minutes = date.getMinutes().toString().padStart(2, '0');

    return `${hours}:${minutes}`
  }

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
              <CreationEditButton text="New Appointment" highlight/>
            </Div>
        </Wrapper>
        <DashboardWrapper>
          <span>Showing: <p>{loading ? 0 : totalItems} appointments</p></span>
          <ListSummary fields={["Time", "Patiet Name", "Insurance", "Procedure", "Technician", "Location", "Status"]}/>
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