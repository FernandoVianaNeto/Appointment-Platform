import { useNavigate } from "react-router-dom";
import type { TPatientData } from "../../types/patient";
import CreationEditButton from "../CreationEditButton";
import { NoPatientsWrapper, Ul } from "./styles";

interface Props {
  results: TPatientData[];
  onSelect: (selectedPatient: TPatientData) => void;
}

function SuggestionDropdown({ results, onSelect }: Props) {
  const navigate = useNavigate()

  return (
    results.length === 0 ? 
        <NoPatientsWrapper>
          <p>No patients found</p>
          <CreationEditButton text="Create new patient" highlight onClick={() => navigate('/patients')}/>
        </NoPatientsWrapper> :
    <Ul>
      {
        results.map((result, index) => (
        <li
          key={result.uuid || index}
          onMouseDown={() => onSelect(result)}
        >
          {result.name}
        </li>
      ))}
    </Ul>
    
  );
}

export default SuggestionDropdown;