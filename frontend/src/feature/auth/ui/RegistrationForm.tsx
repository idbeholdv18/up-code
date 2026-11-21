import { zodResolver } from "@hookform/resolvers/zod";
import { FC, useId } from "react";
import { useForm } from "react-hook-form";
import z from "zod";
import { useRegisterMutation } from "../api/api/api";
import { registrationSchema } from "../api/model/schema";

type RegistrationFormProps = {};

type RegistrationFormValues = z.infer<typeof registrationSchema>;

export const RegistrationForm: FC<RegistrationFormProps> = (props) => {
  const [registerUser, result] = useRegisterMutation();

  const emailInputId = useId();
  const fullNameInputId = useId();
  const passwordInputId = useId();

  const {
    register,
    handleSubmit,
    formState: { errors, isValid },
  } = useForm<RegistrationFormValues>({
    resolver: zodResolver(registrationSchema),
    mode: "onBlur",
  });

  const onSubmit = (data: RegistrationFormValues) => {
    registerUser(data);
  };

  return (
    <form
      onSubmit={handleSubmit(onSubmit)}
      noValidate
      className='text-fg-accent'
    >
      <div>
        <label htmlFor={fullNameInputId}>Имя</label>
        <input
          id={fullNameInputId}
          type='text'
          {...register("fullName")}
          className='border border-fg-primary'
        />
        {errors.fullName && (
          <span style={{ color: "red" }}>{errors.fullName.message}</span>
        )}
      </div>

      <div>
        <label htmlFor={emailInputId}>Email</label>
        <input
          id={emailInputId}
          type='email'
          {...register("email")}
          className='border border-fg-primary'
        />
        {errors.email && (
          <span style={{ color: "red" }}>{errors.email.message}</span>
        )}
      </div>

      <div>
        <label htmlFor={passwordInputId}>Пароль</label>
        <input
          id={passwordInputId}
          type='password'
          {...register("password")}
          className='border border-fg-primary'
        />
        {errors.password && (
          <span style={{ color: "red" }}>{errors.password.message}</span>
        )}
      </div>

      <button type='submit' disabled={result.isLoading || !isValid}>
        {result.isLoading ? "Регистрация..." : "Зарегистрироваться"}
      </button>
    </form>
  );
};
