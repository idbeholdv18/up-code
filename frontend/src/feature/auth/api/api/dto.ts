export type RegisterRequestDto = {
  email: string;
  password: string;
  full_name: string;
};

export type RegisterResponseDto = {
  email: string;
  full_name: string;
};
