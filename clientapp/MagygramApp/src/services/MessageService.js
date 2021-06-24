import Axios from "axios";
import { messageConstants } from "../constants/MessageConstants";
import { authHeader } from "../helpers/auth-header";

export const messageService = {
	getUserConversations,
	getMessagesFromUser,
	sendMessage,
	viewMessages,
};

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

function sendMessage(messageDTO, dispatch) {
	dispatch(request());

	Axios.post(`/api/messages`, messageDTO, { validateStatus: () => true, headers: authHeader() })
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
	function success(conversation) {
		return { type: messageConstants.SEND_MESSAGE_SUCCESS, conversation };
	}
	function failure(error) {
		return { type: messageConstants.SEND_MESSAGE_FAILURE, errorMessage: error };
	}
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

async function getMessagesFromUser(userId, dispatch) {
	dispatch(request());

	await Axios.get(`/api/messages/${userId}`, { validateStatus: () => true, headers: authHeader() })
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
		return { type: messageConstants.SET_USER_MESSAGES_REQUEST };
	}
	function success(messages) {
		return { type: messageConstants.SET_USER_MESSAGES_SUCCESS, messages };
	}
	function failure(error) {
		return { type: messageConstants.SET_USER_MESSAGES_FAILURE, errorMessage: error };
	}
}
