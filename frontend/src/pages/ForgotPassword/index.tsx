import { useState } from "react";
import { ErrorText, Input, ModalContainer, ModalWrapper, Overlay } from "./styles";
import { generateResetPasswordCode, resetPasswordCall, validateResetPasswordCode } from "../../core/services/authService";
import { useNavigate } from "react-router-dom";

function ForgotPassword() {
  const navigate = useNavigate();
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [confirmPassword, setConfirmPassword] = useState<string>('');

  const [resetPasswordCode, setResetPasswordCode] = useState<number>(0);

  const [generateResetCodeError, setGenerateResetCodeError] = useState<boolean>(false);
  const [validationError, setValidationError] = useState<boolean>(false);
  const [resetPasswordError, setResetPasswordError] = useState<boolean>(false);
  const [differentPasswords, setDifferentPasswords] = useState<boolean>(false);
  const [passwordReseted, setPasswordReseted] = useState<boolean>(false);
  const [emailSent, setEmailSent] = useState<boolean>(true);
  const [validateCode, setValidateCode] = useState<boolean>(false);
  const [resetPassword, setResetPassword] = useState<boolean>(false);

  async function handleGeneratePasswordCode() {
    try {
      await generateResetPasswordCode(email);
      setEmailSent(false);
      setValidateCode(true);
    } catch {
      setGenerateResetCodeError(true);
    }
  }

  async function handleValidateResetPasswordCode() {
    try {
      await validateResetPasswordCode(email, resetPasswordCode);
      setValidateCode(false);
      setResetPassword(true);
    } catch {
      setValidationError(true);
    }
  }

  async function handleResetPassword() {
    try {
      if (password !== confirmPassword) {
        setDifferentPasswords(true);
        return;
      }

      await resetPasswordCall(email, resetPasswordCode, password);
      setPasswordReseted(true);
      setResetPassword(false);
    } catch {
      setResetPasswordError(true);
    }
  }


  return (
    <Overlay>
      {
        emailSent && 
        <ModalContainer>
          <ModalWrapper>
            <h2>Forgot your password?</h2>
            <p id="subtitle">Please enter your registered email address below. We’ll send you a reset code.</p>

            <label htmlFor="email">Email address*</label>
            <Input
              id="email"
              type="email"
              name="email"
              placeholder="you@example.com"
              required
              onChange={e => setEmail(e.target.value)}
            />
            {
              generateResetCodeError && <ErrorText>Error on reset code generation</ErrorText>
            }

            <div className="actions">
              <button type="submit" onClick={handleGeneratePasswordCode} className="primary">Send code</button>
              <button type="button" onClick={() => navigate('/login')} className="secondary">Cancel</button>
            </div>
          </ModalWrapper>
        </ModalContainer>
      }
      {
        validateCode && 
        <ModalContainer>
          <h2>Type bellow the code you received in your registered email</h2>
          <ModalWrapper>
            <label htmlFor="email">Code:</label>
            <Input
              id="reset-password-code"
              type="number"
              name="reset-password-code"
              placeholder="123456"
              required
              onChange={(e) => setResetPasswordCode(Number(e.target.value))}
            />
            {
              validationError && <ErrorText>Invalid code</ErrorText>
            }

            <div className="actions">
              <button type="button" onClick={handleValidateResetPasswordCode} className="primary">Validate code</button>
              <button type="button" onClick={() => navigate('/login')} className="secondary">Cancel</button>
            </div>
          </ModalWrapper>
        </ModalContainer>
      }
      {
        resetPassword && 
        <ModalContainer>
          <h2>Reset your password</h2>
          <ModalWrapper>
            <label htmlFor="email">New Password</label>
            <Input
              id="new-password"
              type="password"
              name="new-password"
              placeholder="*******"
              required
              onChange={(e) => setPassword(e.target.value)}
            />

            <label htmlFor="email">Confirm new password</label>
            <Input
              id="new-password"
              type="password"
              name="confirm-password"
              placeholder="*******"
              required
              onChange={(e) => setConfirmPassword(e.target.value)}
            />
            {
              differentPasswords ? <ErrorText>Password and Confirm password are differents</ErrorText> : (resetPasswordError && <ErrorText>Could not reset password. Unknown error</ErrorText>)
            }

            <div className="actions">
              <button type="button" onClick={handleResetPassword} className="primary">Set new password</button>
              <button type="button" onClick={() => navigate('/login')} className="secondary">Cancel</button>
            </div>
          </ModalWrapper>
        </ModalContainer>
      }
      {
        passwordReseted && 
        <ModalContainer>
          <h2>Password successfully reseted</h2>
          <ModalWrapper>
            <div className="actions">
              <button type="button" onClick={() => navigate('/login')} className="secondary">Go back to Login page</button>
            </div>
          </ModalWrapper>
        </ModalContainer>
      }
    </Overlay>
  );
}

export default ForgotPassword;
