import React, { useContext } from "react";
import { messageConstants } from "../../constants/MessageConstants";
import { MessageContext } from "../../contexts/MessageContext";
import { getDateTime } from "../../helpers/datetime-helper";
import { messageService } from "../../services/MessageService";

const ConversationsSidebar = () => {
	const { messageState, dispatch } = useContext(MessageContext);
	const imgStyle = { transform: "scale(1.5)", width: "100%", position: "absolute", left: "0" };

	const handleLoadConversation = async (userId, conversationId) => {
		await messageService.getMessagesFromUser(userId, dispatch).then(handleViewMessages(conversationId));
	};

	const handleLoadConversationRequest = async (userId, conversationId) => {
		await messageService.getUserConversationRequestMessages(userId, dispatch).then(handleSetConversationId(conversationId));
	};

	const handleViewMessages = (conversationId) => {
		return new Promise(function () {
			messageService.viewMessages(conversationId, dispatch);
		});
	};

	const handleSetConversationId = (conversationId) => {
		return new Promise(function () {
			dispatch({ type: messageConstants.SET_SELECTED_REQUEST_CONVERSATION, conversationId });
		});
	};

	const showConversations = () => {
		dispatch({ type: messageConstants.SHOW_CONVERSATIONS });
	};

	const showConversationRequests = () => {
		dispatch({ type: messageConstants.SHOW_CONVERSATION_REQUESTS });
	};

	return (
		<React.Fragment>
			<div className="row" hidden={messageState.conversationRequests === null || messageState.conversationRequests.length == 0}>
				<div className="col-6">
					<button disabled={messageState.showedConversations} type="button" className="btn btn-link btn-fw text-primary w-100 border-0" onClick={showConversations}>
						Conversations
					</button>
				</div>
				<div className="col-6">
					<button disabled={!messageState.showedConversations} type="button" className="btn btn-link btn-fw text-secondary w-100 border-0" onClick={showConversationRequests}>
						Requests
					</button>
				</div>
			</div>
			{messageState.showedConversations &&
				messageState.conversations !== null &&
				messageState.conversations.map((conversation) => {
					return (
						<div key={conversation.id} className="d-flex flex-row align-items-center mt-2" onClick={() => handleLoadConversation(conversation.participant.Id, conversation.id)}>
							<div className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border sidenav-profile-photo">
								<img src={conversation.participant.ImageURL === "" ? "assets/img/profile.jpg" : conversation.participant.ImageURL} alt="" style={imgStyle} />
							</div>
							<div className="profile-info ml-3">
								<span className="profile-info-username">{conversation.participant.Username}</span>
								{conversation.lastMessage.messageType === "TEXT" && (
									<span>{conversation.lastMessage.viewed ? conversation.lastMessage.text.substring(0, 12) : <b>{conversation.lastMessage.text.substring(0, 12)}</b>}</span>
								)}

								{conversation.lastMessage.messageType === "MEDIA" && <span>{conversation.lastMessage.viewed ? "media" : <b>media</b>}</span>}

								<span className="text-secondary ml-2">
									{conversation.lastMessage.viewed ? getDateTime(conversation.lastMessage.timestamp) : <b>{getDateTime(conversation.lastMessage.timestamp)}</b>}
								</span>
							</div>
						</div>
					);
				})}

			{!messageState.showedConversations &&
				messageState.conversationRequests !== null &&
				messageState.conversationRequests.map((conversation) => {
					return (
						<div key={conversation.id} className="d-flex flex-row align-items-center mt-2" onClick={() => handleLoadConversationRequest(conversation.participant.Id, conversation.id)}>
							<div className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border sidenav-profile-photo">
								<img src={conversation.participant.ImageURL === "" ? "assets/img/profile.jpg" : conversation.participant.ImageURL} alt="" style={imgStyle} />
							</div>
							<div className="profile-info ml-3">
								<span className="profile-info-username">{conversation.participant.Username}</span>
								{conversation.lastMessage.messageType === "TEXT" && (
									<span>{conversation.lastMessage.viewed ? conversation.lastMessage.text.substring(0, 12) : <b>{conversation.lastMessage.text.substring(0, 12)}</b>}</span>
								)}

								{conversation.lastMessage.messageType === "MEDIA" && <span>{conversation.lastMessage.viewed ? "media" : <b>media</b>}</span>}

								<span className="text-secondary ml-2">
									{conversation.lastMessage.viewed ? getDateTime(conversation.lastMessage.timestamp) : <b>{getDateTime(conversation.lastMessage.timestamp)}</b>}
								</span>
							</div>
						</div>
					);
				})}
		</React.Fragment>
	);
};

export default ConversationsSidebar;
