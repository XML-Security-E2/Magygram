import Axios from "axios";
import { config } from "../config/config";

export const userService = {
	login,
	logout,
	register,
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
