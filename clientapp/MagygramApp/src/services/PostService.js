import Axios from "axios";
import { postConstants } from "../constants/PostConstants";
import { authHeader } from "../helpers/auth-header";

export const postService = {
	createPost,
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
