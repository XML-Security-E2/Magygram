import { useContext, useState } from "react";
import { UserContext } from "../contexts/UserContext";
import { userService } from "../services/UserService";

const LoginForm = (props) => {
	const { userState, dispatch } = useContext(UserContext);
	
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");


	const handleSubmit = (e) => {
		e.preventDefault();

		let loginRequest = {
			email,
			password,
		};

		userService.login(loginRequest,dispatch)
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
			<div className="form-group text-center" style={{ color: 'red' , fontSize:'0.8em'}} hidden={!userState.loginError.showError}>
				{userState.loginError.errorMessage}
			</div>
			<a className="forgot" href="#/forgot-password">
				Forgot your password?
			</a>
			<div className="form-group">
				<input className="btn btn-primary btn-block" type="submit" value="Log In" />
			</div>
			<div className="forgot">
				Don't have an account?
				<a className="forgot" href="#/registration">
					Sign up
				</a>
			</div>
		</form>
	);
};

export default LoginForm;
