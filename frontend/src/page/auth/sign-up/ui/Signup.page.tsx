import { RegistrationForm } from "@/feature/auth";
import { FC } from "react";

type SignupPageProps = {};

const SignupPage: FC<SignupPageProps> = (props) => {
  return (
    <div>
      <h1>Sign Up</h1>
      <RegistrationForm />
    </div>
  );
};

export default SignupPage;
