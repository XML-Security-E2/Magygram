import React, { useState, useContext } from "react";
import { Link } from "react-router-dom";
import { UserContext } from "../contexts/UserContext";
import { userService } from "../services/UserService";

const RegistrationForm = () => {
	const {userState, dispatch } = useContext(UserContext);

	const [name, setName] = useState("");
	const [surname, setSurname] = useState("");
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");
	const [repeatedPassword, setRepeatedPassword] = useState("");

	const handleSubmit = (e) => {
		e.preventDefault();

		let user = {
			name,
			surname,
			email,
			password,
			repeatedPassword,
		};

		userService.register(user,dispatch);
	};


	return (
		<form method="post" onSubmit={handleSubmit}>
			<div hidden={userState.registrationError.showSuccessMessage}>
				<h2 className="text-center">
					<strong>Create</strong> an account.
				</h2>
				<div className="form-group">
					<input className="form-control" required type="text" name="nameInput" placeholder="Name" value={name} onChange={(e) => setName(e.target.value)} />
				</div>

				<div className="form-group">
					<input className="form-control" required type="text" name="surnameInput" placeholder="Surname" value={surname} onChange={(e) => setSurname(e.target.value)} />
				</div>

				<div className="form-group">
					<input className="form-control" required type="email" name="email" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
				</div>

				<div className="form-group">
					<input className="form-control" required type="password" name="passwordInput" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
				</div>

				<div className="form-group">
					<input
						className="form-control"
						required
						type="password"
						name="password-repeat"
						placeholder="Password (repeat)"
						value={repeatedPassword}
						onChange={(e) => setRepeatedPassword(e.target.value)}
					/>
				</div>
				<div className="form-group text-center" style={{ color: 'red' , fontSize:'0.8em'}} hidden={!userState.registrationError.showError}>
					{userState.registrationError.errorMessage}
				</div>
				<div className="form-group">
					<div className="form-check">
						<label className="form-check-label">
							<input className="form-check-input" type="checkbox" />I agree to the license terms
						</label>
					</div>
				</div>
				<div className="form-group">
					<input className="btn btn-primary btn-block" type="submit" value="Sign In" />
				</div>
				<a className="already" href="#/login">
					You already have an account? Login here.
				</a>
			</div>

			<section hidden={!userState.registrationError.showSuccessMessage} className="login-clean">
				<div className="form-group text-center" style={{ fontSize:'1.3em'}}>
					We sent an email to <b>{userState.registrationError.emailAddress}</b> with a activation link to activate your account.
				</div>
				<div className="form-group">
					<Link className="btn btn-primary btn-block" to="/login">Back to login</Link>
				</div>
			</section>

		</form>
	);
};

export default RegistrationForm;
