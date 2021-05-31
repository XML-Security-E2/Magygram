import { postConstants } from "../constants/PostConstants";

export const postReducer = (state, action) => {
	switch (action.type) {
		case postConstants.CREATE_POST_REQUEST:
			return state;
		case postConstants.CREATE_POST_SUCCESS:
			return state;
		case postConstants.CREATE_POST_FAILURE:
			return state;
		default:
			return state;
	}
};
