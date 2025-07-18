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
import api from '../../core/services/api';
import { listAppointments } from '../../core/services/appointmentsService';
import { useNavigate } from 'react-router-dom';

function Appointments() {
  const navigate = useNavigate();
  const [selected, setSelected] = useState('')
  const [appointmentsCount, setAppointmentsCount] = useState(0)
  const [appointments, setAppointments] = useState()

  const handleSelectChange = (value: string) => {
    setSelected(value)
  };

  useEffect(() => {
    try {
      async function fetchAppointmentsList() {
        const appointments = await listAppointments();
        setAppointments(appointments)
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
          <span>Showing: <p>{appointmentsCount} appointments</p></span>
          <DashboardList>
            {
              appointmentsCount === 0 ?
              <p>No appoitments available</p> : 
              // appointments.map((appointment) => {
                // return 
                <ListCard 
                  key={"1"}
                  endDate='06-07-2025T12:56:45'
                  startDate='06-07-2025T12:56:45'
                  insurance='Private'
                  location='Rua professor gerson pinto, 251'
                  patientName='Fernando Viana'
                  procedure='Remoção capilar'
                  status='Confirmed'
                  technician='Fernando'
                />
              // })
              
            }
          </DashboardList>
        </DashboardWrapper>
        
      </Dashboard>
    </Container>
  );
}

export default Appointments;