import Axios from "axios";
import { config } from "../config/config";

export const userService = {
	login,
	logout,
	register,
	resetPasswordLinkRequest,
	resetPasswordRequest,
};

function login(loginRequest) {
	Axios.post(`${config.API_URL}/users/login`, loginRequest)
		.then((res) => {
			console.log(res);
			localStorage.setItem("accessToken", res.data.accessToken);
			localStorage.setItem("role", res.data.role);
		})
		.catch((err) => {
			console.log(err);
		});
	return "";
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
