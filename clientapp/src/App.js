import "./App.css";
import { HashRouter as Router, Link, Switch, Route } from "react-router-dom";
import LoginPage from "./pages/LoginPage";
import RegistrationPage from "./pages/RegistrationPage";
import ForgotPasswordPage from "./pages/ForgotPasswordPage";
import ResetPasswordPage from "./pages/ResetPasswordPage";

function App() {
	return (
		<Router>
			<Switch>
				<Link exact to="/login" path="/login" component={LoginPage} />
				<Link exact to="/forgot-password" path="/forgot-password" component={ForgotPasswordPage}/>
				<Link exact to="/registration" path="/registration" component={RegistrationPage} />
				<Route path="/reset-password/:id" component={ResetPasswordPage} />
			</Switch>
		</Router>
	);
}

export default App;
