import { Container, Form, Input, Button, Title } from './styles';

function Login() {
  return (
    <Container>
      <Form>
        <Title>Login</Title>
        <Input type="text" placeholder="Email" />
        <Input type="password" placeholder="Senha" />
        <Button>Entrar</Button>
      </Form>
    </Container>
  );
}

export default Login;