import Axios from "axios";
import { modalConstants } from "../constants/ModalConstants";
import { storyConstants } from "../constants/StoryConstants";
import { authHeader } from "../helpers/auth-header";

export const storyService = {
	createStory,
	findStoriesForStoryline,
	findAllUserStories,
	GetStoriesForUser,
	visitedByUser,
};

function createStory(storyDTO, dispatch) {
	const formData = fetchFormData(storyDTO);
	dispatch(request());
	console.log(formData);

	Axios.post(`/api/story`, formData, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);

			if (res.status === 201) {
				dispatch(success("Story successfully created"));
			} else {
				dispatch(failure(res.data.message));
			}
		})
		.catch((err) => {
			console.log(err);
		});

	function request() {
		return { type: storyConstants.CREATE_STORY_REQUEST };
	}
	function success(message) {
		return { type: storyConstants.CREATE_STORY_SUCCESS, successMessage: message };
	}
	function failure(message) {
		return { type: storyConstants.CREATE_STORY_FAILURE, errorMessage: message };
	}
}

async function findStoriesForStoryline(dispatch) {
	dispatch(request());
	await Axios.get(`/api/story`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
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
		return { type: storyConstants.STORYLINE_STORY_REQUEST };
	}

	function success(data) {
		return { type: storyConstants.STORYLINE_STORY_SUCCESS, stories: data };
	}
	function failure() {
		return { type: storyConstants.STORYLINE_STORY_FAILURE };
	}
}

async function findAllUserStories(dispatch) {
	dispatch(request());
	await Axios.get(`/api/story/user`, { validateStatus: () => true, headers: authHeader() })
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
		return { type: storyConstants.USER_HIGHLIGHTS_STORY_REQUEST };
	}

	function success(data) {
		return { type: storyConstants.USER_HIGHLIGHTS_STORY_SUCCESS, stories: data };
	}
	function failure() {
		return { type: storyConstants.USER_HIGHLIGHTS_STORY_FAILURE };
	}
}

function fetchFormData(storyDTO) {
	let formData = new FormData();

	if (storyDTO.media !== "") {
		formData.append(`images`, storyDTO.media);
	} else {
		formData.append("images", null);
	}
	return formData;
}

function GetStoriesForUser(userId, dispatch) {
	Axios.get(`/api/story/` + userId, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				//failure()
			}
		})
		.catch((err) => {
			//failure()
		});

	function success(data) {
		return { type: modalConstants.SHOW_STORY_SLIDER_MODAL, stories: data };
	}
	//function failure() {
	//	return { type: storyConstants.STORYLINE_STORY_FAILURE };
	//}
}

function visitedByUser(storyId, dispatch) {
	Axios.put(`/api/story/${storyId}/visited`, {}, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			if (res.status === 200) {
				//dispatch(success(res.data));
			} else {
				//failure()
			}
		})
		.catch((err) => {
			//failure()
		});

	function success(data) {
		return { type: modalConstants.SHOW_STORY_SLIDER_MODAL, stories: data };
	}
	//function failure() {
	//	return { type: storyConstants.STORYLINE_STORY_FAILURE };
	//}
}
