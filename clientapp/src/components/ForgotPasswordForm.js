import { useContext, useState } from "react";
import { Link } from "react-router-dom";
import { UserContext } from "../contexts/UserContext";
import { userService } from "../services/UserService";

const ForgotPasswordForm = () => {
	const {userState, dispatch } = useContext(UserContext);

	const [email, setEmail] = useState("");

	const handleSubmit = (e) => {
		e.preventDefault();

		let resetPasswordLinkRequest = {
			email,
		};

		userService.resetPasswordLinkRequest(resetPasswordLinkRequest,dispatch);
	};

	return (
		<form method="post" onSubmit={handleSubmit}>
			<h2 className="sr-only">Login Form</h2>
			<div className="illustration">
				<i className="icon ion-ios-navigate"></i>
			</div>
			<div hidden={userState.forgotPasswordLinkError.showSuccessMessage} className="form-group">
				<input className="form-control" type="email" required name="email" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)}></input>
			</div>
			<div className="form-group text-center" style={{ color: 'red' , fontSize:'0.8em'}} hidden={!userState.forgotPasswordLinkError.showError}>
				{userState.forgotPasswordLinkError.errorMessage}
			</div>
			<div hidden={userState.forgotPasswordLinkError.showSuccessMessage} className="form-group">
				<input className="btn btn-primary btn-block" type="submit" value="Send reset email" />
			</div>

			<div hidden={!userState.forgotPasswordLinkError.showSuccessMessage} className="form-group text-center" style={{ fontSize:'1.3em'}}>
				We sent an email to <b>{userState.forgotPasswordLinkError.emailAddress}</b> with a link to get back into your account.
			</div>
			<div hidden={!userState.forgotPasswordLinkError.showSuccessMessage} className="form-group">
				<Link className="btn btn-primary btn-block" to="/login">Back to login</Link>
			</div>
		</form>
	);
};

export default ForgotPasswordForm;
