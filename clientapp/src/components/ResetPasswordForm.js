import { useContext, useState } from "react";
import { userConstants } from "../constants/UserConstants";
import { UserContext } from "../contexts/UserContext";


const ResetPasswordForm = (props) => {
	const { dispatch } = useContext(UserContext);

	const [password, setPassword] = useState("");
	const [passwordRepeat, setPasswordRepeat] = useState("");
    let resetPasswordId= props.id

	const handleSubmit = (e) => {
		e.preventDefault();

		let resetPasswordRequest = {
            resetPasswordId,
			password,
			passwordRepeat,
		};

		alert('test')

		dispatch({ type: userConstants.RESET_PASSWORD_REQUEST, resetPasswordRequest });
	};

	return (
		<form method="post" onSubmit={handleSubmit}>
			<h2 className="sr-only">Login Form</h2>
			<div className="illustration">
				<i className="icon ion-ios-navigate"></i>
			</div>
			<div className="form-group">
				<input className="form-control" type="password" name="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)}></input>
			</div>
			<div className="form-group">
				<input className="form-control" type="password" name="password" placeholder="Repeat your password" value={passwordRepeat} onChange={(e) => setPasswordRepeat(e.target.value)}></input>
			</div>
			<div className="form-group">
				<input className="btn btn-primary btn-block" type="submit" value="Reset password" />
			</div>
		</form>
	);
};

export default ResetPasswordForm;
