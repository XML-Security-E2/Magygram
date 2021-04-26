import Axios from "axios";
import { config } from "../config/config";
import { userConstants } from "../constants/UserConstants";

export const userService = {
	login,
	logout,
	register,
	resetPasswordLinkRequest,
	resetPasswordRequest,
	resendActivationLink,
};

function login(loginRequest,dispatch) {

	dispatch(request());

    Axios.post(`${config.API_URL}/users/login`, loginRequest, {validateStatus : () => true})
		.then((res) => {
			if(res.status === 200){
				dispatch(success());
				window.location = "#/";
			}else if(res.status ===401){
				dispatch(failure(res.data))
			}
			dispatch({ type: userConstants.REGISTER_FAILURE });
		}).catch (err => console.error(err))

	function request() { return { type: userConstants.LOGIN_REQUEST } }
    function success() { return { type: userConstants.LOGIN_SUCCESS } }
    function failure(error) { return { type: userConstants.LOGIN_FAILURE, error } }
}

function resendActivationLink(resendActivationLink) {

	Axios.post(`${config.API_URL}/users/resend-activation-link`, resendActivationLink)
		.then((res) => {
			alert('test')
		})
		.catch((err) => {
		});
	return userConstants.USERS_LOGIN_FAILURE
}

function resetPasswordLinkRequest(resetPasswordLinkRequest) {
	Axios.post(`${config.API_URL}/users/reset-password-link-request`, resetPasswordLinkRequest)
		.then((res) => {
			console.log(res);
		})
		.catch((err) => {
			console.log(err);
		});
	return "";
}

function resetPasswordRequest(resetPasswordRequest) {
	Axios.post(`${config.API_URL}/users/reset-password`, resetPasswordRequest)
		.then((res) => {
			console.log(res);
		})
		.catch((err) => {
			console.log(err);
		});
	return "";
}

function logout() {}

function register(user) {
	Axios.post(`${config.API_URL}/users`, user)
		.then((res) => {
			console.log(res);
		})
		.catch((err) => {
			console.log(err);
		});

	return "";
}
