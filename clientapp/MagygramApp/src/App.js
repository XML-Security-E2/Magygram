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
import EditProfilePage from "./pages/EditProfilePage";
import UserProfilePage from "./pages/UserProfilePage";
import SearchPostPage from "./pages/SearchedPostPage";
import LikedPostPage from "./pages/LikedPostPage";
import DislikedPostPage from "./pages/DislikedPostPage";

function App() {
	return (
		<Router>
			<Switch>
				<ProtectedRoute roles={"*"} exact path="/" redirectTo="/unauthorized" component={HomePage} />
				<ProtectedRoute roles={[]} redirectTo="/" path="/login" component={LoginPage} />
				<ProtectedRoute roles={[]} redirectTo="/" path="/forgot-password" component={ForgotPasswordPage} />
				<ProtectedRoute roles={[]} redirectTo="/" path="/registration" component={RegistrationPage} />
				<ProtectedRoute roles={[]} redirectTo="/" path="/reset-password/:id" component={ResetPasswordPage} />
				<ProtectedRoute roles={[]} redirectTo="/" path="/blocked-user/:id" component={UserActivateRequestPage} />
				<ProtectedRoute roles={["user"]} redirectTo="/unauthorized" path="/edit-profile" component={EditProfilePage} />
				<ProtectedRoute roles={["user"]} redirectTo="/unauthorized" path="/add-posts" component={CreatePostPage} />
				<Route path="/profile" component={UserProfilePage} />
				<ProtectedRoute roles={["user"]} redirectTo="/unauthorized" path="/edit-profile" component={EditProfilePage} />
				<ProtectedRoute roles={["user"]} redirectTo="/unauthorized" path="/add-posts" component={CreatePostPage} />
				<ProtectedRoute roles={["user","admin"]} redirectTo="/unauthorized" path="/search/hashtag/:id" component={SearchPostPage} />
				<ProtectedRoute roles={["user","admin"]} redirectTo="/unauthorized" path="/search/location/:id" component={SearchPostPage} />
				<ProtectedRoute roles={["user"]} redirectTo="/unauthorized" path="/liked" component={LikedPostPage} />
				<ProtectedRoute roles={["user"]} redirectTo="/unauthorized" path="/disliked" component={DislikedPostPage} />

				<Route path="/unauthorized" component={UnauthorizedPage} />

				<Route path="/404" component={PageNotFound} />
				<Redirect to="/404" />
			</Switch>
		</Router>
	);
}

export default App;
