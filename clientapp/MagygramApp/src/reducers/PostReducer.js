import { postConstants } from "../constants/PostConstants";
import { modalConstants } from "../constants/ModalConstants";

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
		case modalConstants.SHOW_POST_LIKED_BY_DETAILS:
			return {
				...state,
				postLikedBy:{
					showModal: true,
					likedBy: action.LikedBy
				}
			};
		case modalConstants.HIDE_POST_LIKED_BY_DETAILS:
			return {
				...state,
				postLikedBy:{
					showModal: false,
					likedBy: []
				}
			};
		case modalConstants.SHOW_POST_DISLIKES_MODAL:
			return {
				...state,
				postDislikes:{
					showModal: true,
					dislikes: action.Dislikes
				}
			};
		case modalConstants.HIDE_POST_DISLIKES_MODAL:
			return {
				...state,
				postDislikes:{
					showModal: false,
					dislikes: []
				}
			};
		case postConstants.COMMENT_POST_REQUEST:
			return {
				...state,
			};
		case postConstants.COMMENT_POST_SUCCESS:
			strcpy =  {
				...state,
			};

			postCopy = strcpy.timeline.posts.find(post => post.Id === action.postId);

			if (postCopy.Comments.find(comment => comment.Id === action.comment.Id) === undefined){
				postCopy.Comments.push(action.comment)
			}
			
			return strcpy;
		case postConstants.COMMENT_POST_FAILURE:
			return {
				...state,
			};
		case modalConstants.SHOW_VIEW_POST_MODAL:
			return {
				...state,
				viewPostModal:{
					showModal: true,
					post: action.post
				}
			};
		case modalConstants.HIDE_VIEW_POST_MODAL:
			return {
				...state,
				viewPostModal: {
					showModal:false,
					post:action.post
				}
			};
		default:
			return state;
	}
};
