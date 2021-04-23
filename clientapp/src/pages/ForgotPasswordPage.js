import React, { Component } from "react";
import ForgotPasswordForm from "../components/ForgotPasswordForm";
import UserContextProvider from "../contexts/UserContext";

class ForgotPasswordPage extends Component {
	render() {
		return (
			<React.Fragment>
				<section className="login-clean">
					<UserContextProvider>
                        <ForgotPasswordForm/>
					</UserContextProvider>
				</section>
			</React.Fragment>
		);
	}
}

export default ForgotPasswordPage;
