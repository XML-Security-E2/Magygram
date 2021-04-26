import { useContext, useState } from "react";
import { userConstants } from "../constants/UserConstants";
import { UserContext } from "../contexts/UserContext";

const LoginForm = () => {
	const { dispatch } = useContext(UserContext);

	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");

	const handleSubmit = (e) => {
		e.preventDefault();

		let loginRequest = {
			email,
			password,
		};

		dispatch({ type: userConstants.LOGIN_REQUEST, loginRequest });
	};

	return (
		<form method="post" onSubmit={handleSubmit}>
			<h2 className="sr-only">Login Form</h2>
			<div className="illustration">
				<i className="icon ion-ios-navigate"></i>
			</div>
			<div  className="form-group">
				<input  className="form-control" required type="email" name="email" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)}></input>
			</div>
			
			<div className="form-group">
				<input  className="form-control" required type="password" name="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)}></input>
			</div>
			<div className="form-group">
				<input className="btn btn-primary btn-block" type="submit" value="Log In" />
			</div>
			<a className="forgot" href="#/forgot-password">
				Forgot your email or password?
			</a>
		</form>
	);
};

export default LoginForm;
