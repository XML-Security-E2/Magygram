import Axios from "axios";
import { notificationConstants } from "../constants/NotificationConstants";
import { authHeader } from "../helpers/auth-header";

export const notificationService = {
	getUserNotifiactions,
	viewNotifications,
};

async function getUserNotifiactions(dispatch) {
	dispatch(request());

	await Axios.get(`/api/notifications`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				dispatch(failure("Sorry, we have some internal problem."));
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: notificationConstants.SET_USER_NOTIFICATIONS_REQUEST };
	}
	function success(notifications) {
		return { type: notificationConstants.SET_USER_NOTIFICATIONS_SUCCESS, notifications };
	}
	function failure(error) {
		return { type: notificationConstants.SET_USER_NOTIFICATIONS_FAILURE, errorMessage: error };
	}
}

function viewNotifications(dispatch) {
	dispatch(request());

	Axios.put(`/api/notifications/view`, null, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success());
			} else {
				dispatch(failure("Sorry, we have some internal problem."));
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: notificationConstants.VIEW_NOTIFICATIONS_REQUEST };
	}
	function success() {
		return { type: notificationConstants.VIEW_NOTIFICATIONS_SUCCESS };
	}
	function failure(error) {
		return { type: notificationConstants.VIEW_NOTIFICATIONS_FAILURE, errorMessage: error };
	}
}
