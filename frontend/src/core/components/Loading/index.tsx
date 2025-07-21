import { Spinner, SpinnerContainer } from "./styles";



interface LoadingSpinnerProps {
  message?: string;
}

function LoadingSpinner({ message }: LoadingSpinnerProps) {
  return (
    <SpinnerContainer>
      <Spinner />
      {message && <span style={{ marginLeft: 12 }}>{message}</span>}
    </SpinnerContainer>
  );
}

export default LoadingSpinner;
