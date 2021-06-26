import React, { createContext, useReducer } from "react";
import { messageReducer } from "../reducers/MessageReducer";

export const MessageContext = createContext();

const MessageContextProvider = (props) => {
	const [messageState, dispatch] = useReducer(messageReducer, {
		selectUserModal: {
			showModal: false,
			selectedUser: {
				Id: "",
				Username: "",
				ImageURL: "",
			},
		},
		sendPostModal: {
			showModal: false,
			selectedUser: {
				Id: "",
				Username: "",
				ImageURL: "",
			},
		},
		sendStoryModal: {
			showModal: false,
		},
		storyModal: {
			showModal: false,
			stories: "",
		},
		conversationWithId: "",
		showedConversations: true,
		loadedConversationRequests: false,
		conversations: [],
		conversationRequests: [],
		showedMessages: [],
		selectedConversationId: "",
	});

	return <MessageContext.Provider value={{ messageState, dispatch }}>{props.children}</MessageContext.Provider>;
};

export default MessageContextProvider;
