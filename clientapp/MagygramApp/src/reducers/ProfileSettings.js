import { profileSettingsConstants } from "../constants/ProfileSettingsConstants";

export const profileSettingsReducer = (state, action) => {
	switch (action.type) {
		case profileSettingsConstants.SHOW_EDIT_PROFILE_PAGE:
			return {
				...state,
				activeSideBar: {
					showEditProfile: true,
					showVerifyAccount: false,
					showEditNotifications: false,
				},
			};
		case profileSettingsConstants.SHOW_VERIFY_ACCOUNT_PAGE:
			return {
				...state,
				activeSideBar: {
					showEditProfile: false,
					showVerifyAccount: true,
					showEditNotifications: false,
				},
			};

		case profileSettingsConstants.SHOW_EDIT_NOTIFICATIONS_PAGE:
			return {
				...state,
				activeSideBar: {
					showEditProfile: false,
					showVerifyAccount: false,
					showEditNotifications: true,
				},
			};
		case profileSettingsConstants.CREATE_VERIFICATION_REQUEST_REQUEST:
			return {
				...state,
				sendRequest: {
					showError: false,
					errorMessage: "",
				},
			};
		case profileSettingsConstants.CREATE_VERIFICATION_REQUEST_SUCCESS:
			return {
				...state,
				sendedVerifyRequest: true,
				sendRequest: {
					showError: false,
					errorMessage: "",
				},
			};
		case profileSettingsConstants.CREATE_VERIFICATION_REQUEST_FAILURE:
			return {
				...state,
				sendRequest: {
					showError: true,
					errorMessage: action.error,
				},
			};
		default:
			return state;
	}
};
