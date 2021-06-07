import { modalConstants } from "../constants/ModalConstants";
import { userConstants } from "../constants/UserConstants";

var a = {};

export const userReducer = (state, action) => {
	switch (action.type) {
		case userConstants.REGISTER_REQUEST:
			return {
				registrationError: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					emailAddress: "",
				},
			};
		case userConstants.REGISTER_SUCCESS:
			return {
				registrationError: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: true,
					emailAddress: action.emailAddress,
				},
			};
		case userConstants.REGISTER_FAILURE:
			return {
				registrationError: {
					showError: true,
					errorMessage: action.errorMessage,
					showSuccessMessage: false,
					emailAddress: "",
				},
			};
		case userConstants.REGISTER_VALIDATION_FAILURE:
			return {
				registrationError: {
					showError: true,
					errorMessage: action.errorMessage,
					showSuccessMessage: false,
					emailAddress: "",
				},
			};
		case userConstants.LOGIN_REQUEST:
			return {
				loginError: {
					showError: false,
					errorMessage: "",
				},
			};
		case userConstants.LOGIN_FAILURE:
			return {
				loginError: {
					showError: true,
					errorMessage: action.error,
				},
			};
		case userConstants.LOGIN_SUCCESS:
			return {
				...state,
				loginError: {
					showError: false,
					errorMessage: "",
				},
			};
		case userConstants.RESET_PASSWORD_LINK_REQUEST:
			return {
				forgotPasswordLinkError: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					emailAddress: "",
				},
			};
		case userConstants.RESET_PASSWORD_LINK_SUCCESS:
			return {
				forgotPasswordLinkError: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: true,
					emailAddress: action.emailAddress,
				},
			};
		case userConstants.RESET_PASSWORD_LINK_FAILURE:
			return {
				forgotPasswordLinkError: {
					showError: true,
					errorMessage: action.errorMessage,
					showSuccessMessage: false,
					emailAddress: "",
				},
			};

		case userConstants.RESET_PASSWORD_REQUEST:
			return {
				resetPassword: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
				},
			};
		case userConstants.RESET_PASSWORD_SUCCESS:
			return {
				resetPassword: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: true,
				},
			};
		case userConstants.RESET_PASSWORD_FAILURE:
			return {
				resetPassword: {
					showError: true,
					errorMessage: action.errorMessage,
					showSuccessMessage: false,
				},
			};
		case userConstants.RESEND_ACTIVATION_LINK_REQUEST:
			return {
				blockedUser: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
				},
			};
		case userConstants.RESEND_ACTIVATION_LINK_SUCCESS:
			return {
				blockedUser: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: true,
				},
			};
		case userConstants.RESEND_ACTIVATION_LINK_FAILURE:
			return {
				blockedUser: {
					showError: true,
					errorMessage: action.errorMessage,
					showSuccessMessage: false,
				},
			};
		case userConstants.BLOCKED_USER_EMAIL_REQUEST:
			return {
				blockedUser: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					emailAddress: action.emailAddress,
				},
			};

		case userConstants.GET_USER_PROFILE_REQUEST:
			return {
				...state,
				userProfile: {
					showedUserId: "",
					user: {
						username: "",
						name: "",
						surname: "",
						website: "",
						bio: "",
						following: "",
						gender: "",
						imageUrl: "",
						postNumber: "",
						followersNumber: "",
						followingNumber: "",
					},
				},
			};
		case userConstants.GET_USER_PROFILE_SUCCESS:
			return {
				...state,
				userProfile: {
					showedUserId: action.userId,
					user: action.user,
				},
			};

		case userConstants.GET_USER_PROFILE_FAILURE:
			return state;

		case modalConstants.HIDE_FOLLOWING_MODAL:
			return {
				...state,
				userProfileFollowingModal: {
					showModal: false,
					userInfos: [],
					modalHeader: "",
				},
			};

		case userConstants.SET_USER_FOLLOWING_REQUEST:
			a = { ...state };
			a.userProfileFollowingModal.userInfos = [];
			return a;

		case userConstants.SET_USER_FOLLOWING_SUCCESS:
			a = { ...state };
			a.userProfileFollowingModal.userInfos = action.userInfos;
			a.userProfileFollowingModal.showModal = true;
			a.userProfileFollowingModal.modalHeader = action.header;
			return a;
		case userConstants.SET_USER_FOLLOWING_FAILURE:
			a = { ...state };
			a.userProfileFollowingModal.userInfos = [];
			return a;

		case userConstants.FOLLOW_USER_REQUEST:
			return state;

		case userConstants.FOLLOW_USER_SUCCESS:
			a = { ...state };
			let cp = a.userProfileFollowingModal.userInfos.find((info) => info.userInfo.id === action.userId);
			if (cp !== undefined) {
				cp.following = true;
			}

			if (a.userProfile.showedUserId === localStorage.getItem("userId")) {
				a.userProfile.user.followingNumber = a.userProfile.user.followingNumber + 1;
			} else {
				a.userProfile.user.followersNumber = a.userProfile.user.followersNumber + 1;
			}
			a.userProfile.user.following = true;

			return a;
		case userConstants.FOLLOW_USER_FAILURE:
			return state;

		case userConstants.UNFOLLOW_USER_REQUEST:
			return state;

		case userConstants.UNFOLLOW_USER_SUCCESS:
			a = { ...state };
			let cccp = a.userProfileFollowingModal.userInfos.find((info) => info.userInfo.id === action.userId);
			if (cccp !== undefined) {
				cccp.following = false;
			}
			if (a.userProfile.showedUserId === localStorage.getItem("userId")) {
				a.userProfile.user.followingNumber = a.userProfile.user.followingNumber - 1;
			} else {
				a.userProfile.user.followersNumber = a.userProfile.user.followersNumber - 1;
			}
			a.userProfile.user.following = false;

			return a;
		case userConstants.UNFOLLOW_USER_FAILURE:
			return state;

		default:
			return state;
	}
};
