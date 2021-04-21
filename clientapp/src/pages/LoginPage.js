import React, { Component } from "react";
import LoginForm from "../components/LoginForm";
import UserContextProvider from "../contexts/UserContext";

class LoginPage extends Component {
	render() {
		return (
			<React.Fragment>
				<section className="login-clean">
					<UserContextProvider>
						<LoginForm />
					</UserContextProvider>
				</section>
			</React.Fragment>
		);
	}
}

export default LoginPage;
