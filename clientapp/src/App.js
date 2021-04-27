import "./App.css";
import { HashRouter as Router, Switch, Route, Redirect } from "react-router-dom";
import LoginPage from "./pages/LoginPage";
import RegistrationPage from "./pages/RegistrationPage";
import ForgotPasswordPage from "./pages/ForgotPasswordPage";
import ResetPasswordPage from "./pages/ResetPasswordPage";
import UserActivateRequestPage from "./pages/UserActivateRequestPage";
import HomePage from "./pages/HomePage";
import PageNotFound from "./pages/PageNotFound";

function App() {
	return (
		<Router>
			<Switch>
				<Route exact path="/" component={HomePage} />
				<Route path="/login" component={LoginPage} />
				<Route path="/forgot-password" component={ForgotPasswordPage} />
				<Route path="/registration" component={RegistrationPage} />
				<Route path="/reset-password/:id" component={ResetPasswordPage} />
				<Route path="/blocked-user/:id" component={UserActivateRequestPage} />

				<Route path="/404" component={PageNotFound} />
				<Redirect to="/404" />
			</Switch>
		</Router>
	);
}

export default App;
