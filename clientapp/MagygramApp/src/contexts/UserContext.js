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
			id:'732c397d-f6a3-4fdd-8388-85a55420bf3a',
			username: 'nikolakolovic',
			imageURL:'assets/images/profiles/profile-1.jpg'
		}
	});

	return <UserContext.Provider value={{ userState, dispatch }}>{props.children}</UserContext.Provider>;
};

export default UserContextProvider;
