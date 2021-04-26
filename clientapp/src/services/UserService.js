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
	checkIfUserIdExist,
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
			}else if(res.status===403){
				window.location = "#/blocked-user/" + res.data.userId;
			}else{
				dispatch({ type: userConstants.LOGIN_FAILURE });
			}
		}).catch (err => console.error(err))

	function request() { return { type: userConstants.LOGIN_REQUEST } }
    function success() { return { type: userConstants.LOGIN_SUCCESS } }
    function failure(error) { return { type: userConstants.LOGIN_FAILURE, error } }
}

function resendActivationLink(resendActivationLink,dispatch) {

	Axios.post(`${config.API_URL}/users/resend-activation-link`, resendActivationLink)
		.then((res) => {

		})
		.catch((err) => {

		});
	return userConstants.USERS_LOGIN_FAILURE
}

function resetPasswordLinkRequest(resetPasswordLinkRequest,dispatch) {
	
	dispatch(request())

	Axios.post(`${config.API_URL}/users/reset-password-link-request`, resetPasswordLinkRequest, {validateStatus : () => true})
		.then((res) => {
			if(res.status===200){
				dispatch(success(resetPasswordLinkRequest.email))
			}else if(res.status===404){
				dispatch(failure("Sorry, your email was not found. Please double-check your email."))
			}else{
				dispatch(failure(res.data))
			}
		})
		.catch((err) => {
			console.log(err);
		});
	
	function request() { return { type: userConstants.RESET_PASSWORD_LINK_REQUEST } }
	function success(emailAddress) { return { type: userConstants.RESET_PASSWORD_LINK_SUCCESS, emailAddress } }
	function failure(error) { return { type: userConstants.RESET_PASSWORD_LINK_FAILURE, errorMessage: error } }
}

function resetPasswordRequest(resetPasswordRequest) {
	Axios.post(`${config.API_URL}/users/reset-password`, resetPasswordRequest, {validateStatus : () => true})
		.then((res) => {
			if(res.status===200){
				
			}
		})
		.catch((err) => {
			console.log(err);
		});
	return "";
}

function logout() {}

function register(user,dispatch) {

	if(validateUser(user,dispatch)){
		dispatch(request())
		Axios.post(`${config.API_URL}/users`, user, {validateStatus : () => true})
		.then((res) => {
			if(res.status===201){
				dispatch(success())
			}else{
				dispatch(failure(res.data))
			}
		})
		.catch((err) => {
			console.log(err);
		});
	}

	function request() { return { type: userConstants.REGISTER_REQUEST } }
	function success() { return { type: userConstants.REGISTER_SUCCESS } }
    function failure(error) { return { type: userConstants.REGISTER_FAILURE, errorMessage:error } }
}

function validateUser(user,dispatch){
	const regexPassword = /^(.{0,7}|[^0-9]*|[^A-Z]*|[^a-z]*|[^!@#$%^&*(),.?":{}|<>~'_+=]*)$/;

	if (user.name.length < 2) {
		dispatch(validatioFailure("Name must contain minimum two letters"));
		return false;
	}else if (user.surname.length < 2) {
		dispatch(validatioFailure("Surname must contain minimum two letters"));
		return false;
	}else if (regexPassword.test(user.password) === true) {
		dispatch(validatioFailure("Password must contain minimum eight characters, at least one capital letter, one number and one special character."));
		return false;
	}else if (user.password !== user.repeatedPassword) {
		dispatch(validatioFailure("Passwords must be the same."));
		return false;
	}

	function validatioFailure(message) { return { type: userConstants.REGISTER_VALIDATION_FAILURE, errorMessage: message } }

	return true;
}

function checkIfUserIdExist(userId,dispatch) {

	Axios.get(`${config.API_URL}/users/check-existance/`+ userId, {validateStatus : () => true})
		.then((res) => {
			if(res.status===200){
				dispatch(success(res.data.emailAddress))
			}else if(res.status===404){
				//TODO: redirect to page not found
			}
			
		})
		.catch((err) => {
		});

		function success(emailAddress) { return { type: userConstants.BLOCKED_USER_EMAIL_REQUEST, emailAddress } }

}