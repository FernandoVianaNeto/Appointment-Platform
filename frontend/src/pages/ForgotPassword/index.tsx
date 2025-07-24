import { useState } from "react";
import { Input, ModalContainer, Overlay } from "./styles";

function ForgotPassword() {
  const [emailSent, setEmailSent] = useState<boolean>(false);

  return (
    <Overlay>
      <ModalContainer>
        <h2>Insert the registered email</h2>
        <form method="post">
          <label>
            Email*:
            <Input className="email-input" type="text" name="email" />
          </label>

          <div className="actions">
            <button type="submit">Sent</button>
            <button type="button" >Cancel</button>
          </div>

        </form>
      </ModalContainer>
    </Overlay>
  );
};

export default ForgotPassword;