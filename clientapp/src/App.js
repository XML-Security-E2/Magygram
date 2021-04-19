import "./App.css";
import { HashRouter as Router, Link, Switch } from "react-router-dom";
import LoginPage from "./pages/LoginPage";
import RegistrationPage from "./pages/RegistrationPage";

function App() {
	return (
		<Router>
			<Switch>
				<Link exact to="/login" path="/login" component={LoginPage} />
				<Link exact to="/registration" path="/registration" component={RegistrationPage} />
			</Switch>
		</Router>
	);
}

export default App;
