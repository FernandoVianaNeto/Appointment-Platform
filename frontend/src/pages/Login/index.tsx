import { useState } from 'react';
import { Container, Form, Input, Button, Title } from './styles';
import { login } from '../../core/services/authService';
import { useNavigate } from 'react-router-dom';

function Login() {
    const navigate = useNavigate();
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const [loading, setLoading] = useState(false);

    const handleSubmit = async (e: React.FormEvent) => {
      // e.preventDefault();
      // setLoading(true);
  
      // try {
      //   await login(email, password);
      //   navigate('/home');
      // } catch (err: any) {
      //   // setError(err.message);
      // } finally {
      //   setLoading(false);
      // }
    };
    
    return (
      <Container>
        <Form>
          <Title>Appointment Plataform</Title>
          <Input type="text" placeholder="Email" onChange={e => setEmail(e.target.value)}/>
          <Input type="password" placeholder="Password" onChange={e => setPassword(e.target.value)}/>
          <Button onClick={handleSubmit} disabled={loading}>{loading ? 'Logging in...' : 'Login'}</Button>
        </Form>
      </Container>
    );
}

export default Login;