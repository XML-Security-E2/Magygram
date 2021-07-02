import { messageConstants } from "../constants/MessageConstants";
import { modalConstants } from "../constants/ModalConstants";
import { getUserInfo } from "../helpers/auth-header";

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

		case modalConstants.SHOW_SEND_POST_TO_USER_MODAL:
			stateCpy = { ...state };
			stateCpy.sendPostModal.postId = action.postId;
			stateCpy.sendPostModal.showModal = true;

			return stateCpy;
		case modalConstants.HIDE_SEND_POST_TO_USER_MODAL:
			stateCpy = { ...state };
			stateCpy.sendPostModal.showModal = false;

			return stateCpy;

		case modalConstants.SHOW_SEND_STORY_TO_USER_MODAL:
			stateCpy = { ...state };
			stateCpy.sendStoryModal.storyId = action.storyId;
			stateCpy.sendStoryModal.showModal = true;

			return stateCpy;
		case modalConstants.HIDE_SEND_STORY_TO_USER_MODAL:
			stateCpy = { ...state };
			stateCpy.sendStoryModal.showModal = false;

			return stateCpy;

		case modalConstants.SHOW_STORY_MESSAGE_MODAL:
			stateCpy = { ...state };

			stateCpy.storyModal.stories = createStory(action.story);

			stateCpy.storyModal.showModal = true;

			return stateCpy;
		case modalConstants.HIDE_STORY_MESSAGE_MODAL:
			stateCpy = { ...state };
			stateCpy.storyModal.showModal = false;

			return stateCpy;

		case messageConstants.SET_USER_MESSAGES_REQUEST:
			return state;

		case messageConstants.SET_USER_MESSAGES_SUCCESS:
			stateCpy = { ...state };
			stateCpy.conversationWithId = action.userId;
			stateCpy.selectUserModal.showModal = false;
			stateCpy.selectUserModal.selectedUser = action.messages.userInfo;
			stateCpy.showedMessages = action.messages.messages;
			stateCpy.loadedConversationRequests = false;

			stateCpy.showedMessages.forEach((message) => {
				if (message.messageType === "POST") {
					message.post = {
						Id: "",
						MediaType: "",
						Url: "",
						Description: "",
						UserId: "",
						UserImageUrl: "",
						Username: "",
						Unauthorized: true,
					};
				} else if (message.messageType === "STORY") {
					message.story = {
						Id: "",
						MediaType: "",
						Url: "",
						UserId: "",
						UserImageUrl: "",
						Username: "",
						Unauthorized: true,
						Expired: true,
					};
				}
			});

			console.log(stateCpy);
			return stateCpy;

		case messageConstants.SET_USER_MESSAGES_FAILURE:
			return state;

		case messageConstants.SET_USER_REQUEST_MESSAGES_REQUEST:
			return state;

		case messageConstants.SET_USER_REQUEST_MESSAGES_SUCCESS:
			stateCpy = { ...state };
			stateCpy.conversationWithId = action.userId;
			stateCpy.selectUserModal.showModal = false;
			stateCpy.selectUserModal.selectedUser = action.messages.userInfo;
			stateCpy.showedMessages = action.messages.messages;
			stateCpy.loadedConversationRequests = true;

			stateCpy.showedMessages.forEach((message) => {
				if (message.messageType === "POST") {
					message.post = {
						Id: "",
						MediaType: "",
						Url: "",
						Description: "",
						UserId: "",
						UserImageUrl: "",
						Username: "",
						Unauthorized: true,
					};
				} else if (message.messageType === "STORY") {
					message.story = {
						Id: "",
						MediaType: "",
						Url: "",
						UserId: "",
						UserImageUrl: "",
						Username: "",
						Unauthorized: true,
						Expired: true,
					};
				}
			});

			console.log(stateCpy);
			return stateCpy;

		case messageConstants.SET_USER_REQUEST_MESSAGES_FAILURE:
			return state;

		case messageConstants.SET_POST_MESSAGE_REQUEST:
			return state;

		case messageConstants.SET_POST_MESSAGE_SUCCESS:
			stateCpy = { ...state };

			stateCpy.showedMessages.forEach((message) => {
				if (message.contentId === action.post.Id) {
					message.post = {
						Id: action.post.Id,
						MediaType: action.post.Media[0].MediaType,
						Url: action.post.Media[0].Url,
						Description: action.post.Description,
						UserId: action.post.UserInfo.Id,
						UserImageUrl: action.post.UserInfo.ImageURL,
						Username: action.post.UserInfo.Username,
						Unauthorized: false,
					};
				}
			});

			console.log(stateCpy);
			return stateCpy;

		case messageConstants.SET_POST_MESSAGE_UNAUTHORIZED:
			stateCpy = { ...state };

			stateCpy.showedMessages.forEach((message) => {
				if (message.contentId === action.postId) {
					message.post = {
						Id: "",
						MediaType: "",
						Url: "",
						Description: "",
						UserId: action.userInfo.Id,
						UserImageUrl: "",
						Username: action.userInfo.Username,
						Unauthorized: true,
					};
				}
			});

			console.log(stateCpy);
			return stateCpy;

		case messageConstants.SET_POST_MESSAGE_FAILURE:
			return state;

		case messageConstants.SET_STORY_MESSAGE_REQUEST:
			return state;

		case messageConstants.SET_STORY_MESSAGE_SUCCESS:
			stateCpy = { ...state };

			stateCpy.showedMessages.forEach((message) => {
				if (message.contentId === action.story.id) {
					message.story = {
						Id: action.story.Id,
						MediaType: action.story.media.MediaType,
						Url: action.story.media.Url,
						UserId: action.story.userInfo.Id,
						UserImageUrl: action.story.userInfo.ImageURL,
						Username: action.story.userInfo.Username,
						Tags: action.story.media.Tags,
						Unauthorized: false,
						Expired: false,
					};
				}
			});

			console.log(stateCpy);
			return stateCpy;

		case messageConstants.SET_STORY_MESSAGE_UNAUTHORIZED:
			stateCpy = { ...state };

			stateCpy.showedMessages.forEach((message) => {
				if (message.contentId === action.storyId) {
					message.story = {
						Id: "",
						MediaType: "",
						Url: "",
						UserId: action.userInfo.Id,
						UserImageUrl: "",
						Username: action.userInfo.Username,
						Unauthorized: true,
						Expired: false,
					};
				}
			});

			console.log(stateCpy);
			return stateCpy;

		case messageConstants.SET_STORY_MESSAGE_EXPIRED:
			stateCpy = { ...state };

			stateCpy.showedMessages.forEach((message) => {
				if (message.contentId === action.storyId) {
					message.story = {
						Id: "",
						MediaType: "",
						Url: "",
						UserId: action.userInfo.Id,
						UserImageUrl: "",
						Username: action.userInfo.Username,
						Unauthorized: false,
						Expired: true,
					};
				}
			});

			console.log(stateCpy);
			return stateCpy;

		case messageConstants.SET_STORY_MESSAGE_FAILURE:
			return state;

		case messageConstants.SET_USERS_CONVERSATIONS_REQUEST:
			return state;

		case messageConstants.SET_USERS_CONVERSATIONS_SUCCESS:
			stateCpy = { ...state };

			stateCpy.conversations = action.conversations !== null ? action.conversations : [];
			return stateCpy;

		case messageConstants.SET_USERS_CONVERSATIONS_FAILURE:
			return state;

		case messageConstants.SET_USER_REQUESTS_REQUEST:
			return state;

		case messageConstants.SET_USER_REQUESTS_SUCCESS:
			stateCpy = { ...state };

			stateCpy.conversationRequests = action.conversationRequests !== null ? action.conversationRequests : [];
			return stateCpy;

		case messageConstants.SET_USER_REQUESTS_FAILURE:
			return state;

		case messageConstants.SEND_MESSAGE_REQUEST:
			return state;

		case messageConstants.SEND_MESSAGE_SUCCESS:
			stateCpy = { ...state };

			if (!action.response.isMessageRequest) {
				if (stateCpy.conversations.find((conversation) => conversation.id === action.response.conversation.id) === undefined) {
					stateCpy.conversations.unshift(action.response.conversation);
				} else {
					let convCpy = stateCpy.conversations.find((conversation) => conversation.id === action.response.conversation.id);
					convCpy.lastMessage = action.response.conversation.lastMessage;
					convCpy.lastMessageUserId = action.response.conversation.lastMessageUserId;
				}
			} else {
				if (action.response.conversationRequest.lastMessage.messageFromId !== getUserInfo().Id) {
					if (stateCpy.conversationRequests.find((request) => request.id === action.response.conversationRequest.id) === undefined) {
						stateCpy.conversationRequests.unshift(action.response.conversationRequest);
					} else {
						let convCpy = stateCpy.conversationRequests.find((conversation) => conversation.id === action.response.conversationRequest.id);
						convCpy.lastMessage = action.response.conversationRequest.lastMessage;
						convCpy.lastMessageUserId = action.response.conversationRequest.lastMessageUserId;
					}
				}
			}
			stateCpy.sendStoryModal.showModal = false;
			stateCpy.sendPostModal.showModal = false;
			if (!action.response.isMessageRequest) {
				stateCpy.showedMessages.push(action.response.conversation.lastMessage);
			} else {
				stateCpy.showedMessages.push(action.response.conversationRequest.lastMessage);
			}

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

		case messageConstants.SET_SELECTED_REQUEST_CONVERSATION:
			stateCpy = { ...state };
			stateCpy.selectedConversationId = action.conversationId;

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

		case messageConstants.SHOW_CONVERSATIONS:
			stateCpy = { ...state };

			stateCpy.showedConversations = true;
			return stateCpy;

		case messageConstants.SHOW_CONVERSATION_REQUESTS:
			stateCpy = { ...state };

			stateCpy.showedConversations = false;
			return stateCpy;

		case messageConstants.ACCEPT_MESSAGE_REQUEST_REQUEST:
			return state;

		case messageConstants.ACCEPT_MESSAGE_REQUEST_SUCCESS:
			stateCpy = { ...state };

			stateCpy.selectUserModal = {
				showModal: false,
				selectedUser: {
					Id: "",
					Username: "",
					ImageURL: "",
				},
			};
			if (stateCpy.conversations.find((conversation) => conversation.id === action.requestId) === undefined) {
				let st = stateCpy.conversationRequests.find((request) => request.id === action.requestId);
				stateCpy.conversations.unshift(st);
			}
			stateCpy.showedMessages = [];
			stateCpy.conversationRequests = state.conversationRequests.filter((request) => request.id !== action.requestId);
			if (stateCpy.conversationRequests.length === 0) {
				stateCpy.loadedConversationRequests = false;
				stateCpy.showedConversations = true;
				stateCpy.conversationWithId = "";
			}
			console.log(stateCpy);
			return stateCpy;

		case messageConstants.ACCEPT_MESSAGE_REQUEST_FAILURE:
			return state;

		case messageConstants.DENY_MESSAGE_REQUEST_REQUEST:
			return state;

		case messageConstants.DENY_MESSAGE_REQUEST_SUCCESS:
			stateCpy = { ...state };

			stateCpy.selectUserModal = {
				showModal: false,
				selectedUser: {
					Id: "",
					Username: "",
					ImageURL: "",
				},
			};

			stateCpy.showedMessages = [];
			stateCpy.conversationRequests = state.conversationRequests.filter((request) => request.id !== action.requestId);
			if (stateCpy.conversationRequests.length === 0) {
				stateCpy.loadedConversationRequests = false;
				stateCpy.showedConversations = true;
				stateCpy.conversationWithId = "";
			}
			return stateCpy;

		case messageConstants.DENY_MESSAGE_REQUEST_FAILURE:
			return state;

		case messageConstants.DELETE_MESSAGE_REQUEST_REQUEST:
			return state;

		case messageConstants.DELETE_MESSAGE_REQUEST_SUCCESS:
			stateCpy = { ...state };

			stateCpy.selectUserModal = {
				showModal: false,
				selectedUser: {
					Id: "",
					Username: "",
					ImageURL: "",
				},
			};

			stateCpy.showedMessages = [];
			stateCpy.conversationRequests = state.conversationRequests.filter((request) => request.id !== action.requestId);
			if (stateCpy.conversationRequests.length === 0) {
				stateCpy.loadedConversationRequests = false;
				stateCpy.showedConversations = true;
				stateCpy.conversationWithId = "";
			}
			return stateCpy;

		case messageConstants.DELETE_MESSAGE_REQUEST_FAILURE:
			return state;

		case messageConstants.SET_MESSAGE_FOR_NOTIFICATION:
			stateCpy = { ...state };

			if (!action.response.isMessageRequest) {
				if (stateCpy.conversations.find((conversation) => conversation.id === action.response.conversation.id) === undefined) {
					stateCpy.conversations.unshift(action.response.conversation);
				} else {
					let convCpy = stateCpy.conversations.find((conversation) => conversation.id === action.response.conversation.id);
					convCpy.lastMessage = action.response.conversation.lastMessage;
					convCpy.lastMessageUserId = action.response.conversation.lastMessageUserId;
				}
			} else {
				if (action.response.conversationRequest.lastMessage.messageFromId !== getUserInfo().Id) {
					if (stateCpy.conversationRequests.find((request) => request.id === action.response.conversationRequest.id) === undefined) {
						stateCpy.conversationRequests.unshift(action.response.conversationRequest);
					} else {
						let convCpy = stateCpy.conversationRequests.find((conversation) => conversation.id === action.response.conversationRequest.id);
						convCpy.lastMessage = action.response.conversationRequest.lastMessage;
						convCpy.lastMessageUserId = action.response.conversationRequest.lastMessageUserId;
					}
				}
			}
			stateCpy.sendStoryModal.showModal = false;
			stateCpy.sendPostModal.showModal = false;

			if (!action.response.isMessageRequest) {
				if (state.conversationWithId === action.response.conversation.lastMessage.messageFromId) {
					stateCpy.showedMessages.push(action.response.conversation.lastMessage);
				}
			} else {
				if (state.conversationWithId === action.response.conversationRequest.lastMessage.messageFromId) {
					stateCpy.showedMessages.push(action.response.conversationRequest.lastMessage);
				}
			}

			console.log(stateCpy);
			return stateCpy;
		default:
			return state;
	}
};

function createStory(story) {
	var retVal = [];

	retVal.push({
		url: story.Url,
		header: {
			heading: story.Username,
			profileImage: story.UserImageUrl === "" ? "assets/img/profile.jpg" : story.UserImageUrl,
			storyId: story.Id,
		},
		type: story.MediaType === "VIDEO" ? "video" : "image",
		tags: story.Tags,
	});

	return retVal;
}
