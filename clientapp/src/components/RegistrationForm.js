import React, { useState, useContext } from "react";
import { UserContext } from "../contexts/UserContext";
import { userConstants } from "../constants/UserConstants";

const RegistrationForm = () => {
	const { dispatch } = useContext(UserContext);

	const [name, setName] = useState("");
	const [surname, setSurname] = useState("");
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");
	const [repeatedPassword, setRepeatedPassword] = useState("");
	const [passwordError, setRepeatedPasswordError] = useState("none");
	const [passError, setPasswordError] = useState("none");

	const regexPassword = /^(.{0,7}|[^0-9]*|[^A-Z]*|[^a-z]*|[^!@#$%^&*(),.?":{}|<>~'_+=]*)$/;

	const handleSubmit = (e) => {
		setPasswordError("none")
		setRepeatedPasswordError("none")
		e.preventDefault();
		if (regexPassword.test(password) === true) {
			setPasswordError("initial")
		}
		else {
			if (password !== repeatedPassword) {
				setRepeatedPasswordError("initial")
			} else {
				let user = {
					name,
					surname,
					email,
					password,
					repeatedPassword,
				};
			}
		}

	};


	return (
		<form method="post" onSubmit={handleSubmit}>
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
			<div className="text-danger" style={{ display: passwordError }}>
				Passwords must be the same.
									</div>
			<div className="text-danger" style={{ display: passError }}>
				Password must contain minimum eight characters, at least one capital letter, one number and one special character.
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
		</form>
	);
};

export default RegistrationForm;
