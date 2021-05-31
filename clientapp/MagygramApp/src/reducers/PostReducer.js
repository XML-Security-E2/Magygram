import { postConstants } from "../constants/PostConstants";

export const postReducer = (state, action) => {
	switch (action.type) {
		case postConstants.TIMELINE_POSTS_REQUEST:
			return {
				
			};
		case postConstants.TIMELINE_POSTS_SUCCESS:
			return {
				
			};
		case postConstants.TIMELINE_POSTS_FAILURE:
			return {
				
			};
		default:
			return state;
	}
};
