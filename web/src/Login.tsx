import React, { useState } from "react";
import { actions } from "./features/login";
import { useDispatch } from "react-redux";
import { gql, useMutation } from "@apollo/client";
import { requestLogin } from "./gql-types";
import logo from "./logo.png"

const REQUEST_LOGIN = gql`
  mutation requestLogin($email: String!) {
    login(email: $email)
  }
`;

const Login = () => {
  const dispatch = useDispatch();
  const [email, setEmail] = useState("");
  const [requestLogin, { data, loading }] = useMutation<requestLogin>(REQUEST_LOGIN);

  const handleSubmit = (evt: React.FormEvent) => {
    evt.preventDefault();
    requestLogin({ variables: { email } });
  };

  if (data?.login) {
    dispatch(actions.request(data.login))
  }

  const requested = Boolean(loading || data?.login);

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md w-full space-y-8">
        <div>
          <img
            className="mx-auto h-24 w-auto"
            src={logo}
            alt="Workflow"
          />
        </div>
        <form className="mt-8 space-y-6" onSubmit={handleSubmit}>
          <input type="hidden" name="remember" value="true" />
          <div className="rounded-md shadow-sm -space-y-px">
            <div>
              <label htmlFor="email-address" className="sr-only">
                Email address
              </label>
              <input
                id="email-address"
                name="email"
                type="email"
                autoComplete="email"
                required
                className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
                placeholder="Email address"
                value={email}
                onChange={e => setEmail(e.target.value)}
              />
            </div>
          </div>

          <div>
            {requested ?
            <p className="flex justify-center mt-3 text-base text-gray-500 sm:mt-5 sm:text-lg sm:max-w-xl sm:mx-auto md:mt-5 md:text-xl lg:mx-0">Please check your email</p> :
            <button
              disabled={Boolean(loading || data?.login)}
              type="submit"
              className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              <span className="absolute left-0 inset-y-0 flex items-center pl-3">
                <svg
                  className="h-5 w-5 text-indigo-500 group-hover:text-indigo-400"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                  aria-hidden="true"
                >
                  <path
                    fillRule="evenodd"
                    d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z"
                    clipRule="evenodd"
                  />
                </svg>
              </span>
              Sign in
            </button>
    }
          </div>
        </form>
      </div>
    </div>
  );
};

export default Login;
