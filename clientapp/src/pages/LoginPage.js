import React from "react";
import LoginForm from "../components/LoginForm";
import UserContextProvider from "../contexts/UserContext";

const LoginPage = () => {
	return (
		<React.Fragment>
			<section className="login-clean">
				<UserContextProvider>
					<LoginForm />
				</UserContextProvider>
			</section>
		</React.Fragment>
	);
};

export default LoginPage;
