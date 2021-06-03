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
import CreatePostPage from "./pages/CreatePostPage";
import ProfilePage from "./pages/ProfilePage";
import EditProfilePage from "./pages/EditProfile";

function App() {
	return (
		<Router>
			<Switch>
				<ProtectedRoute roles={["user"]} exact path="/" redirectTo="/unauthorized" component={HomePage} />
				<ProtectedRoute roles={[]} redirectTo="/" path="/login" component={LoginPage} />
				<ProtectedRoute roles={[]} redirectTo="/" path="/forgot-password" component={ForgotPasswordPage} />
				<ProtectedRoute roles={[]} redirectTo="/" path="/registration" component={RegistrationPage} />
				<ProtectedRoute roles={[]} redirectTo="/" path="/reset-password/:id" component={ResetPasswordPage} />
				<ProtectedRoute roles={[]} redirectTo="/" path="/blocked-user/:id" component={UserActivateRequestPage} />
				<ProtectedRoute roles={["user"]} redirectTo="/" path="/profile" component={ProfilePage} />
				<ProtectedRoute roles={["user"]} redirectTo="/" path="/edit-profile" component={EditProfilePage} />
				<ProtectedRoute roles={["user"]} redirectTo="/" path="/add-posts" component={CreatePostPage} />

				<Route path="/unauthorized" component={UnauthorizedPage} />

				<Route path="/404" component={PageNotFound} />
				<Redirect to="/404" />
			</Switch>
		</Router>
	);
}

export default App;
