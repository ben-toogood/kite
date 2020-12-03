import Feed from "./Feed";
import Login from "./Login";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import { useSelector } from "./store";

function App() {
  const loggedIn = useSelector(s => s.login.loggedIn);

  //   const { data, loading, error } = useQuery<getUsers>(GET_USERS);

  //   if (loading) return <p>Loading...</p>;
  //   if (error) return <p>ERROR: {error.message}</p>;
  //   console.log(data);
  //   if (!data) return null;

  return (
    <Router>
      <Switch>
        <Route path="/">{loggedIn ? <Feed /> : <Login />}</Route>
      </Switch>
    </Router>
  );
}

export default App;
