import { useEffect } from "react";
import { actions } from "./features/login";
import { useDispatch } from "react-redux";
import { useLocation, useHistory } from "react-router-dom";
import { refreshLogin } from "./gql-types";
import { gql, useMutation } from "@apollo/client";

const REFRESH_LOGIN = gql`
  mutation refreshLogin($refreshToken: String!) {
    refreshTokens(refreshToken: $refreshToken) {
      accessToken
      refreshToken
    }
  }
`;

const useQuery = () => {
  return new URLSearchParams(useLocation().search);
};

const Code = () => {
  const dispatch = useDispatch();
  const history = useHistory();
  const query = useQuery();
  const [refresh, { data }] = useMutation<refreshLogin>(REFRESH_LOGIN);

  const code = query.get("code");
  useEffect(() => {
    if (data?.refreshTokens?.accessToken) {
      dispatch(
        actions.login([
          data?.refreshTokens?.accessToken,
          data?.refreshTokens?.refreshToken,
        ])
      );
      history.push("/");
    } else {
      refresh({ variables: { refreshToken: code } });
    }
  }, [refresh, code, data, dispatch, history]);

  return null;
};

export default Code;
