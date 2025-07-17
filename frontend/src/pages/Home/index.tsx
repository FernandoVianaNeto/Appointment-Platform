import Dashboard from '../../core/components/Dashboard';
import SideBar from '../../core/components/SideBar';
import SideBarButton from '../../core/components/SideBarButton';
import { Container } from './styles';

function Home() {
  return (
    <Container>
      <SideBar>
          <SideBarButton text="Appointments" highlight/>
          <SideBarButton text="Patients" />
          <SideBarButton text="Appointments"/>
      </SideBar>
      <Dashboard>

      </Dashboard>
    </Container>
  );
}

export default Home;