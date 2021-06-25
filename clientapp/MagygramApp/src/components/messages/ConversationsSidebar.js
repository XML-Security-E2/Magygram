import React, { useContext } from "react";
import { MessageContext } from "../../contexts/MessageContext";
import { getDateTime } from "../../helpers/datetime-helper";
import { messageService } from "../../services/MessageService";

const ConversationsSidebar = () => {
	const { messageState, dispatch } = useContext(MessageContext);
	const imgStyle = { transform: "scale(1.5)", width: "100%", position: "absolute", left: "0" };

	const handleLoadConversation = async (userId, conversationId) => {
		await messageService.getMessagesFromUser(userId, dispatch).then(handleViewMessages(conversationId));
	};

	const handleViewMessages = (conversationId) => {
		return new Promise(function () {
			messageService.viewMessages(conversationId, dispatch);
		});
	};

	return (
		<React.Fragment>
			{messageState.conversations !== null &&
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
		</React.Fragment>
	);
};

export default ConversationsSidebar;
