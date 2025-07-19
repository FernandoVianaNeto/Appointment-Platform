interface Props {
  results: { name: string; uuid: string }[];
  onSelect: (name: string) => void;
}

export function PatientSuggestionsDropdown({ results, onSelect }: Props) {
  return (
    <ul style={{
      marginTop: 4,
      padding: 0,
      listStyle: 'none',
      border: '1px solid #ccc',
      borderRadius: 8,
      maxHeight: 150,
      overflowY: 'auto',
      boxShadow: '0 4px 10px rgba(0, 0, 0, 0.1)',
      position: 'absolute',
      backgroundColor: 'white',
      zIndex: 1000,
      width: '100%'
    }}>
      {results.map((patient, index) => (
        <li
          key={patient.uuid || index}
          style={{ padding: '10px', cursor: 'pointer' }}
          onMouseDown={() => onSelect(patient.name)}
        >
          {patient.name}
        </li>
      ))}
    </ul>
  );
}
