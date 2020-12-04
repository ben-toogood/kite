import Feed from "./Feed";
import Code from "./Code";
import Login from "./Login";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import { useSelector } from "./store";

function App() {
  const loggedIn = useSelector(s => s.login.loggedIn);

  return (
    <Router>
      <Switch>
        <Route path="/" exact>
          {loggedIn ? <Feed /> : <Login />}
        </Route>
        <Route path="/login?">
          <Code />
        </Route>
      </Switch>
    </Router>
  );
}

export default App;
