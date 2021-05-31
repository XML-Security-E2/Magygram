import "./App.css";
import { HashRouter as Router, Switch, Route, Redirect } from "react-router-dom";
import LoginPage from "./pages/LoginPage";
import RegistrationPage from "./pages/RegistrationPage";
import ForgotPasswordPage from "./pages/ForgotPasswordPage";
import ResetPasswordPage from "./pages/ResetPasswordPage";
import UserActivateRequestPage from "./pages/UserActivateRequestPage";
import HomePage from "./pages/HomePage";
import PageNotFound from "./pages/PageNotFound";
import { ProtectedRoute } from "./router/ProtectedRouter";
import UnauthorizedPage from "./pages/UnauthorizedPage";
import PostContextProvider from "./contexts/PostContext";

function App() {
	return (
		<Router>
			<Switch>
			<PostContextProvider>
				<ProtectedRoute roles={[]} exact path="/" redirectTo="/unauthorized" component={HomePage} />
			</PostContextProvider>
				<ProtectedRoute roles={["user"]} redirectTo="/" path="/login" component={LoginPage} />
				<ProtectedRoute roles={[]} redirectTo="/" path="/forgot-password" component={ForgotPasswordPage} />
				<ProtectedRoute roles={[]} redirectTo="/" path="/registration" component={RegistrationPage} />
				<ProtectedRoute roles={[]} redirectTo="/" path="/reset-password/:id" component={ResetPasswordPage} />
				<ProtectedRoute roles={[]} redirectTo="/" path="/blocked-user/:id" component={UserActivateRequestPage} />
				<Route path="/unauthorized" component={UnauthorizedPage} />

				<Route path="/404" component={PageNotFound} />
				<Redirect to="/404" />
			</Switch>
		</Router>
	);
}

export default App;
