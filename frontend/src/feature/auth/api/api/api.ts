import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { RegisterRequestDto, RegisterResponseDto } from "./dto";
import { RegisterRequest, RegisterResponse } from "../model/register.model";
import { baseApi } from "@/shared/api/baseApi";
import { SESSION_TAG } from "@/shared/api/tags";

export const authApi = baseApi.injectEndpoints({
  overrideExisting: false,
  endpoints: (builder) => ({
    register: builder.mutation<RegisterResponse, RegisterRequest>({
      query: (body) => ({
        url: "auth/register",
        method: "POST",
        body: {
          email: body.email,
          full_name: body.fullName,
          password: body.password,
        } satisfies RegisterRequestDto,
      }),
      transformResponse: (response: RegisterResponseDto): RegisterResponse => ({
        email: response.email,
        fullName: response.full_name,
      }),
      invalidatesTags: [SESSION_TAG],
    }),
  }),
});

export const { useRegisterMutation } = authApi;
