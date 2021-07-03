import { modalConstants } from "../constants/ModalConstants";
import { userConstants } from "../constants/UserConstants";

var a = {};

export const userReducer = (state, action) => {
	switch (action.type) {
		case userConstants.REGISTER_REQUEST:
			return {
				...state,
				registrationError: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					emailAddress: "",
				},
			};
		case userConstants.REGISTER_SUCCESS:
			return {
				...state,
				registrationError: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: true,
					emailAddress: action.emailAddress,
					imageData: action.imageData,
				},
			};
		case userConstants.REGISTER_AGENT_SUCCESS:
			return {
				...state,
				registrationError: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: true,
					emailAddress: action.emailAddress,
					imageData: "",
				},
			};
		case userConstants.REGISTER_FAILURE:
			return {
				...state,
				registrationError: {
					showError: true,
					errorMessage: action.errorMessage,
					showSuccessMessage: false,
					emailAddress: "",
				},
			};
		case userConstants.REGISTRATION_SHOW_QR_CODE:
			return {
				...state,
				registrationShowQr: true,
			};
		case userConstants.REGISTER_VALIDATION_FAILURE:
			return {
				...state,
				registrationError: {
					showError: true,
					errorMessage: action.errorMessage,
					showSuccessMessage: false,
					emailAddress: "",
				},
			};
		case userConstants.LOGIN_REQUEST:
			return {
				...state,
				loginFirstError: {
					showError: false,
					errorMessage: "",
				},
				loginSecondError: {
					showError: false,
					errorMessage: "",
				},
				showTwoFactorAuth: false,
			};
		case userConstants.SHOW_INFLUENCER_CAMPAIGN_TAB: {
			return {
				...state,
				activeTab: {
					verificationRequestsShow: false,
					contentReportShow: true,
					agentRequestsShow: false,
				},
			};
		}
		case userConstants.HIDE_INFLUENCER_CAMPAIGN_TAB: {
			return {
				...state,
				activeTab: {
					verificationRequestsShow: false,
					contentReportShow: false,
					agentRequestsShow: true,
				},
			};
		}
		case userConstants.LOGIN_FAILURE:
			return {
				...state,
				loginFirstError: {
					showError: true,
					errorMessage: action.error,
				},
				loginSecondError: {
					showError: false,
					errorMessage: "",
				},
				showTwoFactorAuth: false,
			};
		case userConstants.LOGIN_SUCCESS:
			return {
				...state,
				loginFirstError: {
					showError: false,
					errorMessage: "",
				},
				loginSecondError: {
					showError: false,
					errorMessage: "",
				},
				showTwoFactorAuth: false, // edit na true za aktiviranje 2fa
			};
		case userConstants.LOGIN_TWO_REQUEST:
			return {
				...state,
				loginFirstError: {
					showError: false,
					errorMessage: "",
				},
				loginSecondError: {
					showError: false,
					errorMessage: "",
				},
				showTwoFactorAuth: true,
			};
		case userConstants.LOGIN_TWO_FAILURE:
			return {
				...state,
				loginSecondError: {
					showError: true,
					errorMessage: action.error,
				},
				showTwoFactorAuth: true,
			};
		case userConstants.LOGIN_TWO_SUCCESS:
			return {
				...state,
				loginFirstError: {
					showError: false,
					errorMessage: "",
				},
				loginSecondError: {
					showError: false,
					errorMessage: "",
				},
				showTwoFactorAuth: true,
			};
		case userConstants.LOGIN_DATA_REQUEST:
			return {
				...state,
				loginFirstError: {
					showError: false,
					errorMessage: "",
				},
				loginSecondError: {
					showError: false,
					errorMessage: "",
				},
				showTwoFactorAuth: false, // edit na true za aktiviranje 2fa
			};
		case userConstants.LOGIN_DATA_FAILURE:
			return {
				...state,
				loginSecondError: {
					showError: true,
					errorMessage: action.error,
				},
				showTwoFactorAuth: false, // edit na true za aktiviranje 2fa
			};
		case userConstants.LOGIN_DATA_SUCCESS:
			return {
				...state,
				loginFirstError: {
					showError: false,
					errorMessage: "",
				},
				loginSecondError: {
					showError: false,
					errorMessage: "",
				},
				showTwoFactorAuth: false, // edit na true za aktiviranje 2fa
			};
		case userConstants.RESET_PASSWORD_LINK_REQUEST:
			return {
				...state,
				forgotPasswordLinkError: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					emailAddress: "",
				},
			};
		case userConstants.RESET_PASSWORD_LINK_SUCCESS:
			return {
				...state,
				forgotPasswordLinkError: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: true,
					emailAddress: action.emailAddress,
				},
			};
		case userConstants.RESET_PASSWORD_LINK_FAILURE:
			return {
				...state,
				forgotPasswordLinkError: {
					showError: true,
					errorMessage: action.errorMessage,
					showSuccessMessage: false,
					emailAddress: "",
				},
			};

		case userConstants.RESET_PASSWORD_REQUEST:
			return {
				...state,
				resetPassword: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
				},
			};
		case userConstants.RESET_PASSWORD_SUCCESS:
			return {
				...state,
				resetPassword: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: true,
				},
			};
		case userConstants.RESET_PASSWORD_FAILURE:
			return {
				...state,
				resetPassword: {
					showError: true,
					errorMessage: action.errorMessage,
					showSuccessMessage: false,
				},
			};
		case userConstants.RESEND_ACTIVATION_LINK_REQUEST:
			return {
				...state,
				blockedUser: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
				},
			};
		case userConstants.RESEND_ACTIVATION_LINK_SUCCESS:
			return {
				...state,
				blockedUser: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: true,
				},
			};
		case userConstants.RESEND_ACTIVATION_LINK_FAILURE:
			return {
				...state,
				blockedUser: {
					showError: true,
					errorMessage: action.errorMessage,
					showSuccessMessage: false,
				},
			};
		case userConstants.BLOCKED_USER_EMAIL_REQUEST:
			return {
				...state,
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
						gender: "MALE",
						imageUrl: "",
						postNumber: "",
						followersNumber: "",
						followingNumber: "",
						sentFollowRequest: false,
						blocked: false,
						birthDate: new Date(),
						notificationSettings: {
							notifyStory: false,
							notifyPost: false,
							notifyLike: false,
							notifyDislike: false,
							notifyFollow: false,
							notifyFollowRequest: false,
							notifyAcceptFollowRequest: false,
							notifyComments: false,
						},
						privacySettings: {
							isPrivate: false,
							receiveMessages: true,
							isTaggable: true,
						},
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
		case userConstants.SET_USER_CAMPAIGNS_REQUEST: {
			let strCpy = {
				...state,
			};
			strCpy.campaigns = [];
			return strCpy;
		}
		case userConstants.SET_USER_CAMPAIGNS: {
			let strCpy = {
				...state,
			};
			strCpy.campaigns = action.campaigns;
			return strCpy;
		}

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
			a.userProfile.user.sentFollowRequest = false;

			return a;

		case userConstants.FOLLOW_USER_SEND_REQUEST_SUCCESS:
			a = { ...state };
			a.userProfile.user.sentFollowRequest = true;
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

		case userConstants.UPDATE_USER_REQUEST:
			return {
				...state,
				editProfile: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					successMessage: "",
				},
			};

		case userConstants.UPDATE_USER_SUCCESS:
			return {
				...state,
				editProfile: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: true,
					successMessage: action.successMessage,
				},
			};
		case userConstants.UPDATE_USER_FAILURE:
			return {
				...state,
				editProfile: {
					showError: true,
					errorMessage: action.errorMessage,
					showSuccessMessage: false,
					successMessage: "",
				},
			};

		case userConstants.FOLLOW_REQUESTS_REQUEST:
			return {
				...state,
				userFollowRequests: {
					userInfos: [],
				},
			};

		case userConstants.FOLLOW_REQUESTS_SUCCESS:
			return {
				...state,
				userFollowRequests: {
					userInfos: action.userInfos,
				},
			};
		case userConstants.FOLLOW_REQUESTS_FAILURE:
			return {
				...state,
				userFollowRequests: {
					userInfos: [],
				},
			};

		case userConstants.ACCEPT_FOLLOW_REQUESTS_REQUEST:
			return state;

		case userConstants.ACCEPT_FOLLOW_REQUESTS_SUCCESS:
			a = { ...state };

			let usrCpy = a.userFollowRequests.userInfos.find((uinfo) => uinfo.userInfo.id === action.userId);
			usrCpy.following = true;
			console.log(action.userInfos);
			return a;
		case userConstants.ACCEPT_FOLLOW_REQUESTS_FAILURE:
			return state;
		case userConstants.MUTE_USER_SUCCESS:
			a = { ...state };
			a.userProfile.user.muted = true;
			return a;
		case userConstants.UNMUTE_USER_SUCCESS:
			a = { ...state };
			a.userProfile.user.muted = false;
			return a;
		case userConstants.BLOCK_USER_SUCCESS:
			a = { ...state };
			a.userProfile.user.blocked = true;
			a.userProfile.user.muted = false;
			a.userProfile.user.following = false;
			return a;
		case userConstants.UNBLOCK_USER_SUCCESS:
			a = { ...state };
			a.userProfile.user.blocked = false;
			return a;
		case userConstants.BLOCK_USER_REQUEST:
			return state;
		case userConstants.UNBLOCK_USER_REQUEST:
			return state;
		case userConstants.GET_FOLLOW_RECOMMENDATION_REQUEST:
			return {
				...state,
				followRecommendationInfo: {
					imageUrl: "",
					name: "",
					surname: "",
					username: "",
					recommendUserInfo: [],
				},
			};
		case userConstants.GET_FOLLOW_RECOMMENDATION_FAILURE:
			return {
				...state,
				followRecommendationInfo: {
					imageUrl: "",
					name: "",
					surname: "",
					username: "",
					recommendUserInfo: [],
				},
			};
		case userConstants.GET_FOLLOW_RECOMMENDATION_SUCCESS:
			return {
				...state,
				followRecommendationInfo: {
					imageUrl: action.data.ImageURL,
					name: action.data.Name,
					surname: action.data.Surname,
					username: action.data.Username,
					recommendUserInfo: action.data.RecommendedUsers,
				},
			};

		case userConstants.RECOMMENDED_FOLLOW_USER_REQUEST:
			return state;
		case userConstants.RECOMMENDED_FOLLOW_USER_SUCCESS:
			a = { ...state };
			let recommendFollowSuccess = a.followRecommendationInfo.recommendUserInfo.find((info) => info.Id === action.userId);
			if (recommendFollowSuccess !== undefined) {
				recommendFollowSuccess.Followed = true;
			}

			return a;

		case userConstants.RECOMMENDED_FOLLOW_USER_SEND_REQUEST_SUCCESS:
			a = { ...state };
			let recommendFollow = a.followRecommendationInfo.recommendUserInfo.find((info) => info.Id === action.userId);
			if (recommendFollow !== undefined) {
				recommendFollow.SendedRequest = true;
			}

			return a;
		case userConstants.RECOMMENDED_FOLLOW_USER_FAILURE:
			return state;
		case userConstants.SHOW_AGENT_REGISTRATION_TAB:
			return {
				...state,
				registrationTab: {
					showUserRegistrationTab: false,
					showAgentRegistrationTab: true,
				},
				registrationError: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					emailAddress: "",
				},
			};
		case userConstants.SHOW_USER_REGISTRATION_TAB:
			return {
				...state,
				registrationTab: {
					showUserRegistrationTab: true,
					showAgentRegistrationTab: false,
				},
				registrationError: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					emailAddress: "",
				},
			};

		case modalConstants.SHOW_USER_REPORT_MODAL:
			return {
				...state,
				userReport: {
					showModal: true,
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					successMessage: "",
				},
			};
		case modalConstants.HIDE_USER_REPORT_MODAL:
			return {
				...state,
				userReport: {
					showModal: false,
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					successMessage: "",
				},
			};

		case userConstants.REPORT_USER_REQUEST:
			a = { ...state };
			a.userReport.showError = false;
			a.userReport.errorMessage = "";
			a.userReport.showSuccessMessage = false;
			a.userReport.successMessage = "";

			return a;
		case userConstants.REPORT_USER_SUCCESS:
			a = { ...state };
			a.userReport.showError = false;
			a.userReport.errorMessage = "";
			a.userReport.showSuccessMessage = true;
			a.userReport.successMessage = action.successMessage;

			return a;

		case userConstants.REPORT_USER_FAILURE:
			a = { ...state };
			a.userReport.showError = true;
			a.userReport.errorMessage = action.errorMessage;
			a.userReport.showSuccessMessage = false;
			a.userReport.successMessage = "";

			return a;
		default:
			return state;
	}
};
