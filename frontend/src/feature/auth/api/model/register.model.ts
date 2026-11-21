export type RegisterRequest = {
  email: string;
  password: string;
  fullName: string;
};

export type RegisterResponse = {
  email: string;
  fullName: string;
};
