import { useContext, useState } from "react";
import { Link } from "react-router-dom";
import { UserContext } from "../contexts/UserContext";
import { userService } from "../services/UserService";

const ResetPasswordForm = (props) => {
	const { userState, dispatch } = useContext(UserContext);

	const [password, setPassword] = useState("");
	const [passwordRepeat, setPasswordRepeat] = useState("");
	let resetPasswordId = props.id;

	const handleSubmit = (e) => {
		e.preventDefault();

		let resetPasswordRequest = {
			resetPasswordId,
			password,
			passwordRepeat,
		};

		userService.resetPasswordRequest(resetPasswordRequest, dispatch);
	};

	return (
		<form method="post" onSubmit={handleSubmit}>
			<h2 className="sr-only">Login Form</h2>
			<div className="illustration">
				<i className="icon ion-ios-navigate"></i>
			</div>
			<div hidden={userState.resetPassword.showSuccessMessage} className="form-group">
				<input className="form-control" type="password" name="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)}></input>
			</div>
			<div hidden={userState.resetPassword.showSuccessMessage} className="form-group">
				<input className="form-control" type="password" name="password" placeholder="Repeat your password" value={passwordRepeat} onChange={(e) => setPasswordRepeat(e.target.value)}></input>
			</div>
			<div className="form-group text-center" style={{ color: "red", fontSize: "0.8em" }} hidden={!userState.resetPassword.showError}>
				{userState.resetPassword.errorMessage}
			</div>
			<div hidden={userState.resetPassword.showSuccessMessage} className="form-group">
				<input className="btn btn-primary btn-block" type="submit" value="Reset password" />
			</div>

			<div hidden={!userState.resetPassword.showSuccessMessage} className="form-group text-center" style={{ fontSize: "1.3em" }}>
				You successfully changed your password.
			</div>
			<div hidden={!userState.resetPassword.showSuccessMessage} className="form-group">
				<Link className="btn btn-primary btn-block" to="/login">
					Back to login
				</Link>
			</div>
		</form>
	);
};

export default ResetPasswordForm;
