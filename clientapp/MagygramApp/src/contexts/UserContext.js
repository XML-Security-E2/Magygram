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
		loggedUser: {
			id:'3b88ea7b-4a1d-43d5-b741-d08c413336a7',
			username: 'nikolakolovic',
			imageURL:'assets/images/profiles/profile-1.jpg'
		}
	});

	return <UserContext.Provider value={{ userState, dispatch }}>{props.children}</UserContext.Provider>;
};

export default UserContextProvider;
