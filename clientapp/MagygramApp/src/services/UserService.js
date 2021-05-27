import Axios from "axios";
import { config } from "../config/config";
import { userConstants } from "../constants/UserConstants";
import { deleteLocalStorage, setAuthInLocalStorage } from "../helpers/auth-header";

export const userService = {
	login,
	logout,
	register,
	resetPasswordLinkRequest,
	resetPasswordRequest,
	resendActivationLink,
	checkIfUserIdExist,
};

function login(loginRequest, dispatch) {
	dispatch(request());

	Axios.post(`${config.API_URL}/users/login`, loginRequest, { validateStatus: () => true })
		.then((res) => {
			if (res.status === 200) {
				setAuthInLocalStorage(res.data);
				dispatch(success());
				window.location = "#/";
			} else if (res.status === 401) {
				dispatch(failure(res.data.message));
			} else if (res.status === 403) {
				window.location = "#/blocked-user/" + res.data.userId;
			} else {
				dispatch({ type: userConstants.LOGIN_FAILURE });
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: userConstants.LOGIN_REQUEST };
	}
	function success() {
		return { type: userConstants.LOGIN_SUCCESS };
	}
	function failure(error) {
		return { type: userConstants.LOGIN_FAILURE, error };
	}
}

function resendActivationLink(resendActivationLink, dispatch) {
	dispatch(request());

	Axios.post(`${config.API_URL}/users/resend-activation-link`, resendActivationLink, { validateStatus: () => true })
		.then((res) => {
			if (res.status === 200) {
				dispatch(success());
			} else {
				dispatch(failure("Activation mail was not sent. Please, try again."));
			}
		})
		.catch((err) => {});

	function request() {
		return { type: userConstants.RESEND_ACTIVATION_LINK_REQUEST };
	}
	function success() {
		return { type: userConstants.RESEND_ACTIVATION_LINK_SUCCESS };
	}
	function failure(error) {
		return { type: userConstants.RESEND_ACTIVATION_LINK_FAILURE, errorMessage: error };
	}
}

function resetPasswordLinkRequest(resetPasswordLinkRequest, dispatch) {
	dispatch(request());

	Axios.post(`${config.API_URL}/users/reset-password-link-request`, resetPasswordLinkRequest, { validateStatus: () => true })
		.then((res) => {
			if (res.status === 200) {
				dispatch(success(resetPasswordLinkRequest.email));
			} else if (res.status === 404) {
				dispatch(failure("Sorry, your email was not found. Please double-check your email."));
			} else {
				dispatch(failure(res.data.message));
			}
		})
		.catch((err) => {
			console.log(err);
		});

	function request() {
		return { type: userConstants.RESET_PASSWORD_LINK_REQUEST };
	}
	function success(emailAddress) {
		return { type: userConstants.RESET_PASSWORD_LINK_SUCCESS, emailAddress };
	}
	function failure(error) {
		return { type: userConstants.RESET_PASSWORD_LINK_FAILURE, errorMessage: error };
	}
}

function resetPasswordRequest(resetPasswordRequest, dispatch) {
	let [passwordValid, passwordErrorMessage] = validatePasswords(resetPasswordRequest.password, resetPasswordRequest.passwordRepeat);

	if (passwordValid === true) {
		dispatch(request());

		Axios.post(`${config.API_URL}/users/reset-password`, resetPasswordRequest, { validateStatus: () => true })
			.then((res) => {
				if (res.status === 200) {
					dispatch(success());
				} else {
					console.log(res);
					dispatch(failure(res.data.message));
				}
			})
			.catch((err) => {
				console.log(err);
			});
	} else {
		dispatch(failure(passwordErrorMessage));
	}

	function request() {
		return { type: userConstants.RESET_PASSWORD_REQUEST };
	}
	function success() {
		return { type: userConstants.RESET_PASSWORD_SUCCESS };
	}
	function failure(error) {
		return { type: userConstants.RESET_PASSWORD_FAILURE, errorMessage: error };
	}
}

function logout() {
	deleteLocalStorage();
	window.location = "#/login";
}

function register(user, dispatch) {
	if (validateUser(user, dispatch)) {
		dispatch(request());
		Axios.post(`${config.API_URL}/users`, user, { validateStatus: () => true })
			.then((res) => {
				if (res.status === 201) {
					dispatch(success(user.email));
				} else {
					dispatch(failure(res.data.message));
				}
			})
			.catch((err) => {
				console.log(err);
			});
	}

	function request() {
		return { type: userConstants.REGISTER_REQUEST };
	}
	function success(emailAddress) {
		return { type: userConstants.REGISTER_SUCCESS, emailAddress };
	}
	function failure(error) {
		return { type: userConstants.REGISTER_FAILURE, errorMessage: error };
	}
}

function validatePasswords(password, repeatedPassword) {
	const regexPassword = /^(.{0,7}|[^0-9]*|[^A-Z]*|[^a-z]*|[^!@#$%^&*(),.?":{}|<>~'_+=]*)$/;

	if (regexPassword.test(password) === true) {
		return [false, "Password must contain minimum eight characters, at least one capital letter, one number and one special character."];
	} else if (password !== repeatedPassword) {
		return [false, "Passwords must be the same."];
	} else {
		return [true, ""];
	}
}

function validateUser(user, dispatch) {
	const [passwordValid, passwordErrorMessage] = validatePasswords(user.password, user.repeatedPassword);

	if (user.name.length < 2) {
		dispatch(validatioFailure("Name must contain minimum two letters"));
		return false;
	} else if (user.surname.length < 2) {
		dispatch(validatioFailure("Surname must contain minimum two letters"));
		return false;
	} else if (passwordValid === false) {
		dispatch(validatioFailure(passwordErrorMessage));
		return false;
	}

	function validatioFailure(message) {
		return { type: userConstants.REGISTER_VALIDATION_FAILURE, errorMessage: message };
	}

	return true;
}

function checkIfUserIdExist(userId, dispatch) {
	Axios.get(`${config.API_URL}/users/check-existence/` + userId, { validateStatus: () => true })
		.then((res) => {
			if (res.status === 200) {
				dispatch(success(res.data.emailAddress));
			} else if (res.status === 404) {
				window.location = "#/404";
			}
		})
		.catch((err) => {});

	function success(emailAddress) {
		return { type: userConstants.BLOCKED_USER_EMAIL_REQUEST, emailAddress };
	}
}
