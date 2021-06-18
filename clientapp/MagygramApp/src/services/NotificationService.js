import Axios from "axios";
import { notificationConstants } from "../constants/NotificationConstants";
import { authHeader } from "../helpers/auth-header";

export const notificationService = {
	getUserNotifiactions,
	viewNotifications,
	getProfileNotificationsSettings,
	editProfileNotificationsSettings,
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

async function getProfileNotificationsSettings(userId, dispatch) {
	dispatch(request());

	await Axios.get(`/api/users/notifications/get/${userId}`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200 || res.status === 404) {
				dispatch(success(res.data));
			} else {
				dispatch(failure("Sorry, we have some internal problem."));
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: notificationConstants.SET_PROFILE_NOTIFICATIONS_REQUEST };
	}
	function success(notifications) {
		return { type: notificationConstants.SET_PROFILE_NOTIFICATIONS_SUCCESS, notifications };
	}
	function failure(error) {
		return { type: notificationConstants.SET_PROFILE_NOTIFICATIONS_FAILURE, errorMessage: error };
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

function editProfileNotificationsSettings(settingsDTO, userId, dispatch) {
	dispatch(request());

	Axios.post(`/api/users/notifications/settings/${userId}`, settingsDTO, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success("Notifications setttings updated successfully"));
			} else {
				dispatch(failure("Sorry, we have some internal problem."));
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: notificationConstants.EDIT_PROFILE_NOTIFICATIONS_REQUEST };
	}
	function success(message) {
		return { type: notificationConstants.EDIT_PROFILE_NOTIFICATIONS_SUCCESS, successMessage: message };
	}
	function failure(error) {
		return { type: notificationConstants.EDIT_PROFILE_NOTIFICATIONS_FAILURE, errorMessage: error };
	}
}
