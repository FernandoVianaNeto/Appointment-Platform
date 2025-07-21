import { useState } from 'react';
import { Container, Form, Input, Button, Title } from './styles';
import { useNavigate } from 'react-router-dom';
import { login } from '../../core/services/authService';
import LoadingSpinner from '../../core/components/Loading';

function Login() {
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState<boolean>(false);
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    console.log("before preventDefault", e);
    e.preventDefault();
    console.log("after preventDefault");
    setLoading(true);
  
    try {
      await login(email, password);
      navigate('/appointments');
    } catch (err: any) {
      console.error('Login error:', err);
  
      if (err instanceof Error && err.message === 'Invalid email or password') {
        console.log(err)
        setError(true);
      } else {
        alert('Unexpected error occurred. Please try again later.');
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <Container>
      <Form onSubmit={handleSubmit}>
        <Title>Appointment Platform</Title>
        <Input type="text" placeholder="Email" onChange={e => setEmail(e.target.value)} />
        <Input type="password" placeholder="Password" onChange={e => setPassword(e.target.value)} />
        <Button type="submit">Login</Button>
        {error && <p>Incorrect Email or Password</p>}
      </Form>
    </Container>
  );
}

export default Login;