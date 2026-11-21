import {
  BaseQueryFn,
  createApi,
  FetchArgs,
  fetchBaseQuery,
  FetchBaseQueryError,
} from "@reduxjs/toolkit/query/react";
import { Mutex } from "async-mutex";
import { API_TAGS } from "./tags";

const mutex = new Mutex();

const rawBaseQuery = fetchBaseQuery({
  // TODO move to .env
  baseUrl: "http://localhost:8080",
  credentials: "include",
});

export const baseQueryWithReauth: BaseQueryFn<
  string | FetchArgs,
  unknown,
  FetchBaseQueryError
> = async (args, api, extraOptions) => {
  await mutex.waitForUnlock();
  let result = await rawBaseQuery(args, api, extraOptions);

  if (result.error && result.error.status === 401) {
    if (!mutex.isLocked()) {
      const release = await mutex.acquire();
      try {
        const refreshResult = await rawBaseQuery(
          {
            url: "/auth/refresh",
            method: "POST",
          },
          api,
          extraOptions,
        );

        if (!refreshResult.error) {
          result = await rawBaseQuery(args, api, extraOptions);
        } else {
          api.dispatch(baseApi.util.resetApiState());
        }
      } finally {
        release();
      }
    } else {
      await mutex.waitForUnlock();
      result = await rawBaseQuery(args, api, extraOptions);
    }
  }
  return result;
};

export const baseApi = createApi({
  endpoints: () => ({}),
  baseQuery: baseQueryWithReauth,
  reducerPath: "api",
  tagTypes: API_TAGS,
});
