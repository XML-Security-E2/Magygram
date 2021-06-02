import { modalConstants } from "../constants/ModalConstants";
import { postConstants } from "../constants/PostConstants";

let strcpy = {};
let postCopy = {};

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
					posts: [],
				},
			};
		case postConstants.TIMELINE_POSTS_SUCCESS:
			return {
				...state,
				timeline: {
					posts: action.posts,
				},
			};
		case postConstants.TIMELINE_POSTS_FAILURE:
			return {
				...state,
				timeline: {
					posts: [],
				},
			};
		case postConstants.LIKE_POST_REQUEST:
			return {
				...state,
			};
		case postConstants.LIKE_POST_SUCCESS:
			strcpy = {
				...state,
			};
			postCopy = strcpy.timeline.posts.find((post) => post.Id === action.postId);
			postCopy.Liked = true;
			if (postCopy.LikedBy.find((likedByUserInfo) => likedByUserInfo.Id === state.loggedUserInfo.Id) === undefined) {
				postCopy.LikedBy.push(state.loggedUserInfo);
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
			strcpy = {
				...state,
			};
			postCopy = strcpy.timeline.posts.find((post) => post.Id === action.postId);
			postCopy.Liked = false;
			var newLikedByList = postCopy.LikedBy.filter((likedByUserInfo) => likedByUserInfo.Id !== state.loggedUserInfo.Id);
			postCopy.LikedBy = newLikedByList;
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
			strcpy = {
				...state,
			};
			postCopy = strcpy.timeline.posts.find((post) => post.Id === action.postId);
			postCopy.Disliked = true;
			if (postCopy.DislikedBy.find((dislikedByUserInfo) => dislikedByUserInfo.Id === state.loggedUserInfo.Id) === undefined) {
				postCopy.DislikedBy.push(state.loggedUserInfo);
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
			strcpy = {
				...state,
			};
			postCopy = strcpy.timeline.posts.find((post) => post.Id === action.postId);
			postCopy.Disliked = false;
			var newDisikedByList = postCopy.DislikedBy.filter((dislikedByUserInfo) => dislikedByUserInfo.Id !== state.loggedUserInfo.Id);
			postCopy.DislikedBy = newDisikedByList;
			return strcpy;
		case postConstants.UNDISLIKE_POST_FAILURE:
			return {
				...state,
			};

		case postConstants.SET_USER_COLLECTIONS_REQUEST:
			return {
				...state,
				userCollections: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					successMessage: "",
					collections: [],
				},
			};
		case postConstants.SET_USER_COLLECTIONS_SUCCESS:
			return {
				...state,
				userCollections: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					successMessage: "",
					collections: action.collections,
				},
			};
		case postConstants.SET_USER_COLLECTIONS_FAILURE:
			return {
				...state,
				userCollections: {
					showError: true,
					errorMessage: action.errorMessage,
					showSuccessMessage: false,
					successMessage: "",
					collections: [],
				},
			};
		case modalConstants.OPEN_ADD_TO_COLLECTION_MODAL:
			return {
				...state,
				addToFavouritesModa: {
					renderCollectionSwitch: !state.addToFavouritesModa.renderCollectionSwitch,
					showModal: true,
					selectedPostId: action.postId,
				},
			};
		case modalConstants.CLOSE_ADD_TO_COLLECTION_MODAL:
			strcpy = {
				...state,
			};
			strcpy.addToFavouritesModa.showModal = false;
			strcpy.addToFavouritesModa.selectedPostId = "";

			return strcpy;

		case postConstants.ADD_POST_TO_COLLECTION_REQUEST:
			strcpy = {
				...state,
			};
			strcpy.userCollections.showError = false;
			strcpy.userCollections.errorMessage = "";
			strcpy.userCollections.showSuccessMessage = false;
			strcpy.userCollections.successMessage = "";
			return strcpy;

		case postConstants.ADD_POST_TO_COLLECTION_SUCCESS:
			strcpy = {
				...state,
			};
			strcpy.userCollections.showError = false;
			strcpy.userCollections.errorMessage = "";
			strcpy.userCollections.showSuccessMessage = true;
			strcpy.userCollections.successMessage = action.successMessage;
			return strcpy;
		case postConstants.ADD_POST_TO_COLLECTION_FAILURE:
			strcpy = {
				...state,
			};
			strcpy.userCollections.showError = true;
			strcpy.userCollections.errorMessage = action.errorMessage;
			strcpy.userCollections.showSuccessMessage = false;
			strcpy.userCollections.successMessage = "";
			return strcpy;
		default:
			return state;
	}
};
