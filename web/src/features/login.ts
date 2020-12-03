import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface Login {
  loggedIn: boolean;
}

const initialState: Login = { loggedIn: false };

const loginSlice = createSlice({
  name: "login",
  initialState,
  reducers: {
    login: (state, action: PayloadAction<boolean>) => {
      state.loggedIn = action.payload;
    }
  }
});

export const { actions } = loginSlice;
export const reducer = loginSlice.reducer;
