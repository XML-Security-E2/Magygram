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
		conversations: [],
		messageRequests: [],
		showedMessages: [],
		selectedConversationId: "",
	});

	return <MessageContext.Provider value={{ messageState, dispatch }}>{props.children}</MessageContext.Provider>;
};

export default MessageContextProvider;
