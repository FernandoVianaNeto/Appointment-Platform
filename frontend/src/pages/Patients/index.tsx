import { useEffect, useState } from 'react';
import Dashboard from '../../core/components/Dashboard';
import Header from '../../core/components/Header';
import HeaderInput from '../../core/components/HeaderInput';
import HeaderSelect from '../../core/components/HeaderSelect';
import SideBar from '../../core/components/SideBar';
import SideBarButton from '../../core/components/SideBarButton';
import { Container, DashboardWrapper, Div, H1, ListOptionsWrapper, Wrapper } from './styles';
import CreationEditButton from '../../core/components/CreationEditButton';
import DashboardList from '../../core/components/DashboardList';
import { useNavigate } from 'react-router-dom';
import ListSummary from '../../core/components/ListSummary';
import { MdDeleteOutline } from "react-icons/md";
import LoadingSpinner from '../../core/components/Loading';
import type { TPatientData } from '../../core/types/patient';
import { createPatient, deletePatients, editPatient, listPatients } from '../../core/services/patientService';
import PatientCard from '../../core/components/PatientCard';
import CreatePatientModal from '../../core/components/CreatePatientModal';
import EditPatientModal from '../../core/components/EditPatientModal';
import ConfirmationModal from '../../core/components/ConfirmationModal';

function Patients() {
  const navigate = useNavigate();
  // const initialLoad = useRef(true);
  const [searchTerm, setSearchTerm] = useState<string>();
  const [totalItems, setTotalItems] = useState(0)
  const [loading, setLoading] = useState(true);
  const [rowSelection, setRowSelection] = useState(false);
  const [allRowsSelected, setAllRowSelected] = useState(false);
  const [createModalOpen, setCreateModalOpen] = useState(false);
  const [editModalOpen, setEditModalOpen] = useState(false);

  const [isConfirmationModalOpened, setIsConfirmationModalOpen] = useState(false);

  const [patientToBeEditted, setPatientToBeEditted] = useState<TPatientData>();
  const [hasMorePatients, setHasMorePatients] = useState(true);
  const [page, setPage] = useState(1);
  const [patients, setPatients] = useState<TPatientData[]>([]);
  const [selectedItems, setSelectedItems] = useState<string[]>([]);

  const handleCreatePatient = async (formData: any) => {
    try {
      await createPatient(formData);
      setCreateModalOpen(false);
      window.location.reload();
    } catch (err) {
      console.error("Error on appointment creation", err);
    }
  };

  const handleEditPatient = async (formData: any) => {
    try {
      console.log("FORM DATA", formData)
      await editPatient(formData);
      setEditModalOpen(false);
      window.location.reload();
    } catch (err) {
      console.error("Error on patient edditing", err);
    }
  };

  const getPatientsListByFilters = async (customPage = page, isNewFilter = false) => {
    setLoading(true);
    try {
      const patientsResponse = await listPatients({ 
        page: customPage, 
        searchTerm, 
      });
  
      if (isNewFilter || customPage === 1) {
        setPatients(patientsResponse.data);
      } else {
        setPatients(prev => {
          const existingIds = new Set(prev.map(item => item.uuid));
          const newItems = patientsResponse.data.filter((item: any) => !existingIds.has(item.uuid));
          return [...prev, ...newItems];
        });
      }
  
      setTotalItems(patientsResponse?.metadata.totalItems ?? 0);
      setPage(patientsResponse?.metadata.next);
      setHasMorePatients(patientsResponse?.metadata.next !== 0);
    } catch (error: any) {
      navigate('/login');
    } finally {
      setLoading(false);
    }
  };
  
  const handleSubmitFilter = async () => {
    try {
      setPatients([]);
      setPage(1);
      await getPatientsListByFilters(1, true);
    } catch (error: any) {
      if (error === "unauthorized") navigate('/');
    }
  };

  const handleGetNextPageAppointments = async () => {
      setLoading(true);
      try {
        getPatientsListByFilters(page, false);
      }
      catch (error: any) {
      if (error === "unauthorized") navigate('/login');
    }
  };

  const handleDeletePatients = async () => {
    setLoading(true);
    try {
      await deletePatients(selectedItems);
      await getPatientsListByFilters();
      setLoading(false);
    }
    catch (error: any) {
      if (error === "unauthorized") navigate('/login');
    }
  };

  useEffect(() => {
    patients.map((item) => {
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
    try {
      async function fetchPatientsList() {
        const patientsResponse = await listPatients({ page: page });
        setPatients(patientsResponse?.data)
        setTotalItems(patientsResponse?.metadata.totalItems ?? 0);
        setLoading(false);
      }
    
      fetchPatientsList();
    } catch (error: any) {
      if (error === "unauthorized") navigate('/');
    }
  }, [])

  return (
    <Container>
      <CreatePatientModal onSave={(e) => handleCreatePatient(e)} isOpen={createModalOpen} onClose={() => setCreateModalOpen(false)} />
      <EditPatientModal onEdit={(e) => handleEditPatient(e)} isOpen={editModalOpen} onClose={() => setEditModalOpen(false)} patient={patientToBeEditted as TPatientData}/>
      <ConfirmationModal 
        confirmationText={selectedItems.length > 1 ? `Are you sure you want to delete the patients?` : `Are you sure you want to delete the patient?` }
        isOpen={isConfirmationModalOpened} 
        onClose={() => setIsConfirmationModalOpen(false)} 
        onConfirm={() => {
          setIsConfirmationModalOpen(false)
          handleDeletePatients()
        }}
      />
      <SideBar>
          <SideBarButton text="Appointments" onClick={() => navigate('/appointments')}/>
          <SideBarButton text="Patients" highlight />
      </SideBar>

      <Dashboard>
        <Header onSubmit={handleSubmitFilter}>
          <HeaderSelect options={['All']} />
          <HeaderInput onChange={(searchTerm) => setSearchTerm(searchTerm)}/>  
        </ Header>

        <Wrapper>
          <H1>Patients</H1>
            <Div>
              <CreationEditButton text="New Patient" highlight onClick={() => setCreateModalOpen(true)}/>
            </Div>
        </Wrapper>
  
        <DashboardWrapper>
          <ListOptionsWrapper deleteSelection={rowSelection}>
            <span>Showing: <p>{loading ? 0 : totalItems} patients</p></span>
            <button type="button" onClick={() => setIsConfirmationModalOpen(true)}>
              <MdDeleteOutline />
            </button>
          </ListOptionsWrapper>
          
          <ListSummary 
            fields={["Patient Name", "Insurance", "Phone", "Email", "Address"]}
            onChange={() => setAllRowSelected(!allRowsSelected)}
          />
          {loading ? <LoadingSpinner /> : (
            <DashboardList 
              noContent={patients?.length === 0}
              hasMore={hasMorePatients}
              fetchMoreData={handleGetNextPageAppointments}
            >
              {
                patients?.length === 0 ?
                <p>No patient registered</p> : 
                  patients?.map((patient: TPatientData) => (
                    <PatientCard
                      key={patient.uuid}
                      uuid={patient.uuid}
                      insurance={patient.insurance}
                      phone={patient.phone}
                      patientName={patient.name}
                      address={patient.address}
                      email={patient.email}
                      rowSelected={selectedItems.includes(patient.uuid)}
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
                      onEdit={() => { 
                        setEditModalOpen(true)
                        setPatientToBeEditted(patient)
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

export default Patients;