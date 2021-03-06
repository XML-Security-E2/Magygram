import Axios from "axios";
import { messageConstants } from "../constants/MessageConstants";
import { authHeader } from "../helpers/auth-header";

export const messageService = {
	getUserConversations,
	getMessagesFromUser,
	sendMessage,
	viewMessages,
	viewMediaMessages,
	findPostById,
	findStoryById,
	getUserConversationRequests,
	getUserConversationRequestMessages,
	acceptRequest,
	denyRequest,
	deleteRequest,
};

function acceptRequest(requestId, dispatch) {
	dispatch(request());

	Axios.put(`/api/conversations/request/${requestId}/accept`, null, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(requestId));
			} else {
				dispatch(failure("Sorry, we have some internal problem."));
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: messageConstants.ACCEPT_MESSAGE_REQUEST_REQUEST };
	}
	function success(requestId) {
		return { type: messageConstants.ACCEPT_MESSAGE_REQUEST_SUCCESS, requestId };
	}
	function failure(error) {
		return { type: messageConstants.ACCEPT_MESSAGE_REQUEST_FAILURE, errorMessage: error };
	}
}

function denyRequest(requestId, dispatch) {
	dispatch(request());

	Axios.put(`/api/conversations/request/${requestId}/deny`, null, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(requestId));
			} else {
				dispatch(failure("Sorry, we have some internal problem."));
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: messageConstants.DENY_MESSAGE_REQUEST_REQUEST };
	}
	function success(requestId) {
		return { type: messageConstants.DENY_MESSAGE_REQUEST_SUCCESS, requestId };
	}
	function failure(error) {
		return { type: messageConstants.DENY_MESSAGE_REQUEST_FAILURE, errorMessage: error };
	}
}

function deleteRequest(requestId, dispatch) {
	dispatch(request());

	Axios.delete(`/api/conversations/request/${requestId}`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(requestId));
			} else {
				dispatch(failure("Sorry, we have some internal problem."));
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: messageConstants.DELETE_MESSAGE_REQUEST_REQUEST };
	}
	function success(requestId) {
		return { type: messageConstants.DELETE_MESSAGE_REQUEST_SUCCESS, requestId };
	}
	function failure(error) {
		return { type: messageConstants.DELETE_MESSAGE_REQUEST_FAILURE, errorMessage: error };
	}
}

async function findPostById(postId, dispatch) {
	dispatch(request());
	await Axios.get(`/api/posts/messages/id/${postId}`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else if (res.status === 401) {
				dispatch(unauthorized(res.data, postId));
			} else {
				dispatch(failure("Error while loading collections"));
			}
		})
		.catch((err) => {
			dispatch(failure());
		});

	function request() {
		return { type: messageConstants.SET_POST_MESSAGE_REQUEST };
	}

	function success(data) {
		return { type: messageConstants.SET_POST_MESSAGE_SUCCESS, post: data };
	}
	function failure(message) {
		return { type: messageConstants.SET_POST_MESSAGE_FAILURE, errorMessage: message };
	}

	function unauthorized(userInfo, postId) {
		return { type: messageConstants.SET_POST_MESSAGE_UNAUTHORIZED, userInfo, postId };
	}
}

async function findStoryById(storyId, dispatch) {
	dispatch(request());

	await Axios.get(`/api/story/id/${storyId}`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else if (res.status === 401) {
				dispatch(unauthorized(res.data, storyId));
			} else if (res.status === 403) {
				dispatch(expired(res.data, storyId));
			} else {
				dispatch(failure("Error while loading collections"));
			}
		})
		.catch((err) => {
			dispatch(failure());
		});

	function request() {
		return { type: messageConstants.SET_STORY_MESSAGE_REQUEST };
	}

	function success(data) {
		return { type: messageConstants.SET_STORY_MESSAGE_SUCCESS, story: data };
	}
	function failure(message) {
		return { type: messageConstants.SET_STORY_MESSAGE_FAILURE, errorMessage: message };
	}
	function unauthorized(userInfo, storyId) {
		return { type: messageConstants.SET_STORY_MESSAGE_UNAUTHORIZED, userInfo, storyId };
	}
	function expired(userInfo, storyId) {
		return { type: messageConstants.SET_STORY_MESSAGE_EXPIRED, userInfo, storyId };
	}
}

function viewMessages(conversationId, dispatch) {
	dispatch(request());

	Axios.put(`/api/messages/${conversationId}/view`, null, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(conversationId));
			} else {
				dispatch(failure("Sorry, we have some internal problem."));
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: messageConstants.VIEW_MESSAGES_REQUEST };
	}
	function success(conversationId) {
		return { type: messageConstants.VIEW_MESSAGES_SUCCESS, conversationId };
	}
	function failure(error) {
		return { type: messageConstants.VIEW_MESSAGES_FAILURE, errorMessage: error };
	}
}

function viewMediaMessages(conversationId, messageId, dispatch) {
	dispatch(request());

	Axios.put(`/api/messages/${conversationId}/${messageId}/view`, null, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(messageId));
			} else {
				dispatch(failure("Sorry, we have some internal problem."));
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: messageConstants.VIEW_MEDIA_MESSAGE_REQUEST };
	}
	function success(messageId) {
		return { type: messageConstants.VIEW_MEDIA_MESSAGE_SUCCESS, messageId };
	}
	function failure(error) {
		return { type: messageConstants.VIEW_MEDIA_MESSAGE_FAILURE, errorMessage: error };
	}
}

function sendMessage(messageDTO, dispatch) {
	const formData = fetchFormData(messageDTO);

	dispatch(request());

	Axios.post(`/api/messages`, formData, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 201) {
				dispatch(success(res.data));
			} else {
				dispatch(failure("Sorry, we have some internal problem."));
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: messageConstants.SEND_MESSAGE_REQUEST };
	}
	function success(response) {
		return { type: messageConstants.SEND_MESSAGE_SUCCESS, response };
	}
	function failure(error) {
		return { type: messageConstants.SEND_MESSAGE_FAILURE, errorMessage: error };
	}
}

function fetchFormData(messageDTO) {
	let formData = new FormData();

	if (messageDTO.media !== "") {
		formData.append("media", messageDTO.media);
	} else {
		formData.append("media", null);
	}

	formData.append("messageTo", messageDTO.messageTo);
	formData.append("messageType", messageDTO.messageType);
	formData.append("text", messageDTO.text);
	formData.append("contentId", messageDTO.contentId);

	return formData;
}

async function getUserConversations(dispatch) {
	dispatch(request());

	await Axios.get(`/api/conversations`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				dispatch(failure("Sorry, we have some internal problem."));
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: messageConstants.SET_USERS_CONVERSATIONS_REQUEST };
	}
	function success(conversations) {
		return { type: messageConstants.SET_USERS_CONVERSATIONS_SUCCESS, conversations };
	}
	function failure(error) {
		return { type: messageConstants.SET_USERS_CONVERSATIONS_FAILURE, errorMessage: error };
	}
}

async function getUserConversationRequests(dispatch) {
	dispatch(request());

	await Axios.get(`/api/messages/requests`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				dispatch(failure("Sorry, we have some internal problem."));
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: messageConstants.SET_USER_REQUESTS_REQUEST };
	}
	function success(conversationRequests) {
		return { type: messageConstants.SET_USER_REQUESTS_SUCCESS, conversationRequests };
	}
	function failure(error) {
		return { type: messageConstants.SET_USER_REQUESTS_FAILURE, errorMessage: error };
	}
}

async function getMessagesFromUser(userId, dispatch) {
	dispatch(request());

	await Axios.get(`/api/messages/${userId}`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data, userId));
			} else {
				dispatch(failure("Sorry, we have some internal problem."));
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: messageConstants.SET_USER_MESSAGES_REQUEST };
	}
	function success(messages, userId) {
		return { type: messageConstants.SET_USER_MESSAGES_SUCCESS, messages, userId };
	}
	function failure(error) {
		return { type: messageConstants.SET_USER_MESSAGES_FAILURE, errorMessage: error };
	}
}

async function getUserConversationRequestMessages(userId, dispatch) {
	dispatch(request());

	await Axios.get(`/api/messages/${userId}/requests`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data, userId));
			} else {
				dispatch(failure("Sorry, we have some internal problem."));
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: messageConstants.SET_USER_REQUEST_MESSAGES_REQUEST };
	}
	function success(messages, userId) {
		return { type: messageConstants.SET_USER_REQUEST_MESSAGES_SUCCESS, messages, userId };
	}
	function failure(error) {
		return { type: messageConstants.SET_USER_REQUEST_MESSAGES_FAILURE, errorMessage: error };
	}
}
