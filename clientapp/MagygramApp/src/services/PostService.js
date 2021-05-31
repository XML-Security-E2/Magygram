import Axios from "axios";
import { postConstants } from "../constants/PostConstants";

export const postService = {
	findPostsForTimeline,
};

async function findPostsForTimeline(dispatch) {

	await Axios.get(`/user-api/api/users/employee/waiter`, { validateStatus: () => true })
		.then((res) => {
			if (res.status === 200) {
			} else {
			}
		})
		.catch((err) => {
			console.log(err);
		});

	function request() {
		//return { type: userConstants.SET_WAITERS_REQUEST };
	}
	function success(data) {
		//return { type: userConstants.SET_WAITERS_SUCCESS, waiters: data };
	}
	function failure(message) {
		//return { type: userConstants.SET_WAITERS_ERROR, errorMessage: message };
	}
}
