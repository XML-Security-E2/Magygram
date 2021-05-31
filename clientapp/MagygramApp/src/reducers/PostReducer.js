import { postConstants } from "../constants/PostConstants";

export const postReducer = (state, action) => {
	switch (action.type) {
		case postConstants.REGISTER_REQUEST:
			return {
				registrationError: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					emailAddress: "",
				},
			};
		default:
			return state;
	}
};
