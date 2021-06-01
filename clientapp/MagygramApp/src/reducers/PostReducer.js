import { postConstants } from "../constants/PostConstants";

let strcpy = {}
let postCopy = {}

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
			strcpy =  {
				...state,
			};
			postCopy = strcpy.timeline.posts.find(post => post.Id === action.postId);
			postCopy.Liked= true;
			if (postCopy.LikedBy.find(likedByUserInfo => likedByUserInfo.Id === state.loggedUserInfo.Id) === undefined){
				postCopy.LikedBy.push(state.loggedUserInfo)
			}
			return strcpy;
		case postConstants.LIKE_POST_FAILURE:
			return {
				...state,
			};
		case postConstants.UNLIKE_POST_REQUEST:
			return {
				...state,
			};
		case postConstants.UNLIKE_POST_SUCCESS:
			strcpy =  {
				...state,
			};
			postCopy = strcpy.timeline.posts.find(post => post.Id === action.postId);
			postCopy.Liked= false;
			var newLikedByList = postCopy.LikedBy.filter((likedByUserInfo) => likedByUserInfo.Id !== state.loggedUserInfo.Id);
			postCopy.LikedBy = newLikedByList
			return strcpy;
		case postConstants.UNLIKE_POST_FAILURE:
			return {
				...state,
			};
		case postConstants.DISLIKE_POST_REQUEST:
			return {
				...state,
			};
		case postConstants.DISLIKE_POST_SUCCESS:
			strcpy =  {
				...state,
			};
			postCopy = strcpy.timeline.posts.find(post => post.Id === action.postId);
			postCopy.Disliked= true;
			if (postCopy.DislikedBy.find(dislikedByUserInfo => dislikedByUserInfo.Id === state.loggedUserInfo.Id) === undefined){
				postCopy.DislikedBy.push(state.loggedUserInfo)
			}
			return strcpy;
		case postConstants.DISLIKE_POST_FAILURE:
			return {
				...state,
			};
		case postConstants.UNDISLIKE_POST_REQUEST:
			return {
				...state,
			};
		case postConstants.UNDISLIKE_POST_SUCCESS:
			strcpy =  {
				...state,
			};
			postCopy = strcpy.timeline.posts.find(post => post.Id === action.postId);
			postCopy.Disliked= false;
			var newDisikedByList = postCopy.DislikedBy.filter((dislikedByUserInfo) => dislikedByUserInfo.Id !== state.loggedUserInfo.Id);
			postCopy.DislikedBy = newDisikedByList
			return strcpy;
		case postConstants.UNDISLIKE_POST_FAILURE:
			return {
				...state,
			};
		default:
			return state;
	}
};
