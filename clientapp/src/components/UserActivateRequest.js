import React, { useContext, useEffect } from "react";
import { UserContext } from "../contexts/UserContext";
import { userService } from "../services/UserService";

const UserActivateRequest = (props) => {
	const { userState, dispatch } = useContext(UserContext);

	const userId = props.id;

	const handleSubmit = (e) => {
		e.preventDefault();
		let resendActivationLink = {
			email: userState.blockedUser.emailAddress,
		};

		userService.resendActivationLink(resendActivationLink, dispatch);
	};

	useEffect(() => {
		userService.checkIfUserIdExist(userId, dispatch);
	}, [userId, dispatch]);

	return (
		<React.Fragment>
			<div className="illustration">
				<i className="icon ion-ios-navigate"></i>
			</div>
			<form method="post" onSubmit={handleSubmit}>
				<div hidden={userState.blockedUser.showSuccessMessage}>
					Vas nalog je privremeno blokiran. Ukoliko zelite ponovno aktiviranje naloga pritisnite na dugme ispod nakon cega ce Vam na email: <b>{userState.blockedUser.emailAddress}</b> stici
					aktivacioni link.
				</div>
				<div className="form-group text-center" style={{ color: "red", fontSize: "0.8em" }} hidden={!userState.blockedUser.showError}>
					{userState.blockedUser.errorMessage}
				</div>
				<div hidden={userState.blockedUser.showSuccessMessage} className="form-group">
					<input className="btn btn-primary btn-block" type="submit" value="Send activation mail" />
				</div>
				<div hidden={!userState.blockedUser.showSuccessMessage} className="form-group text-center" style={{ fontSize: "1.3em" }}>
					Activation mail was sent successfully.
				</div>
			</form>
		</React.Fragment>
	);
};

export default UserActivateRequest;
