import { postConstants } from "../constants/PostConstants";

export const postReducer = (state, action) => {
	switch (action.type) {
		case postConstants.CREATE_POST_REQUEST:
			return {
				...state,
				createPost: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					successMessage: "",
				},
			};
		case postConstants.CREATE_POST_SUCCESS:
			return {
				...state,
				createPost: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: true,
					successMessage: action.successMessage,
				},
			};
		case postConstants.CREATE_POST_FAILURE:
			return {
				...state,
				createPost: {
					showError: true,
					errorMessage: action.errorMessage,
					showSuccessMessage: false,
					successMessage: "",
				},
			};
		default:
			return state;
	}
};
