import z from "zod";

export const registrationSchema = z.object({
  fullName: z.string().min(2, "Enter at least 2 symbols"),
  email: z
    .string()
    .min(1, "Email required")
    .pipe(z.email("Invalid email format")),
  password: z.string().min(8, "Password must contain at least 8 symbols"),
});
