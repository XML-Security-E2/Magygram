import React, { createContext, useReducer } from "react";
import { userReducer } from "../reducers/UserReducer";

export const UserContext = createContext();

const UserContextProvider = (props) => {
	const [userState, dispatch] = useReducer(userReducer, {
		loginFirstError: {
			showError: false,
			errorMessage: "",
		},
		loginSecondError: {
			showError: false,
			errorMessage: "",
		},
		showTwoFactorAuth: false,
		registrationError: {
			isActiveRegistrationTab: false,
			showError: false,
			errorMessage: "",
			showSuccessMessage: false,
			emailAddress: "",
			imageData: "",
		},
		registrationTab: {
			showUserRegistrationTab: true,
			showAgentRegistrationTab: false,
		},
		registrationShowQr: false,
		forgotPasswordLinkError: {
			showError: false,
			errorMessage: "",
			showSuccessMessage: false,
			emailAddress: "",
		},
		resetPassword: {
			showError: false,
			errorMessage: "",
			showSuccessMessage: false,
		},
		blockedUser: {
			showError: false,
			errorMessage: "",
			showSuccessMessage: false,
			emailAddress: "",
		},
		userProfileFollowingModal: {
			showModal: false,
			modalHeader: "",
			userInfos: [],
		},
		userFollowRequests: {
			userInfos: [],
		},
		activeTab:{
            verificationRequestsShow: true,
            contentReportShow: false,
            agentRequestsShow: false,
        },
		campaigns:[],
		campaignOptions: {
			showModal: false,
		},
		editProfile: {
			showError: false,
			errorMessage: "",
			showSuccessMessage: false,
			successMessage: "",
		},
		userProfile: {
			showedUserId: "",
			user: {
				username: "",
				name: "",
				following: "",
				surname: "",
				website: "",
				bio: "",
				email: "",
				gender: "",
				imageUrl: "",
				postNumber: "",
				followersNumber: "",
				followingNumber: "",
				sentFollowRequest: false,
				muted: false,
				blocked: false,
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
		followRecommendationInfo: {
			imageUrl: "",
			name: "",
			surname: "",
			username: "",
			recommendUserInfo: [],
		},
		userReport: {
			showModal: false,
			showError: false,
			errorMessage: "",
			showSuccessMessage: false,
			successMessage: "",
		},
	});

	return <UserContext.Provider value={{ userState, dispatch }}>{props.children}</UserContext.Provider>;
};

export default UserContextProvider;
