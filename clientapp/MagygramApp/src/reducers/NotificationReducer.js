import { modalConstants } from "../constants/ModalConstants";
import { notificationConstants } from "../constants/NotificationConstants";

var stateCpy = {};

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

		case notificationConstants.SET_PROFILE_NOTIFICATIONS_REQUEST:
			return {
				...state,
				notificationSettingsModal: {
					showModal: false,
					showSuccessMessage: false,
					successMessage: "",
					settings: {
						notifyPost: false,
						notifyStory: false,
					},
				},
			};
		case notificationConstants.SET_PROFILE_NOTIFICATIONS_SUCCESS:
			return {
				...state,
				notificationSettingsModal: {
					showModal: true,
					showSuccessMessage: false,
					successMessage: "",
					settings: action.notifications,
				},
			};
		case notificationConstants.SET_PROFILE_NOTIFICATIONS_FAILURE:
			return {
				...state,
				notificationSettingsModal: {
					showModal: false,
					showSuccessMessage: false,
					successMessage: "",
					settings: {
						notifyPost: false,
						notifyStory: false,
					},
				},
			};

		case modalConstants.HIDE_NOTIFICATION_SETTINGS_MODAL:
			return {
				...state,
				notificationSettingsModal: {
					showModal: false,
					showSuccessMessage: false,
					successMessage: "",
					settings: {
						notifyPost: false,
						notifyStory: false,
					},
				},
			};

		case notificationConstants.HIDE_NOTIFICATION_SETTINGS_SUCCESS_MESSAGE:
			stateCpy = { ...state };
			stateCpy.notificationSettingsModal.showSuccessMessage = false;
			stateCpy.notificationSettingsModal.successMessage = "";

			return stateCpy;

		case notificationConstants.EDIT_PROFILE_NOTIFICATIONS_REQUEST:
			stateCpy = { ...state };
			stateCpy.notificationSettingsModal.showSuccessMessage = false;
			stateCpy.notificationSettingsModal.successMessage = "";

			return stateCpy;
		case notificationConstants.EDIT_PROFILE_NOTIFICATIONS_SUCCESS:
			stateCpy = { ...state };
			stateCpy.notificationSettingsModal.showSuccessMessage = true;
			stateCpy.notificationSettingsModal.successMessage = action.successMessage;

			return stateCpy;
		case notificationConstants.EDIT_PROFILE_NOTIFICATIONS_FAILURE:
			return state;

		case notificationConstants.VIEW_NOTIFICATIONS_REQUEST:
			return state;

		case notificationConstants.VIEW_NOTIFICATIONS_SUCCESS:
			return {
				...state,
				notificationsNumber: 0,
			};
		case notificationConstants.VIEW_NOTIFICATIONS_FAILURE:
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
