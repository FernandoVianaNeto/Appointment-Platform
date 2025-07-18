import { useState } from 'react';
import Dashboard from '../../core/components/Dashboard';
import Header from '../../core/components/Header';
import HeaderInput from '../../core/components/HeaderInput';
import HeaderSelect from '../../core/components/HeaderSelect';
import SideBar from '../../core/components/SideBar';
import SideBarButton from '../../core/components/SideBarButton';
import { Container, Div, H1, Wrapper } from './styles';
import DateSelector from '../../core/components/DateSelector';

function Appointments() {
  const [selected, setSelected] = useState('')

  const handleSelectChange = (value: string) => {
    setSelected(value)
  };

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
            </Div>
        </Wrapper>
      </Dashboard>
    </Container>
  );
}

export default Appointments;