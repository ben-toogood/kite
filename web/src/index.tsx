import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { ApolloProvider, ApolloClient, InMemoryCache } from "@apollo/client";
import { store } from "./store";
import App from "./App";
import "./index.css";

const client = new ApolloClient({
  uri: "https://api.deploy.wtf/graphql",
  cache: new InMemoryCache()
});

ReactDOM.render(
  <React.StrictMode>
    <Provider store={store}>
      <ApolloProvider client={client}>
        <App />
      </ApolloProvider>
    </Provider>
    ,
  </React.StrictMode>,
  document.getElementById("root")
);
