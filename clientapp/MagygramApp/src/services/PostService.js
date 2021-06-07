import Axios from "axios";
import { postConstants } from "../constants/PostConstants";
import { authHeader } from "../helpers/auth-header";

export const postService = {
	findPostsForTimeline,
	findPostById,
	findAllUserPosts,
	findAllUsersCollections,
	findAllUsersProfileCollections,
	findAllPostsFromCollection,
	addPostToCollection,
	createCollection,
	deletePostFromCollection,
	createPost,
	editPost,
	likePost,
	unlikePost,
	dislikePost,
	undislikePost,
	commentPost,
	findPostsForGuestByHashtag,
	findPostsForUserByHashtag,
	findPostsForGuestByLocation,
	findPostsForUserByLocation,
};

async function findPostsForTimeline(dispatch) {
	dispatch(request());
	await Axios.get(`/api/posts`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				failure();
			}
		})
		.catch((err) => {
			failure();
		});

	function request() {
		return { type: postConstants.TIMELINE_POSTS_REQUEST };
	}

	function success(data) {
		return { type: postConstants.TIMELINE_POSTS_SUCCESS, posts: data };
	}
	function failure() {
		return { type: postConstants.TIMELINE_POSTS_FAILURE };
	}
}

async function findAllUserPosts(userId, dispatch) {
	dispatch(request());
	await Axios.get(`/api/posts/${userId}`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else if (res.status === 401) {
				dispatch(unauthorizedFailure());
			} else {
				dispatch(failure(res.data.message));
			}
		})
		.catch((err) => {
			dispatch(failure("error"));
		});

	function request() {
		return { type: postConstants.SET_USER_POSTS_REQUEST };
	}

	function success(data) {
		return { type: postConstants.SET_USER_POSTS_SUCCESS, posts: data };
	}
	function failure(error) {
		return { type: postConstants.SET_USER_POSTS_FAILURE, errorMessage: error };
	}
	function unauthorizedFailure() {
		return { type: postConstants.SET_USER_POSTS_UNAUTHORIZED_FAILURE };
	}
}

async function findAllUsersCollections(dispatch) {
	dispatch(request());
	await Axios.get(`/api/users/collections/except-default`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				dispatch(failure("Error while loading collections"));
			}
		})
		.catch((err) => {
			dispatch(failure());
		});

	function request() {
		return { type: postConstants.SET_USER_COLLECTIONS_REQUEST };
	}

	function success(data) {
		return { type: postConstants.SET_USER_COLLECTIONS_SUCCESS, collections: data };
	}
	function failure(message) {
		return { type: postConstants.SET_USER_COLLECTIONS_FAILURE, errorMessage: message };
	}
}

async function findPostById(postId, dispatch) {
	dispatch(request());
	await Axios.get(`/api/posts/id/${postId}`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				dispatch(failure("Error while loading collections"));
			}
		})
		.catch((err) => {
			dispatch(failure());
		});

	function request() {
		return { type: postConstants.PROFILE_POST_DETAILS_REQUEST };
	}

	function success(data) {
		return { type: postConstants.PROFILE_POST_DETAILS_SUCCESS, post: data };
	}
	function failure(message) {
		return { type: postConstants.PROFILE_POST_DETAILS_FAILURE, errorMessage: message };
	}
}

async function findAllPostsFromCollection(collectionName, dispatch) {
	dispatch(request());
	await Axios.get(`/api/users/collections/${collectionName}/posts`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data, collectionName));
			} else {
				dispatch(failure("Error while loading collections"));
			}
		})
		.catch((err) => {
			dispatch(failure("error"));
		});

	function request() {
		return { type: postConstants.SET_USER_PROFILE_COLLECTIONS_POSTS_REQUEST };
	}

	function success(data, collectionName) {
		return { type: postConstants.SET_USER_PROFILE_COLLECTIONS_POSTS_SUCCESS, collectionPosts: data, collectionName };
	}
	function failure(message) {
		return { type: postConstants.SET_USER_PROFILE_COLLECTIONS_POSTS_FAILURE, errorMessage: message };
	}
}

async function findAllUsersProfileCollections(dispatch) {
	dispatch(request());
	await Axios.get(`/api/users/collections`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				failure("Error while loading collections");
			}
		})
		.catch((err) => {
			failure();
		});

	function request() {
		return { type: postConstants.SET_USER_PROFILE_COLLECTIONS_REQUEST };
	}

	function success(data) {
		return { type: postConstants.SET_USER_PROFILE_COLLECTIONS_SUCCESS, collections: data };
	}
	function failure(message) {
		return { type: postConstants.SET_USER_PROFILE_COLLECTIONS_FAILURE, errorMessage: message };
	}
}

function createPost(postDTO, dispatch) {
	const formData = fetchFormData(postDTO);
	dispatch(request());
	console.log(formData);

	Axios.post(`/api/posts`, formData, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);

			if (res.status === 201) {
				dispatch(success("Post successfully created"));
			} else {
				dispatch(failure(res.data.message));
			}
		})
		.catch((err) => {
			console.log(err);
		});

	function request() {
		return { type: postConstants.CREATE_POST_REQUEST };
	}
	function success(message) {
		return { type: postConstants.CREATE_POST_SUCCESS, successMessage: message };
	}
	function failure(message) {
		return { type: postConstants.CREATE_POST_FAILURE, errorMessage: message };
	}
}

function editPost(postDTO, dispatch) {
	dispatch(request());

	Axios.put(`/api/posts`, postDTO, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);
			if (res.status === 200) {
				dispatch(success("Post successfully updated", postDTO));
			} else {
				dispatch(failure(res.data.message));
			}
		})
		.catch((err) => {
			console.log(err);
		});

	function request() {
		return { type: postConstants.EDIT_POST_REQUEST };
	}
	function success(message, post) {
		return { type: postConstants.EDIT_POST_SUCCESS, successMessage: message, post };
	}
	function failure(message) {
		return { type: postConstants.EDIT_POST_FAILURE, errorMessage: message };
	}
}

function addPostToCollection(collectionDTO, dispatch) {
	dispatch(request());

	Axios.post(`/api/users/collections/posts`, collectionDTO, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);
			if (res.status === 201) {
				dispatch(success("Post successfully added to favourites", collectionDTO, collectionDTO.collectionName === ""));
			} else {
				dispatch(failure(res.data.message));
			}
		})
		.catch((err) => {
			console.log(err);
		});

	function request() {
		return { type: postConstants.ADD_POST_TO_COLLECTION_REQUEST };
	}
	function success(message, collectionDTO, defaultCollection) {
		return { type: postConstants.ADD_POST_TO_COLLECTION_SUCCESS, successMessage: message, collectionDTO, defaultCollection };
	}
	function failure(message) {
		return { type: postConstants.ADD_POST_TO_COLLECTION_FAILURE, errorMessage: message };
	}
}

function createCollection(collectionName, dispatch) {
	let formData = new FormData();
	formData.append("name", collectionName);

	dispatch(request());

	Axios.post(`/api/users/collections`, formData, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);
			if (res.status === 201) {
				dispatch(success("Collection successfully created", collectionName));
			} else {
				dispatch(failure(res.data.message));
			}
		})
		.catch((err) => {
			console.log(err);
		});

	function request() {
		return { type: postConstants.CREATE_COLLECTION_REQUEST };
	}
	function success(message, collectionName) {
		return { type: postConstants.CREATE_COLLECTION_SUCCESS, successMessage: message, collectionName };
	}
	function failure(message) {
		return { type: postConstants.CREATE_POST_FAILURE, errorMessage: message };
	}
}

function deletePostFromCollection(postId, dispatch) {
	dispatch(request());

	Axios.delete(`/api/users/collections/posts/${postId}`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);
			if (res.status === 200) {
				dispatch(success("Post successfully deleted from favourites", postId));
			} else {
				dispatch(failure(res.data.message));
			}
		})
		.catch((err) => {
			console.log(err);
		});

	function request() {
		return { type: postConstants.DELETE_POST_FROM_COLLECTION_REQUEST };
	}
	function success(message, postId) {
		return { type: postConstants.DELETE_POST_FROM_COLLECTION_SUCCESS, successMessage: message, postId };
	}
	function failure(message) {
		return { type: postConstants.DELETE_POST_FROM_COLLECTION_FAILURE, errorMessage: message };
	}
}

function fetchFormData(postDTO) {
	let formData = new FormData();

	if (postDTO.postMedia !== "") {
		for (let i = 0; i < postDTO.postMedia.length; i++) {
			formData.append(`images[${i}]`, postDTO.postMedia[i]);
		}
	} else {
		formData.append("images", null);
	}
	formData.append("description", postDTO.description);
	formData.append("location", postDTO.location);
	formData.append("tags", JSON.stringify(postDTO.tags));
	return formData;
}

function likePost(postId, loggedUser, dispatch) {
	dispatch(request());

	Axios.put(`/api/posts/${postId}/like`, {}, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);
			if (res.status === 200) {
				dispatch(success(postId, loggedUser));
			} else {
				dispatch(failure("Error"));
			}
		})
		.catch((err) => {
			console.log(err);
			dispatch(failure("Error"));
		});

	function request() {
		return { type: postConstants.LIKE_POST_REQUEST };
	}
	function success(postId, loggedUser) {
		return { type: postConstants.LIKE_POST_SUCCESS, postId, loggedUser };
	}
	function failure(message) {
		return { type: postConstants.LIKE_POST_FAILURE, errorMessage: message };
	}
}

function unlikePost(postId, loggedUser, dispatch) {
	dispatch(request());
	Axios.put(`/api/posts/${postId}/unlike`, {}, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);
			if (res.status === 200) {
				dispatch(success(postId, loggedUser));
			} else {
				dispatch(failure("Error"));
			}
		})
		.catch((err) => {
			console.log(err);
			dispatch(failure("Error"));
		});

	function request() {
		return { type: postConstants.UNLIKE_POST_REQUEST };
	}
	function success(postId, loggedUser) {
		return { type: postConstants.UNLIKE_POST_SUCCESS, postId, loggedUser };
	}
	function failure(message) {
		return { type: postConstants.UNLIKE_POST_FAILURE, errorMessage: message };
	}
}

function dislikePost(postId, dispatch) {
	dispatch(request());

	Axios.put(`/api/posts/${postId}/dislike`, {}, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);
			if (res.status === 200) {
				dispatch(success(postId));
			} else {
				dispatch(failure("Error"));
			}
		})
		.catch((err) => {
			console.log(err);
			dispatch(failure("Error"));
		});

	function request() {
		return { type: postConstants.DISLIKE_POST_REQUEST };
	}
	function success(postId) {
		return { type: postConstants.DISLIKE_POST_SUCCESS, postId };
	}
	function failure(message) {
		return { type: postConstants.DISLIKE_POST_FAILURE, errorMessage: message };
	}
}

function undislikePost(postId, dispatch) {
	dispatch(request());

	Axios.put(`/api/posts/${postId}/undislike`, {}, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);
			if (res.status === 200) {
				dispatch(success(postId));
			} else {
				dispatch(failure("Error"));
			}
		})
		.catch((err) => {
			console.log(err);
			dispatch(failure("Error"));
		});

	function request() {
		return { type: postConstants.UNDISLIKE_POST_REQUEST };
	}
	function success(postId) {
		return { type: postConstants.UNDISLIKE_POST_SUCCESS, postId };
	}
	function failure(message) {
		return { type: postConstants.UNDISLIKE_POST_FAILURE, errorMessage: message };
	}
}

function commentPost(commentDTO, dispatch) {
	dispatch(request());

	Axios.put(`/api/posts/comments`, commentDTO, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);
			if (res.status === 200) {
				dispatch(success(res.data, commentDTO.PostId));
			} else {
				dispatch(failure("Error"));
			}
		})
		.catch((err) => {
			console.log(err);
			dispatch(failure("Error"));
		});

	function request() {
		return { type: postConstants.COMMENT_POST_REQUEST };
	}
	function success(comment, postId) {
		return { type: postConstants.COMMENT_POST_SUCCESS, comment, postId };
	}
	function failure(message) {
		return { type: postConstants.COMMENT_POST_FAILURE, errorMessage: message };
	}
}

function findPostsForGuestByHashtag(hashtag,dispatch){
	dispatch(request());
	Axios.get(`/api/posts/hashtag/${hashtag}/guest`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				failure();
			}
		})
		.catch((err) => {
			failure();
		});

	function request() {
		return { type: postConstants.GUEST_TIMELINE_POSTS_REQUEST };
	}

	function success(data) {
		return { type: postConstants.GUEST_TIMELINE_POSTS_SUCCESS, posts: data };
	}
	function failure() {
		return { type: postConstants.GUEST_TIMELINE_POSTS_FAILURE };
	}
}


function findPostsForUserByHashtag(hashtag,dispatch){
	dispatch(request());
	Axios.get(`/api/posts/hashtag/${hashtag}/user`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data));
				window.location = "#/search";
			} else {
				failure();
			}
		})
		.catch((err) => {
			failure();
		});

	
	function request() {
		return { type: postConstants.TIMELINE_POSTS_REQUEST };
	}
	
	function success(data) {
		return { type: postConstants.TIMELINE_POSTS_SUCCESS, posts: data };
	}
	
	function failure() {
		return { type: postConstants.TIMELINE_POSTS_FAILURE };
	}
}

function findPostsForGuestByLocation(location,dispatch){
	dispatch(request());
	Axios.get(`/api/posts/location/${location}/guest`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				failure();
			}
		})
		.catch((err) => {
			failure();
		});

	function request() {
		return { type: postConstants.GUEST_TIMELINE_POSTS_REQUEST };
	}

	function success(data) {
		return { type: postConstants.GUEST_TIMELINE_POSTS_SUCCESS, posts: data };
	}
	function failure() {
		return { type: postConstants.GUEST_TIMELINE_POSTS_FAILURE };
	}
}

function findPostsForUserByLocation(location,dispatch){
	dispatch(request());
	Axios.get(`/api/posts/location/${location}/user`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data));
				window.location = "#/search";
			} else {
				failure();
			}
		})
		.catch((err) => {
			failure();
		});

	function request() {
		return { type: postConstants.TIMELINE_POSTS_REQUEST };
	}
	
	function success(data) {
		return { type: postConstants.TIMELINE_POSTS_SUCCESS, posts: data };
	}
		
	function failure() {
		return { type: postConstants.TIMELINE_POSTS_FAILURE };
	}
}