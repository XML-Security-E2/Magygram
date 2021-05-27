import React from "react";
import ForgotPasswordForm from "../components/ForgotPasswordForm";
import UserContextProvider from "../contexts/UserContext";

const ForgotPasswordPage = () => {
	return (
		<React.Fragment>
			<section className="login-clean">
				<UserContextProvider>
					<ForgotPasswordForm />
				</UserContextProvider>
			</section>
		</React.Fragment>
	);
};

export default ForgotPasswordPage;
