import Axios from "axios";
import { modalConstants } from "../constants/ModalConstants";
import { storyConstants } from "../constants/StoryConstants";
import { authHeader } from "../helpers/auth-header";

export const storyService = {
	createStory,
	createAgentStory,
	createHighlight,
	findAllProfileHighlights,
	findAllStoriesByHighlightName,
	findStoriesForStoryline,
	findAllUserStories,
	GetStoriesForUser,
	visitedByUser,
	HaveActiveStoriesLoggedUser,
	findStoryById,
	reportStory,
};

function reportStory(reportDTO, dispatch) {
	dispatch(request());

	Axios.post(`/api/report`, reportDTO, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success("Report sent successfully"));
			} else {
				dispatch(failure(res.data.message));
			}
		})
		.catch((err) => {
			console.log(err);
			dispatch(failure("Internal server error"));
		});

	function request() {
		return { type: storyConstants.REPORT_STORY_REQUEST };
	}
	function success(message) {
		return { type: storyConstants.REPORT_STORY_SUCCESS, successMessage: message };
	}
	function failure(message) {
		return { type: storyConstants.REPORT_STORY_FAILURE, errorMessage: message };
	}
}

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

function createAgentStory(storyCampaignDTO, dispatch) {
	const formData = fetchFormData(storyCampaignDTO);
	dispatch(request());
	console.log(formData);

	formData.append("minAge", storyCampaignDTO.minAge);
	formData.append("maxAge", storyCampaignDTO.maxAge);
	formData.append("displayTime", storyCampaignDTO.displayTime);
	formData.append("frequency", storyCampaignDTO.frequency);
	formData.append("gender", storyCampaignDTO.gender);
	formData.append("startDate", storyCampaignDTO.startDate);
	formData.append("endDate", storyCampaignDTO.endDate);

	Axios.post(`/api/story/campaign`, formData, { validateStatus: () => true, headers: authHeader() })
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
		return { type: storyConstants.CREATE_AGENT_STORY_REQUEST };
	}
	function success(message) {
		return { type: storyConstants.CREATE_AGENT_STORY_SUCCESS, successMessage: message };
	}
	function failure(message) {
		return { type: storyConstants.CREATE_AGENT_STORY_FAILURE, errorMessage: message };
	}
}

function createHighlight(highlightDTO, dispatch) {
	dispatch(request());

	if (validateHighlights(highlightDTO, dispatch)) {
		Axios.post(`/api/users/highlights`, highlightDTO, { validateStatus: () => true, headers: authHeader() })
			.then((res) => {
				console.log(res);
				if (res.status === 201) {
					findAllProfileHighlights(localStorage.getItem("userId"), dispatch);
					dispatch(success(res.data));
				} else {
					dispatch(failure(res.data.message));
				}
			})
			.catch((err) => {
				console.log(err);
			});
	}

	function request() {
		return { type: storyConstants.CREATE_HIGHLIGHTS_STORY_REQUEST };
	}
	function success(data) {
		return { type: storyConstants.CREATE_HIGHLIGHTS_STORY_SUCCESS, highlight: data };
	}
	function failure(message) {
		return { type: storyConstants.CREATE_HIGHLIGHTS_STORY_FAILURE, errorMessage: message };
	}
}

function validateHighlights(highlightDTO, dispatch) {
	if (highlightDTO.storyIds.length === 0) {
		dispatch(failure("You must select story"));
		return false;
	} else {
		return true;
	}

	function failure(message) {
		return { type: storyConstants.CREATE_HIGHLIGHTS_STORY_FAILURE, errorMessage: message };
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

async function findAllProfileHighlights(userId, dispatch) {
	dispatch(request());
	await Axios.get(`/api/users/${userId}/highlights`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);
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
		return { type: storyConstants.PROFILE_HIGHLIGHTS_REQUEST };
	}

	function success(data) {
		return { type: storyConstants.PROFILE_HIGHLIGHTS_SUCCESS, highlights: data };
	}
	function failure() {
		return { type: storyConstants.PROFILE_HIGHLIGHTS_FAILURE };
	}
}

function findAllStoriesByHighlightName(userId, name, dispatch) {
	dispatch(request());

	Axios.get(`/api/users/${userId}/highlights/${name}`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data, name));
			} else {
				dispatch(failure());
			}
		})
		.catch((err) => {
			dispatch(failure());
		});

	function request() {
		return { type: storyConstants.FIND_HIGHLIGHT_BY_NAME_REQUEST };
	}

	function success(data, name) {
		return { type: storyConstants.FIND_HIGHLIGHT_BY_NAME_SUCCESS, highlights: data, name };
	}
	function failure() {
		return { type: storyConstants.FIND_HIGHLIGHT_BY_NAME_FAILURE };
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
	formData.append("tags", JSON.stringify(storyDTO.tags));
	return formData;
}

function findStoryById(storyId, userId, dispatch) {
	Axios.get(`/api/story/${storyId}/getForAdmin`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data, 1));
			} else {
				//failure()
			}
		})
		.catch((err) => {
			//failure()
		});

	function success(data, visited) {
		return { type: modalConstants.SHOW_STORY_SLIDER_MODAL_ADMIN, stories: data, visited, userId };
	}
}

function GetStoriesForUser(userId, visited, dispatch) {
	Axios.get(`/api/story/` + userId, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			if (res.status === 200) {
				dispatch(success(res.data, visited));
			} else {
				//failure()
			}
		})
		.catch((err) => {
			//failure()
		});

	function success(data, visited) {
		return { type: modalConstants.SHOW_STORY_SLIDER_MODAL, stories: data, visited, userId };
	}
	//function failure() {
	//	return { type: storyConstants.STORYLINE_STORY_FAILURE };
	//}
}

function visitedByUser(storyId, dispatch) {
	Axios.put(`/api/story/${storyId}/visited`, {}, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			if (res.status === 200) {
				dispatch(success(storyId));
			} else {
				//failure()
			}
		})
		.catch((err) => {
			//failure()
		});

	function success(storyId, lastStory) {
		return { type: storyConstants.VISITED_STORY_SUCCESS, storyId };
	}
	//function failure() {
	//	return { type: storyConstants.STORYLINE_STORY_FAILURE };
	//}
}

function HaveActiveStoriesLoggedUser(dispatch) {
	dispatch(request());
	Axios.get(`/api/story/activestories`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				dispatch(failure());
			}
		})
		.catch((err) => {
			dispatch(failure());
		});

	function request() {
		return { type: storyConstants.HAVE_LOGGED_USER_STORY_REQUEST };
	}

	function success(data) {
		return { type: storyConstants.HAVE_LOGGED_USER_STORY_SUCCESS, haveStories: data };
	}

	function failure() {
		return { type: storyConstants.HAVE_LOGGED_USER_STORY_FAILURE };
	}
}
