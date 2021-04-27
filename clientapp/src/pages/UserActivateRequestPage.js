import React from "react";
import UserActivateRequest from "../components/UserActivateRequest";
import UserContextProvider from "../contexts/UserContext";

const UserActivateRequestPage = (props) => {
	return (
		<React.Fragment>
			<section className="login-clean">
				<UserContextProvider>
					<UserActivateRequest id={props.match.params.id} />
				</UserContextProvider>
			</section>
		</React.Fragment>
	);
};

export default UserActivateRequestPage;
