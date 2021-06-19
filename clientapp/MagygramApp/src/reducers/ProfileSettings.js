import { profileSettingsConstants } from "../constants/ProfileSettingsConstants";
import { userConstants } from "../constants/UserConstants";

export const profileSettingsReducer = (state, action) => {
	switch (action.type) {
		case profileSettingsConstants.SHOW_EDIT_PROFILE_PAGE:
			return {
				...state,
				activeSideBar: {
					showEditProfile: true,
					showVerifyAccount: false,
					showEditNotifications: false,
					showEditPrivacySettings: false,
				},
			};
		case profileSettingsConstants.SHOW_VERIFY_ACCOUNT_PAGE:
			return {
				...state,
				activeSideBar: {
					showEditProfile: false,
					showVerifyAccount: true,
					showEditNotifications: false,
					showEditPrivacySettings: false,
				},
			};

		case profileSettingsConstants.SHOW_EDIT_NOTIFICATIONS_PAGE:
			return {
				...state,
				activeSideBar: {
					showEditProfile: false,
					showVerifyAccount: false,
					showEditNotifications: true,
					showEditPrivacySettings: false,
				},
			};
		case profileSettingsConstants.SHOW_EDIT_PRIVACY_SETTINGS_PAGE:
			return {
				...state,
				activeSideBar: {
					showEditProfile: false,
					showVerifyAccount: false,
					showEditNotifications: false,
					showEditPrivacySettings: true,
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
		case userConstants.CHECK_IF_USER_VERIFIED_REQUEST:
			return {
				...state,
				isUserVerified:false
			};
		case userConstants.CHECK_IF_USER_VERIFIED_SUCCESS:
			return {
				...state,
				isUserVerified:action.result
			};
		case userConstants.CHECK_IF_USER_VERIFIED_FAILURE:
			return {
				...state,
				isUserVerified:false
			};
		case profileSettingsConstants.CHECK_IF_USER_HAS_PENDING_REQUEST_REQUEST:
			return {
				...state,
				sendedVerifyRequest:false
			};
		case profileSettingsConstants.CHECK_IF_USER_HAS_PENDING_REQUEST_SUCCESS:
			return {
				...state,
				sendedVerifyRequest:action.result
			};
		case profileSettingsConstants.CHECK_IF_USER_HAS_PENDING_REQUEST_FAILURE:
			return {
				...state,
				sendedVerifyRequest:false
			};
		default:
			return state;
	}
};
