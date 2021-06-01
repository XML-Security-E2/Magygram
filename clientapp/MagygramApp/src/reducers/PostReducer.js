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
		case postConstants.TIMELINE_POSTS_REQUEST:
			return {
				...state,
				timeline: {		
					posts : []
				}
			};
		case postConstants.TIMELINE_POSTS_SUCCESS:
			return {
				...state,
				timeline: {		
					posts : action.posts
				}
			};
		case postConstants.TIMELINE_POSTS_FAILURE:
			return {
				...state,
				timeline: {		
					posts : []
				}
			};
		case postConstants.LIKE_POST_REQUEST:
			return {
				...state,
			};
		case postConstants.LIKE_POST_SUCCESS:
			alert('t')
			let strcpy =  {
				...state,
			};
			strcpy.timeline.posts.find(post => post.Id === action.postId).Liked = true;
			return strcpy;
		case postConstants.LIKE_POST_FAILURE:
			return {
				...state,
			};
		default:
			return state;
	}
};
