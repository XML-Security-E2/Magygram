import { useContext, useState } from "react";
import { userConstants } from "../constants/UserConstants";
import { UserContext } from "../contexts/UserContext";

const ForgotPasswordForm = () => {
	const { dispatch } = useContext(UserContext);

	const [email, setEmail] = useState("");

	const handleSubmit = (e) => {
		e.preventDefault();

		let resetPasswordLinkRequest = {
			email,
		};

		dispatch({ type: userConstants.RESET_PASSWORD_LINK_REQUEST, resetPasswordLinkRequest });
	};

	return (
		<form method="post" onSubmit={handleSubmit}>
			<h2 className="sr-only">Login Form</h2>
			<div className="illustration">
				<i className="icon ion-ios-navigate"></i>
			</div>
			<div className="form-group">
				<input className="form-control" type="email" name="email" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)}></input>
			</div>
			<div className="form-group">
				<input className="btn btn-primary btn-block" type="submit" value="Send reset email" />
			</div>
		</form>
	);
};

export default ForgotPasswordForm;
