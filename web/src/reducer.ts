import { combineReducers } from "redux";
import { reducer as login } from "./features/login";

export const reducer = combineReducers({
  login
});

export type RootState = ReturnType<typeof reducer>;
