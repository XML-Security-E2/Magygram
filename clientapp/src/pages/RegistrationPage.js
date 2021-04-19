import React from "react";
import RegistrationForm from "../components/RegistrationForm";
import UserContextProvider from "../contexts/UserContext";

const RegistrationPage = () => {
	return (
		<React.Fragment>
			<section className="register-photo">
				<div className="form-container">
					<div className="image-holder" />
					<UserContextProvider>
						<RegistrationForm />
					</UserContextProvider>
				</div>
			</section>
		</React.Fragment>
	);
};

export default RegistrationPage;
