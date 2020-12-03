import { configureStore, Action } from "@reduxjs/toolkit";
import { createSelectorHook } from "react-redux";
import { ThunkAction } from "redux-thunk";
import { reducer } from "./reducer";

export const store = configureStore({
  reducer
});

if (process.env.NODE_ENV === "development" && module.hot) {
  module.hot.accept("./reducer", () => {
    const newRootReducer = require("./reducer").default;
    store.replaceReducer(newRootReducer);
  });
}

export const useSelector = createSelectorHook<AppState>();

export type AppState = ReturnType<typeof reducer>;
export type AppThunk = ThunkAction<void, AppState, unknown, Action<string>>;
export type Dispatch = typeof store.dispatch;
