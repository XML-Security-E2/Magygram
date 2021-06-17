import { notificationConstants } from "../constants/NotificationConstants";

export const notificationReducer = (state, action) => {
	switch (action.type) {
		case notificationConstants.SET_USER_NOTIFICATIONS_REQUEST:
			return {
				...state,
				notifications: [],
			};
		case notificationConstants.SET_USER_NOTIFICATIONS_SUCCESS:
			return {
				...state,
				notifications: action.notifications,
			};
		case notificationConstants.SET_USER_NOTIFICATIONS_FAILURE:
			return {
				...state,
				notifications: [],
			};

		case notificationConstants.VIEW_NOTIFICATIONS_REQUEST:
			return state;

		case notificationConstants.VIEW_NOTIFICATIONS_SUCCESS:
			return {
				...state,
				notificationsNumber: 0,
			};
		case notificationConstants.VIEW_NOTIFICATIONS_REQUEST:
			return state;

		case notificationConstants.NOTIFICATION_RECEIVED:
			return {
				...state,
				notificationsNumber: action.count,
			};
		default:
			return state;
	}
};
