import React, { createContext, useReducer } from "react";
import { userReducer } from "../reducers/UserReducer";

export const UserContext = createContext();

const UserContextProvider = (props) => {
	const [userState, dispatch] = useReducer(userReducer, {
		loginError: {
			showError: false,
			errorMessage: "",
		},
		registrationError: {
			showError: false,
			errorMessage: "",
			showSuccessMessage: false,
			emailAddress: "",
		},
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
			},
		},
	});

	return <UserContext.Provider value={{ userState, dispatch }}>{props.children}</UserContext.Provider>;
};

export default UserContextProvider;
