import type { TPatientData } from "../../types/patient";
import { Ul } from "./styles";

interface Props {
  results: TPatientData[];
  onSelect: (selectedPatient: TPatientData) => void;
}

function SuggestionDropdown({ results, onSelect }: Props) {
  return (
    <Ul>
      {results.map((result, index) => (
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