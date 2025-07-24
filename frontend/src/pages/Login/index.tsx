import { useState } from 'react';
import { Container, Form, Input, Button, Title, ButtonWrapper } from './styles';
import { useNavigate } from 'react-router-dom';
import { login } from '../../core/services/authService';
import SignUpModal from '../../core/components/SignUpModal';
import { createClinic } from '../../core/services/clinicService';

function Login() {
  const navigate = useNavigate();
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [error, setError] = useState<boolean>(false);

  const [isSignUpModalOpen, setIsSignUpModalOpen] = useState<boolean>(false);

  const handleSignUp = async (data: any) => {
    await createClinic({ email: data.email, password: data.password, name: data.clinicName})
    setIsSignUpModalOpen(false);
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
  
    try {
      await login(email, password);
      navigate('/appointments');
    } catch (err) {
      if (err instanceof Error && err.message === 'Invalid email or password') {
        setError(true);
      } else {
        alert('Unexpected error occurred. Please try again later.');
      }
    }
  };

  return (
    <Container>
      <SignUpModal isOpen={isSignUpModalOpen} onClose={() => setIsSignUpModalOpen(false)} onSave={(e) => handleSignUp(e)}/>
      <Form onSubmit={handleSubmit}>
        <Title>Appointment Platform</Title>
        <Input type="text" placeholder="Email" onChange={e => setEmail(e.target.value)} />
        <Input type="password" placeholder="Password" onChange={e => setPassword(e.target.value)} />
        {error && <p>Incorrect Email or Password</p>}
        <ButtonWrapper>
          <Button type="submit" highlight>Sign in</Button>
          <Button type="button" onClick={() => setIsSignUpModalOpen(true)}>Sign up</Button>
          <Button type="button">Forgot Password</Button>
        </ButtonWrapper>
      </Form>
    </Container>
  );
}

export default Login;