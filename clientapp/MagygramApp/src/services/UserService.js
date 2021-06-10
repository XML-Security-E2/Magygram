import Axios from "axios";
import { userConstants } from "../constants/UserConstants";
import { deleteLocalStorage, setAuthInLocalStorage } from "../helpers/auth-header";
import { authHeader } from "../helpers/auth-header";
import { postService } from "./PostService";

export const userService = {
	loginFirstAuthorization,
	logout,
	register,
	getUserProfileByUserId,
	editUser,
	editUserImage,
	resetPasswordLinkRequest,
	resetPasswordRequest,
	findAllFollowingUsers,
	findAllFollowedUsers,
	findAllFollowRequests,
	acceptFollowRequest,
	resendActivationLink,
	checkIfUserIdExist,
	followUser,
	unfollowUser,
	loginSecondAuthorization
};

async function findAllFollowedUsers(userId, dispatch) {
	dispatch(request());
	await Axios.get(`/api/users/${userId}/followed`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				dispatch(failure());
			}
		})
		.catch((err) => {
			dispatch(failure());
		});

	function request() {
		return { type: userConstants.SET_USER_FOLLOWING_REQUEST };
	}

	function success(data) {
		return { type: userConstants.SET_USER_FOLLOWING_SUCCESS, userInfos: data, header: "Followers" };
	}
	function failure() {
		return { type: userConstants.SET_USER_FOLLOWING_FAILURE };
	}
}

async function findAllFollowRequests(dispatch) {
	dispatch(request());
	await Axios.get(`/api/users/follow-requests`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				dispatch(failure());
			}
		})
		.catch((err) => {
			dispatch(failure());
		});

	function request() {
		return { type: userConstants.FOLLOW_REQUESTS_REQUEST };
	}

	function success(data) {
		return { type: userConstants.FOLLOW_REQUESTS_SUCCESS, userInfos: data };
	}
	function failure() {
		return { type: userConstants.FOLLOW_REQUESTS_FAILURE };
	}
}

async function findAllFollowingUsers(userId, dispatch) {
	dispatch(request());
	await Axios.get(`/api/users/${userId}/following`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				dispatch(failure());
			}
		})
		.catch((err) => {
			dispatch(failure());
		});

	function request() {
		return { type: userConstants.SET_USER_FOLLOWING_REQUEST };
	}

	function success(data) {
		return { type: userConstants.SET_USER_FOLLOWING_SUCCESS, userInfos: data, header: "Following" };
	}
	function failure() {
		return { type: userConstants.SET_USER_FOLLOWING_FAILURE };
	}
}

function loginFirstAuthorization(loginRequest, dispatch) {
	dispatch(request());

	Axios.post(`/api/auth/login`, loginRequest, { validateStatus: () => true })
		.then((res) => {
			if (res.status === 200) {
				dispatch(success())
			} else if (res.status === 401) {
				dispatch(failure("Sorry, your email or password was incorrect. Please double-check your password."));
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

function loginSecondAuthorization(loginRequest, dispatch) {
	dispatch(request());

	Axios.post(`/api/auth/login/two`, loginRequest, { validateStatus: () => true })
		.then((res) => {
			if (res.status === 200) {
				setAuthInLocalStorage(res.data);
				getLoggedData(dispatch);
			} else {
				dispatch(failure("Uneti kod je neispravan"));
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: userConstants.LOGIN_TWO_REQUEST };
	}
	function success() {
		return { type: userConstants.LOGIN_TWO_SUCCESS };
	}
	function failure(error) {
		return { type: userConstants.LOGIN_TWO_FAILURE, error };
	}
}

function acceptFollowRequest(userId, dispatch) {
	dispatch(request());

	Axios.post(`/api/users/follow-requests/${userId}/accept`, null, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			if (res.status === 200) {
				dispatch(success(userId));
			} else {
				dispatch(failure(res.data.message));
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: userConstants.ACCEPT_FOLLOW_REQUESTS_REQUEST };
	}

	function success(userId) {
		return { type: userConstants.ACCEPT_FOLLOW_REQUESTS_SUCCESS, userId };
	}

	function failure(error) {
		return { type: userConstants.ACCEPT_FOLLOW_REQUESTS_FAILURE, errorMessage: error };
	}
}

function editUser(userId, userRequestDTO, dispatch) {
	dispatch(request());

	Axios.put(`/api/users/${userId}`, userRequestDTO, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			if (res.status === 200) {
				dispatch(success("User info successfully changed"));
			} else {
				dispatch(failure(res.data.message));
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: userConstants.UPDATE_USER_REQUEST };
	}

	function success(successMessage) {
		return { type: userConstants.UPDATE_USER_SUCCESS, successMessage };
	}

	function failure(error) {
		return { type: userConstants.UPDATE_USER_FAILURE, errorMessage: error };
	}
}

function editUserImage(userId, image, dispatch) {
	let formData = new FormData();
	formData.append(`images[0]`, image);

	dispatch(request());

	Axios.put(`/api/users/${userId}/image`, formData, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				localStorage.setItem("imageURL", res.data);
				dispatch(success("User image successfully changed"));
			} else {
				dispatch(failure(res.data.message));
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: userConstants.UPDATE_USER_REQUEST };
	}

	function success(successMessage) {
		return { type: userConstants.UPDATE_USER_SUCCESS, successMessage };
	}

	function failure(error) {
		return { type: userConstants.UPDATE_USER_FAILURE, errorMessage: error };
	}
}

function getLoggedData(dispatch) {
	dispatch(request());

	Axios.get(`/api/users/logged`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			if (res.status === 200) {
				localStorage.setItem("userId", res.data.id);
				localStorage.setItem("username", res.data.username);
				localStorage.setItem("imageURL", res.data.imageUrl);

				dispatch(success());
				window.location = "#/";
			} else {
				dispatch(failure("Usled internog problema trenutno nije moguce logovanje"));
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

	Axios.post(`/api/users/resend-activation-link`, resendActivationLink, { validateStatus: () => true })
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

	Axios.post(`/api/users/reset-password-link-request`, resetPasswordLinkRequest, { validateStatus: () => true })
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

		Axios.post(`/api/users/reset-password`, resetPasswordRequest, { validateStatus: () => true })
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

function followUser(userId, dispatch) {
	let formData = new FormData();
	formData.append("userId", userId);
	dispatch(request());

	Axios.post(`/api/users/follow`, formData, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			if (res.status === 200) {
				dispatch(success(userId));
			} else if (res.status === 201) {
				dispatch(followRequestSuccess());
			} else {
				dispatch(failure(res.data.message));
			}
		})
		.catch((err) => {
			console.log(err);
		});

	function request() {
		return { type: userConstants.FOLLOW_USER_REQUEST };
	}
	function success(userId) {
		return { type: userConstants.FOLLOW_USER_SUCCESS, userId };
	}
	function followRequestSuccess() {
		return { type: userConstants.FOLLOW_USER_SEND_REQUEST_SUCCESS };
	}
	function failure() {
		return { type: userConstants.FOLLOW_USER_FAILURE };
	}
}

function unfollowUser(userId, dispatch) {
	let formData = new FormData();
	formData.append("userId", userId);
	dispatch(request());

	Axios.post(`/api/users/unfollow`, formData, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			if (res.status === 200) {
				dispatch(success(userId));
			} else {
				dispatch(failure(res.data.message));
			}
		})
		.catch((err) => {
			console.log(err);
		});

	function request() {
		return { type: userConstants.UNFOLLOW_USER_REQUEST };
	}
	function success(userId) {
		return { type: userConstants.UNFOLLOW_USER_SUCCESS, userId };
	}
	function failure() {
		return { type: userConstants.UNFOLLOW_USER_FAILURE };
	}
}

async function getUserProfileByUserId(userId, dispatch) {
	dispatch(request());

	Axios.get(`/api/users/${userId}/profile`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data, userId));
			} else {
				dispatch(failure("Usled internog problema trenutno nije moguce logovanje"));
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: userConstants.GET_USER_PROFILE_REQUEST };
	}
	function success(user, userId) {
		return { type: userConstants.GET_USER_PROFILE_SUCCESS, user, userId };
	}
	function failure(error) {
		return { type: userConstants.GET_USER_PROFILE_FAILURE, error };
	}
}

function logout() {
	deleteLocalStorage();
	window.location = "#/login";
}

function register(user, dispatch) {
	if (validateUser(user, dispatch)) {
		dispatch(request());
		Axios.post(`/api/users`, user, { responseType: 'arraybuffer' ,validateStatus: () => true })
			.then((res) => {
				if (res.status === 201) {
					let blob = new Blob(
						[res.data], 
						{ type: res.headers['image/png'] }
					)
					let image = URL.createObjectURL(blob)
					dispatch(success(user.email,image));
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
	function success(emailAddress,imageData) {
		return { type: userConstants.REGISTER_SUCCESS, emailAddress ,imageData};
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
	Axios.get(`/api/users/check-existence/` + userId, { validateStatus: () => true })
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
