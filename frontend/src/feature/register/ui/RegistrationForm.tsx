import { FC, useId, useState } from "react";
import { useRegisterMutation } from "../api/api";

type RegistrationFormProps = {};

export const RegistrationForm: FC<RegistrationFormProps> = (props) => {
  const [register, result] = useRegisterMutation();
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [fullName, setFullName] = useState<string>("");

  const emailInputId = useId();
  const fullNameInputId = useId();
  const passwordInputId = useId();

  const onEmailInput = (
    e: React.ChangeEvent<HTMLInputElement, HTMLInputElement>,
  ) => {
    setEmail(e.target.value);
  };

  const onPasswordInput = (
    e: React.ChangeEvent<HTMLInputElement, HTMLInputElement>,
  ) => {
    setPassword(e.target.value);
  };

  const onFullNameInput = (
    e: React.ChangeEvent<HTMLInputElement, HTMLInputElement>,
  ) => {
    setFullName(e.target.value);
  };

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
      }}
      style={{
        display: "flex",
        flexDirection: "column",
        gap: "1rem",
      }}
    >
      <label htmlFor={emailInputId}>email</label>
      <input
        type='text'
        onChange={onEmailInput}
        value={email}
        id={emailInputId}
      />
      <label htmlFor={fullNameInputId}>full name</label>
      <input
        type='text'
        onChange={onFullNameInput}
        value={fullName}
        id={fullNameInputId}
      />
      <label htmlFor={passwordInputId}>password</label>
      <input
        type='text'
        onChange={onPasswordInput}
        value={password}
        id={passwordInputId}
      />
      <button
        onClick={() =>
          register({
            email: email,
            full_name: fullName,
            password: password,
          })
        }
      >
        {result.isLoading ? "loading..." : "create test user"}
      </button>
    </form>
  );
};
