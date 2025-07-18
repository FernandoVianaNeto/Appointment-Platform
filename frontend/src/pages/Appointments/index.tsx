import { useState } from 'react';
import Dashboard from '../../core/components/Dashboard';
import Header from '../../core/components/Header';
import HeaderInput from '../../core/components/HeaderInput';
import HeaderOption from '../../core/components/HeaderOption';
import HeaderSelect from '../../core/components/HeaderSelect';
import SideBar from '../../core/components/SideBar';
import SideBarButton from '../../core/components/SideBarButton';
import { Container } from './styles';

function Appointments() {
  const [selected, setSelected] = useState(false)

  return (
    <Container>
      <SideBar>
          <SideBarButton text="Appointments" highlight/>
          <SideBarButton text="Patients" />
          <SideBarButton text="Settings"/>
      </SideBar>
      <Dashboard>
        <Header>
          <HeaderSelect>
            <HeaderOption selected={selected}>
              All
            </HeaderOption>
            <HeaderOption selected={selected}>
              All
            </HeaderOption>
            <HeaderOption selected={selected}>
              All
            </HeaderOption>
          </HeaderSelect>
          <HeaderInput />  
        </ Header>
      </Dashboard>
    </Container>
  );
}

export default Appointments;