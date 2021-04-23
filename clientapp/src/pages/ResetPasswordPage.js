import React from "react";
import UserContextProvider from "../contexts/UserContext";
import ResetPasswordForm from "../components/ResetPasswordForm"

const ResetPasswordPage = (props) => {

	return (
		<React.Fragment>
				<section className="login-clean">
					<UserContextProvider>
                        <ResetPasswordForm id={props.match.params.id}/>
					</UserContextProvider>
				</section>
			</React.Fragment>
	);
};

export default ResetPasswordPage;
