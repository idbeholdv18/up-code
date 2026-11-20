import { FC, useId, useState } from "react";
import { useRegisterMutation } from "../api/api";
import z from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

type RegistrationFormProps = {};

const registrationSchema = z.object({
  full_name: z.string().min(2, "Enter at least 2 symbols"),
  email: z
    .string()
    .min(1, "Email required")
    .pipe(z.email("Invalid email format")),
  password: z.string().min(8, "Password must contain at least 8 symbols"),
});

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
    <form onSubmit={handleSubmit(onSubmit)} noValidate>
      <div>
        <label htmlFor={fullNameInputId}>Имя</label>
        <input id={fullNameInputId} type='text' {...register("full_name")} />
        {errors.full_name && (
          <span style={{ color: "red" }}>{errors.full_name.message}</span>
        )}
      </div>

      <div>
        <label htmlFor={emailInputId}>Email</label>
        <input id={emailInputId} type='email' {...register("email")} />
        {errors.email && (
          <span style={{ color: "red" }}>{errors.email.message}</span>
        )}
      </div>

      <div>
        <label htmlFor={passwordInputId}>Пароль</label>
        <input id={passwordInputId} type='password' {...register("password")} />
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
