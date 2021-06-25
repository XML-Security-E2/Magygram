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
			if (postCopy.LikedBy.find((likedByUserInfo) => likedByUserInfo.Id === action.loggedUser.Id) === undefined) {
				postCopy.LikedBy.push(action.loggedUser);
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

			var newLikedByList = postCopy.LikedBy.filter((likedByUserInfo) => likedByUserInfo.Id !== action.loggedUser.Id);
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
			if (postCopy.DislikedBy.find((dislikedByUserInfo) => dislikedByUserInfo.Id === action.loggedUser.Id) === undefined) {
				postCopy.DislikedBy.push(action.loggedUser);
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
			var newDisikedByList = postCopy.DislikedBy.filter((dislikedByUserInfo) => dislikedByUserInfo.Id !== action.loggedUser.Id);
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

		case postConstants.PROFILE_POST_DETAILS_REQUEST:
			strcpy = {
				...state,
			};
			strcpy.userProfileContent.showError = false;
			strcpy.userProfileContent.errorMessage = "";
			strcpy.viewPostModal.showModal = false;
			strcpy.viewPostModal.post = {
				Id: "",
				Description: "",
				Location: "",
				ContentType: "",
				Tags: null,
				HashTags: null,
				Media: [{}],
				UserInfo: {},
				LikedBy: [{}],
				DislikedBy: [{}],
				Comments: [{}],
				Liked: false,
				Disliked: false,
			};
			return strcpy;

		case postConstants.PROFILE_POST_DETAILS_SUCCESS:
			strcpy = {
				...state,
			};
			strcpy.userProfileContent.showError = false;
			strcpy.userProfileContent.errorMessage = "";
			strcpy.viewPostModal.showModal = true;
			strcpy.viewPostModal.post = action.post;
			return strcpy;
		case postConstants.PROFILE_POST_DETAILS_FAILURE:
			strcpy = {
				...state,
			};
			strcpy.userProfileContent.showError = true;
			strcpy.userProfileContent.errorMessage = action.errorMessage;
			strcpy.viewPostModal.showModal = false;
			strcpy.viewPostModal.post = {
				Id: "",
				Description: "",
				Location: "",
				ContentType: "",
				Tags: null,
				HashTags: null,
				Media: [{}],
				UserInfo: {},
				LikedBy: [{}],
				DislikedBy: [{}],
				Comments: [{}],
				Liked: false,
				Disliked: false,
			};
			return strcpy;

		case postConstants.SET_USER_PROFILE_COLLECTIONS_POSTS_REQUEST:
			return {
				...state,
				userProfileContent: {
					showError: false,
					errorMessage: "",
					showPosts: false,
					showCollections: false,
					showCollectionPosts: false,
					showUnauthorizedErrorMessage: false,
					selectedCollectionName: "",
					posts: [],
					collections: [],
					collectionPosts: [],
				},
			};
		case postConstants.SET_USER_PROFILE_COLLECTIONS_POSTS_SUCCESS:
			return {
				...state,
				userProfileContent: {
					showError: false,
					errorMessage: "",
					showPosts: false,
					showCollections: false,
					showCollectionPosts: true,
					showUnauthorizedErrorMessage: false,
					selectedCollectionName: action.collectionName,
					posts: [],
					collections: [],
					collectionPosts: action.collectionPosts,
				},
			};
		case postConstants.SET_USER_PROFILE_COLLECTIONS_POSTS_FAILURE:
			return {
				...state,
				userProfileContent: {
					showError: true,
					errorMessage: action.errorMessage,
					showPosts: false,
					showCollections: false,
					showCollectionPosts: false,
					showUnauthorizedErrorMessage: false,
					selectedCollectionName: "",
					posts: [],
					collections: [],
					collectionPosts: [],
				},
			};

		case postConstants.SET_USER_PROFILE_COLLECTIONS_REQUEST:
			return {
				...state,
				userProfileContent: {
					showError: false,
					errorMessage: "",
					showPosts: false,
					showCollections: false,
					showCollectionPosts: false,
					showUnauthorizedErrorMessage: false,
					selectedCollectionName: "",
					posts: [],
					collections: [],
					collectionPosts: [],
				},
			};
		case postConstants.SET_USER_PROFILE_COLLECTIONS_SUCCESS:
			return {
				...state,
				userProfileContent: {
					showError: false,
					errorMessage: "",
					showPosts: false,
					showCollections: true,
					showCollectionPosts: false,
					showUnauthorizedErrorMessage: false,
					selectedCollectionName: "",
					posts: [],
					collections: action.collections,
					collectionPosts: [],
				},
			};
		case postConstants.SET_USER_PROFILE_COLLECTIONS_FAILURE:
			return {
				...state,
				userProfileContent: {
					showError: true,
					errorMessage: action.errorMessage,
					showPosts: false,
					showCollections: false,
					showCollectionPosts: false,
					showUnauthorizedErrorMessage: false,
					selectedCollectionName: "",
					posts: [],
					collections: [],
					collectionPosts: [],
				},
			};

		case postConstants.SET_USER_POSTS_REQUEST:
			return {
				...state,
				userProfileContent: {
					showError: false,
					errorMessage: "",
					showPosts: false,
					showCollections: false,
					showCollectionPosts: false,
					showUnauthorizedErrorMessage: false,
					selectedCollectionName: "",
					posts: [],
					collections: [],
					collectionPosts: [],
				},
			};
		case postConstants.SET_USER_POSTS_SUCCESS:
			return {
				...state,
				userProfileContent: {
					showError: false,
					errorMessage: "",
					showPosts: true,
					showCollections: false,
					showCollectionPosts: false,
					showUnauthorizedErrorMessage: false,
					selectedCollectionName: "",
					posts: action.posts,
					collections: [],
					collectionPosts: [],
				},
			};
		case postConstants.SET_USER_POSTS_FAILURE:
			return {
				...state,
				userProfileContent: {
					showError: true,
					errorMessage: action.errorMessage,
					showPosts: false,
					showCollections: false,
					showCollectionPosts: false,
					showUnauthorizedErrorMessage: false,
					selectedCollectionName: "",
					posts: [],
					collections: [],
					collectionPosts: [],
				},
			};
		case postConstants.SET_USER_POSTS_UNAUTHORIZED_FAILURE:
			return {
				...state,
				userProfileContent: {
					showError: true,
					errorMessage: action.errorMessage,
					showPosts: false,
					showCollections: false,
					showCollectionPosts: false,
					showUnauthorizedErrorMessage: true,
					selectedCollectionName: "",
					posts: [],
					collections: [],
					collectionPosts: [],
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
				if (postCopy !== undefined) {
					postCopy.Favourites = true;
				} else {
					strcpy.viewPostModal.post.Favourites = true;
				}
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
			if (postCopy !== undefined) {
				postCopy.Favourites = false;
			} else {
				strcpy.viewPostModal.post.Favourites = false;
			}

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

			if (strcpy.viewPostModal.showModal) {
				if (strcpy.viewPostModal.post.Comments.find((comment) => comment.Id === action.comment.Id) === undefined) {
					strcpy.viewPostModal.post.Comments.push(action.comment);
				}
			} else {
				postCopy = strcpy.timeline.posts.find((post) => post.Id === action.postId);

				if (postCopy.Comments.find((comment) => comment.Id === action.comment.Id) === undefined) {
					postCopy.Comments.push(action.comment);
				}
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
		case modalConstants.HIDE_VIEW_POST_FOR_GUEST_MODAL:
			return {
				...state,
				viewPostModalForGuest: {
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
						userId: action.post.UserInfo.Id,
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
			strcpy.viewPostModal.showModal = false;
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
		case postConstants.GUEST_TIMELINE_POSTS_REQUEST:
			return {
				...state,
				guestTimeline: {
					posts: [],
				},
			};
		case postConstants.GUEST_TIMELINE_POSTS_SUCCESS:
			return {
				...state,
				guestTimeline: {
					posts: action.posts,
				},
			};
		case postConstants.GUEST_TIMELINE_POSTS_FAILURE:
			return {
				...state,
				guestTimeline: {
					posts: [],
				},
			};
		case postConstants.PROFILE_POST_DETAILS_FOR_GUEST_REQUEST:
			strcpy = {
				...state,
			};
			strcpy.userProfileContent.showError = false;
			strcpy.userProfileContent.errorMessage = "";
			strcpy.viewPostModalForGuest.showModal = false;
			strcpy.viewPostModalForGuest.post = {
				Id: "",
				Description: "",
				Location: "",
				Media: [{}],
				UserInfo: {},
			};
			return strcpy;

		case postConstants.PROFILE_POST_DETAILS_FOR_GUEST_SUCCESS:
			strcpy = {
				...state,
			};
			strcpy.userProfileContent.showError = false;
			strcpy.userProfileContent.errorMessage = "";
			strcpy.viewPostModalForGuest.showModal = true;
			strcpy.viewPostModalForGuest.post = action.post;
			return strcpy;
		case postConstants.PROFILE_POST_DETAILS_FOR_GUEST_FAILURE:
			strcpy = {
				...state,
			};
			strcpy.userProfileContent.showError = true;
			strcpy.userProfileContent.errorMessage = action.errorMessage;
			strcpy.viewPostModalForGuest.showModal = false;
			strcpy.viewPostModalForGuest.post = {
				Id: "",
				Description: "",
				Location: "",
				Media: [{}],
				UserInfo: {},
			};
			return strcpy;
		case postConstants.LIKED_POSTS_REQUEST:
			return {
				...state,
				userLikedPosts: null,
			};
		case postConstants.LIKED_POSTS_SUCCESS:
			return {
				...state,
				userLikedPosts: action.posts,
			};
		case postConstants.LIKED_POSTS_FAILURE:
			return {
				...state,
				userLikedPosts: null,
			};
		case postConstants.DISLIKED_POSTS_REQUEST:
			return {
				...state,
				userDislikedPosts: null,
			};
		case postConstants.DISLIKED_POSTS_SUCCESS:
			return {
				...state,
				userDislikedPosts: action.posts,
			};
		case postConstants.DISLIKED_POSTS_FAILURE:
			return {
				...state,
				userDislikedPosts: null,
			};

		case postConstants.REPORT_POST_REQUEST:
			postCopy = { ...state };
			postCopy.postReport.showError = false;
			postCopy.postReport.errorMessage = "";
			postCopy.postReport.showSuccessMessage = false;
			postCopy.postReport.successMessage = "";

			return postCopy;
		case postConstants.REPORT_POST_SUCCESS:
			postCopy = { ...state };
			postCopy.postReport.showError = false;
			postCopy.postReport.errorMessage = "";
			postCopy.postReport.showSuccessMessage = true;
			postCopy.postReport.successMessage = action.successMessage;

			return postCopy;

		case postConstants.REPORT_POST_FAILURE:
			postCopy = { ...state };
			postCopy.postReport.showError = true;
			postCopy.postReport.errorMessage = action.errorMessage;
			postCopy.postReport.showSuccessMessage = false;
			postCopy.postReport.successMessage = "";

			return postCopy;

		case postConstants.SET_POST_FOR_PAGE_REQUEST:
			return state;
		case postConstants.SET_POST_FOR_PAGE_SUCCESS:
			postCopy = { ...state };
			postCopy.postDetailsPage.post = action.post;
			return postCopy;

		case postConstants.SET_POST_FOR_PAGE_FAILURE:
			return state;
		default:
			return state;
	}
};
