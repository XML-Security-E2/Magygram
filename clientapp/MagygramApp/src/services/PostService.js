import Axios from "axios";
import { postConstants } from "../constants/PostConstants";
import { authHeader } from "../helpers/auth-header";


export const postService = {
	findPostsForTimeline,
	createPost,
	likePost,
	unlikePost,
	dislikePost,
	undislikePost,
};

async function findPostsForTimeline(dispatch) {
	dispatch(request());
	await Axios.get(`/api/posts`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				failure()
			}
		})
			.catch((err) => {
				failure()
			});

		function request() {
			return { type: postConstants.TIMELINE_POSTS_REQUEST};
		}

		function success(data) {
			return { type: postConstants.TIMELINE_POSTS_SUCCESS, posts: data };
		}
		function failure() {
			return { type: postConstants.TIMELINE_POSTS_FAILURE };
		}

};

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

function likePost(postId, dispatch) {
	dispatch(request());

	Axios.put(`/api/posts/${postId}/like`, {} ,{ validateStatus: () => true, headers: authHeader() })
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
		return { type: postConstants.LIKE_POST_REQUEST };
	}
	function success(postId) {
		return { type: postConstants.LIKE_POST_SUCCESS, postId };
	}
	function failure(message) {
		return { type: postConstants.LIKE_POST_FAILURE, errorMessage: message };
	}
}

function unlikePost(postId, dispatch) {
	dispatch(request());
	dispatch(success(postId));
	dispatch(failure("Test"));

	function request() {
		return { type: postConstants.UNLIKE_POST_REQUEST };
	}
	function success(postId) {
		return { type: postConstants.UNLIKE_POST_SUCCESS, postId };
	}
	function failure(message) {
		return { type: postConstants.UNLIKE_POST_FAILURE, errorMessage: message };
	}
}

function dislikePost(postId, dispatch) {
	dispatch(request());
	dispatch(success(postId));
	dispatch(failure("Test"));

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
	dispatch(success(postId));
	dispatch(failure("Test"));

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