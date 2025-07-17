import Dashboard from '../../core/components/Dashboard';
import DashboardHeader from '../../core/components/Dashboard/DashboardHeader';
import SideBar from '../../core/components/SideBar';
import SideBarButton from '../../core/components/SideBarButton';
import { Container } from './styles';

function Appointments() {
  return (
    <Container>
      <SideBar>
          <SideBarButton text="Appointments" highlight/>
          <SideBarButton text="Patients" />
          <SideBarButton text="Settings"/>
      </SideBar>
      <Dashboard>
        <DashboardHeader>
          <select>
            <option>All</option>
          </select>
        </DashboardHeader>
      </Dashboard>
    </Container>
  );
}

export default Appointments;