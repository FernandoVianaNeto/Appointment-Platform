import { useEffect, useRef, useState } from 'react';
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
import { createAppointment, deleteAppointments, listAppointments } from '../../core/services/appointmentsService';
import { useNavigate } from 'react-router-dom';
import type { TAppointmentItem } from '../../core/types/appointments';
import ListSummary from '../../core/components/ListSummary';
import { MdDeleteOutline } from "react-icons/md";
import CreateAppointmentModal from '../../core/components/CreateAppointmentModal';
import LoadingSpinner from '../../core/components/Loading';
import { getHours } from '../../core/helpers/getHours';
import dayjs from 'dayjs';

function Appointments() {
  const navigate = useNavigate();
  const initialLoad = useRef(true);
  const [filterTypeSelected, setFilterTypeSelected] = useState<string>();
  const [filterDate, setFilterDate] = useState<string>();
  const [searchTerm, setSearchTerm] = useState<string>();
  const [totalItems, setTotalItems] = useState(0)
  const [loading, setLoading] = useState(true);
  const [rowSelection, setRowSelection] = useState(false);
  const [allRowsSelected, setAllRowSelected] = useState(false);
  const [createEditModalOpen, setCreateEditModalOpen] = useState(false);
  const [hasMoreAppointments, setHasMoreAppointments] = useState(true);
  const [page, setPage] = useState(1);
  const [appointments, setAppointments] = useState<TAppointmentItem[]>([]);
  const [selectedItems, setSelectedItems] = useState<string[]>([]);

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

  const getAppointmentListByFilters = async (customPage = page, isNewFilter = false) => {
    setLoading(true);
    try {
      const appointmentsResponse = await listAppointments({ 
        page: customPage, 
        searchTerm, 
        filterType: filterTypeSelected, 
        date: filterDate 
      });
  
      if (isNewFilter || customPage === 1) {
        setAppointments(appointmentsResponse.data);
      } else {
        setAppointments(prev => {
          const existingIds = new Set(prev.map(item => item.uuid));
          const newItems = appointmentsResponse.data.filter((item: any) => !existingIds.has(item.uuid));
          return [...prev, ...newItems];
        });
      }
  
      setTotalItems(appointmentsResponse?.metadata.totalItems ?? 0);
      setPage(appointmentsResponse?.metadata.next);
      setHasMoreAppointments(appointmentsResponse?.metadata.next !== 0);
    } catch (error: any) {
      navigate('/login');
    } finally {
      setLoading(false);
    }
  };
  
  const handleSubmitFilter = async () => {
    try {
      setAppointments([]);
      setPage(1);
      await getAppointmentListByFilters(1, true);
    } catch (error: any) {
      if (error === "unauthorized") navigate('/');
    }
  };

  const handleGetNextPageAppointments = async () => {
      setLoading(true);
      try {
        getAppointmentListByFilters(page, false);
      }
      catch (error: any) {
      if (error === "unauthorized") navigate('/login');
    }
  };

  const handleDeleteAppointments = async () => {
    try {
      console.log(selectedItems)
      await deleteAppointments(selectedItems);
    }
    catch (error: any) {
      if (error === "unauthorized") navigate('/login');
    }
  };

  useEffect(() => {
    appointments.map((item) => {
      if (allRowsSelected) {
        setRowSelection(true)
        setSelectedItems((prev) => [...prev, item.uuid])
      } else {
        setRowSelection(false)
        setSelectedItems([])
      }
    })
  }, [allRowsSelected])

  useEffect(() => {
    const now = dayjs().format('YYYY-MM-DD');
    setFilterDate(now)
    try {
      async function fetchAppointmentsList() {
        const appointmentsResponse = await listAppointments({ page: page, date: now });
        setTotalItems(appointmentsResponse?.metadata.totalItems ?? 0);
        setLoading(false);
      }
    
      fetchAppointmentsList();
    } catch (error: any) {
      if (error === "unauthorized") navigate('/');
    }
  }, [])

  useEffect(() => {
    if (initialLoad.current) {
      initialLoad.current = false;
      return;
    }
  
    const fetchFilteredAppointments = async () => {
      try {
        setLoading(true);
        setAppointments([]);
        setPage(1);
        await getAppointmentListByFilters(1);
      } catch (error: any) {
        if (error === "unauthorized") navigate('/');
      }
    };
  
    fetchFilteredAppointments();
  }, [filterDate]);

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
              <DateSelector onClick={(selectedDate) => setFilterDate(selectedDate.format('YYYY-MM-DD'))}/>
              <CreationEditButton text="New Appointment" highlight onClick={() => setCreateEditModalOpen(true)}/>
            </Div>
        </Wrapper>
  
        <DashboardWrapper>
          <ListOptionsWrapper deleteSelection={rowSelection}>
            <span>Showing: <p>{loading ? 0 : totalItems} appointments</p></span>
            <button type="button" onClick={handleDeleteAppointments}>
              <MdDeleteOutline />
            </button>
          </ListOptionsWrapper>
          
          <ListSummary 
            fields={["Time", "Patient Name", "Insurance", "Procedure", "Technician", "Location", "Status"]}
            onChange={() => setAllRowSelected(!allRowsSelected)}
          />
          {loading ? <LoadingSpinner /> : (
            <DashboardList 
              noContent={appointments?.length === 0}
              hasMore={hasMoreAppointments}
              fetchMoreData={handleGetNextPageAppointments}
            >
              {
                appointments?.length === 0 ?
                <p>No appointments available</p> : 
                  appointments?.map((appointment: TAppointmentItem) => (
                    <ListCard 
                      key={appointment.uuid}
                      uuid={appointment.uuid}
                      endDate={getHours(appointment.end_date)}
                      startDate={getHours(appointment.start_date)}
                      insurance={appointment.patient.insurance}
                      location={appointment.location}
                      patientName={appointment.patient.name}
                      procedure={appointment.procedure}
                      status={appointment.status}
                      technician="Fernando"
                      rowSelected={selectedItems.includes(appointment.uuid)}
                      onRowSelected={(uuid) => {
                        setRowSelection(!rowSelection)
                        setSelectedItems(prev => {
                          if (prev.includes(uuid)) {
                            return prev.filter(id => id !== uuid);
                          } else {
                            return [...prev, uuid];
                          }
                        });
                      }}
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