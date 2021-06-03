import { modalConstants } from "../constants/ModalConstants";
import { postConstants } from "../constants/PostConstants";

let strcpy = {};
let postCopy = {};

export const postReducer = (state, action) => {
	switch (action.type) {
		case postConstants.EDIT_POST_REQUEST:
			strcpy = {
				...state,
			};
			strcpy.editPost.showError = false;
			strcpy.editPost.errorMessage = "";
			strcpy.editPost.showSuccessMessage = false;
			strcpy.editPost.successMessage = "";
			return strcpy;

		case postConstants.EDIT_POST_SUCCESS:
			strcpy = {
				...state,
			};

			postCopy = strcpy.timeline.posts.find((post) => post.Id === action.post.id);
			postCopy.Location = action.post.location;
			postCopy.Description = action.post.description;
			postCopy.Tags = action.post.tags;

			strcpy.editPost.showError = false;
			strcpy.editPost.errorMessage = "";
			strcpy.editPost.showSuccessMessage = true;
			strcpy.editPost.successMessage = action.successMessage;
			return strcpy;
		case postConstants.EDIT_POST_FAILURE:
			strcpy = {
				...state,
			};
			strcpy.editPost.showError = true;
			strcpy.editPost.errorMessage = action.errorMessage;
			strcpy.editPost.showSuccessMessage = false;
			strcpy.editPost.successMessage = "";
			return strcpy;
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
				addToFavouritesModal: {
					renderCollectionSwitch: !state.addToFavouritesModal.renderCollectionSwitch,
					showModal: true,
					selectedPostId: action.postId,
				},
			};
		case modalConstants.CLOSE_ADD_TO_COLLECTION_MODAL:
			strcpy = {
				...state,
			};
			strcpy.addToFavouritesModal.showModal = false;
			strcpy.addToFavouritesModal.selectedPostId = "";

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
			postCopy = strcpy.timeline.posts.find((post) => post.Id === action.collectionDTO.postId);

			if (action.defaultCollection) {
				postCopy.Favourites = true;
			} else {
				if (strcpy.userCollections.collections[action.collectionDTO.collectionName].find((col) => col.id === action.collectionDTO.postId) === undefined) {
					strcpy.userCollections.collections[action.collectionDTO.collectionName].push({
						id: action.collectionDTO.postId,
						media: { url: postCopy.Media[0].Url, mediaType: postCopy.Media[0].MediaType },
					});
				}
			}

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

		case postConstants.DELETE_POST_FROM_COLLECTION_REQUEST:
			return state;

		case postConstants.DELETE_POST_FROM_COLLECTION_SUCCESS:
			strcpy = {
				...state,
			};

			postCopy = strcpy.timeline.posts.find((post) => post.Id === action.postId);
			postCopy.Favourites = false;

			for (const [key] of Object.entries(strcpy.userCollections.collections)) {
				strcpy.userCollections.collections[key] = strcpy.userCollections.collections[key].filter((collection) => collection.id !== action.postId);
			}

			console.log(strcpy);

			return strcpy;
		case postConstants.DELETE_POST_FROM_COLLECTION_FAILURE:
			return state;

		case postConstants.CREATE_COLLECTION_REQUEST:
			strcpy = {
				...state,
			};
			strcpy.userCollections.showError = false;
			strcpy.userCollections.errorMessage = "";
			strcpy.userCollections.showSuccessMessage = false;
			strcpy.userCollections.successMessage = "";
			return strcpy;

		case postConstants.CREATE_COLLECTION_SUCCESS:
			strcpy = {
				...state,
			};

			strcpy.userCollections.collections[action.collectionName] = [];
			strcpy.userCollections.showError = false;
			strcpy.userCollections.errorMessage = "";
			strcpy.userCollections.showSuccessMessage = true;
			strcpy.userCollections.successMessage = action.successMessage;

			return strcpy;
		case postConstants.CREATE_COLLECTION_FAILURE:
			strcpy = {
				...state,
			};
			strcpy.userCollections.showError = true;
			strcpy.userCollections.errorMessage = action.errorMessage;
			strcpy.userCollections.showSuccessMessage = false;
			strcpy.userCollections.successMessage = "";
			return strcpy;
		case modalConstants.SHOW_POST_LIKED_BY_DETAILS:
			return {
				...state,
				postLikedBy: {
					showModal: true,
					likedBy: action.LikedBy,
				},
			};
		case modalConstants.HIDE_POST_LIKED_BY_DETAILS:
			return {
				...state,
				postLikedBy: {
					showModal: false,
					likedBy: [],
				},
			};
		case modalConstants.SHOW_POST_DISLIKES_MODAL:
			return {
				...state,
				postDislikes: {
					showModal: true,
					dislikes: action.Dislikes,
				},
			};
		case modalConstants.HIDE_POST_DISLIKES_MODAL:
			return {
				...state,
				postDislikes: {
					showModal: false,
					dislikes: [],
				},
			};
		case postConstants.COMMENT_POST_REQUEST:
			return {
				...state,
			};
		case postConstants.COMMENT_POST_SUCCESS:
			strcpy = {
				...state,
			};

			postCopy = strcpy.timeline.posts.find((post) => post.Id === action.postId);

			if (postCopy.Comments.find((comment) => comment.Id === action.comment.Id) === undefined) {
				postCopy.Comments.push(action.comment);
			}

			return strcpy;
		case postConstants.COMMENT_POST_FAILURE:
			return {
				...state,
			};
		case modalConstants.SHOW_VIEW_POST_MODAL:
			return {
				...state,
				viewPostModal: {
					showModal: true,
					post: action.post,
				},
			};
		case modalConstants.HIDE_VIEW_POST_MODAL:
			return {
				...state,
				viewPostModal: {
					showModal: false,
					post: action.post,
				},
			};

		case modalConstants.SHOW_POST_OPTIONS_MODAL:
			console.log(action.post);
			return {
				...state,
				editPost: {
					showModal: false,
					post: {
						id: action.post.Id,
						location: action.post.Location,
						tags: action.post.Tags,
						description: action.post.Description,
						media: action.post.Media,
					},
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					successMessage: "",
				},
				postOptions: {
					showModal: true,
				},
			};
		case modalConstants.HIDE_POST_OPTIONS_MODAL:
			return {
				...state,
				editPost: {
					showModal: false,
					post: {
						id: "",
						location: "",
						tags: [],
						description: "",
						media: [],
					},
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					successMessage: "",
				},
				postOptions: {
					showModal: false,
				},
			};

		case modalConstants.SHOW_POST_EDIT_MODAL:
			strcpy = {
				...state,
			};
			strcpy.editPost.showModal = true;
			strcpy.postOptions.showModal = false;
			return strcpy;
		case modalConstants.HIDE_POST_EDIT_MODAL:
			return {
				...state,
				editPost: {
					showModal: false,
					post: {
						id: "",
						location: "",
						tags: [],
						description: "",
						media: [],
					},
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					successMessage: "",
				},
				postOptions: {
					showModal: false,
				},
			};
		default:
			return state;
	}
};
