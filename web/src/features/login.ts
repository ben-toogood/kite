import { createSlice, PayloadAction } from "@reduxjs/toolkit";

// Note, you should NOT store access tokens like this in production
// you should use cookies
interface Login {
  loggedIn: boolean;
  requested: boolean;
  accessToken: string | null;
  refreshToken: string | null;
}

const initialState: Login = {
  requested: false,
  loggedIn: Boolean(localStorage.getItem("accessToken")),
  accessToken: localStorage.getItem("accessToken"),
  refreshToken: localStorage.getItem("refreshToken")
};

const loginSlice = createSlice({
  name: "login",
  initialState,
  reducers: {
    request: (state, action: PayloadAction<boolean>) => {
      state.requested = action.payload;
    },
    login: (state, action: PayloadAction<string[]>) => {
      localStorage.setItem("accessToken", action.payload[0]);
      localStorage.setItem("refreshToken", action.payload[1]);
      state.accessToken = action.payload[0];
      state.refreshToken = action.payload[0];
      state.loggedIn = true;
      state.requested = false;
    }
  }
});

export const { actions } = loginSlice;
export const reducer = loginSlice.reducer;
