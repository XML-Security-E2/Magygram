import Axios from "axios";
import { config } from "../config/config";

export const userService = {
	login,
	logout,
	register,
};

function login(username, password) {}

function logout() {}

function register(user) {
	Axios.post(`${config.API_URL}/api/registration`, user)
		.then((res) => {
			console.log(res);
		})
		.catch((err) => {
			console.log(err);
		});

	return "";
}
