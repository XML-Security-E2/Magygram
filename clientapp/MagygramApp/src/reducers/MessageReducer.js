import { messageConstants } from "../constants/MessageConstants";
import { modalConstants } from "../constants/ModalConstants";

var stateCpy = {};

export const messageReducer = (state, action) => {
	switch (action.type) {
		case modalConstants.SHOW_MESSAGE_SELECT_USER_MODAL:
			stateCpy = { ...state };
			stateCpy.selectUserModal.showModal = true;

			return stateCpy;
		case modalConstants.HIDE_MESSAGE_SELECT_USER_MODAL:
			stateCpy = { ...state };
			stateCpy.selectUserModal.showModal = false;

			return stateCpy;

		case messageConstants.SET_USER_MESSAGES_REQUEST:
			return state;

		case messageConstants.SET_USER_MESSAGES_SUCCESS:
			stateCpy = { ...state };
			stateCpy.selectUserModal.showModal = false;
			stateCpy.selectUserModal.selectedUser = action.messages.userInfo;
			stateCpy.showedMessages = action.messages.messages;
			console.log(stateCpy);
			return stateCpy;

		case messageConstants.SET_USER_MESSAGES_FAILURE:
			return state;

		case messageConstants.SET_USERS_CONVERSATIONS_REQUEST:
			return state;

		case messageConstants.SET_USERS_CONVERSATIONS_SUCCESS:
			stateCpy = { ...state };

			stateCpy.conversations = action.conversations !== null ? action.conversations : [];
			return stateCpy;

		case messageConstants.SET_USERS_CONVERSATIONS_FAILURE:
			return state;

		case messageConstants.SEND_MESSAGE_REQUEST:
			return state;

		case messageConstants.SEND_MESSAGE_SUCCESS:
			stateCpy = { ...state };
			console.log(stateCpy);

			if (!action.response.isMessageRequest) {
				if (stateCpy.conversations.find((conversation) => conversation.id === action.response.conversation.id) === undefined) {
					stateCpy.conversations.unshift(action.response.conversation);
				} else {
					let convCpy = stateCpy.conversations.find((conversation) => conversation.id === action.response.conversation.id);
					convCpy.lastMessage = action.response.conversation.lastMessage;
					convCpy.lastMessageUserId = action.response.conversation.lastMessageUserId;
				}
			} else {
				if (stateCpy.messageRequests.find((request) => request.id === action.response.messageRequest.id) === undefined) {
					stateCpy.messageRequests.unshift(action.response.messageRequest);
				}
			}

			stateCpy.showedMessages.push(action.response.conversation.lastMessage);

			console.log(stateCpy);
			return stateCpy;

		case messageConstants.SEND_MESSAGE_FAILURE:
			return state;

		case messageConstants.VIEW_MESSAGES_REQUEST:
			return state;

		case messageConstants.VIEW_MESSAGES_SUCCESS:
			stateCpy = { ...state };
			stateCpy.selectedConversationId = action.conversationId;

			let convCpy = stateCpy.conversations.find((conversation) => conversation.id === action.conversationId);
			convCpy.lastMessage.viewed = true;
			return stateCpy;

		case messageConstants.VIEW_MESSAGES_FAILURE:
			return state;

		case messageConstants.VIEW_MEDIA_MESSAGE_REQUEST:
			return state;

		case messageConstants.VIEW_MEDIA_MESSAGE_SUCCESS:
			stateCpy = { ...state };

			let convvCpy = stateCpy.showedMessages.find((message) => message.id === action.messageId);

			convvCpy.viewedMedia = true;

			return stateCpy;

		case messageConstants.VIEW_MEDIA_MESSAGE_FAILURE:
			return state;
		default:
			return state;
	}
};
