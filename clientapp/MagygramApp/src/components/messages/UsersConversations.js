import React, { useContext, useEffect, useRef, useState } from "react";
import Chat from "./Chat";
import ConversationsSidebar from "./ConversationsSidebar";
import ChatForm from "./ChatForm";
import { getUserInfo } from "../../helpers/auth-header";
import { MessageContext } from "../../contexts/MessageContext";
import { modalConstants } from "../../constants/ModalConstants";
import UserChatHeader from "./UserChatHeader";
import { messageService } from "../../services/MessageService";
import FrontPageChat from "./FrontPageChat";
import { messageConstants } from "../../constants/MessageConstants";

const UsersConversations = () => {
	const { messageState, dispatch } = useContext(MessageContext);
	const [rrmessages, setRrmessages] = useState([]);

	const wss = useRef(null);

	const openUserSelectModal = () => {
		dispatch({ type: modalConstants.SHOW_MESSAGE_SELECT_USER_MODAL });
	};
	useEffect(() => {
		wss.current = new WebSocket(
			process.env.REACT_APP_STAGE === "prod" ? "ws://localhost:467/ws/messages/" + localStorage.getItem("userId") : "wss://localhost:467/ws/messages/" + localStorage.getItem("userId")
		);
		wss.current.onopen = () => console.log("ws je ovdeeeeee opened");
		wss.current.onclose = () => console.log("ws closed");

		const getUserConversations = async () => {
			await messageService.getUserConversations(dispatch);
		};
		const getUserConversationRequests = async () => {
			await messageService.getUserConversationRequests(dispatch);
		};
		getUserConversations();
		getUserConversationRequests();
	}, []);

	useEffect(() => {
		if (!wss.current) return;

		wss.current.onmessage = (evt) => {
			console.log("LALLALLALAL");
			console.log(evt.data);
			let a = [...rrmessages];
			a.push(evt.data);

			setRrmessages(a);
			dispatch({ type: messageConstants.SET_MESSAGE_FOR_NOTIFICATION, response: JSON.parse(evt.data) });
		};
	}, [rrmessages]);

	return (
		<React.Fragment>
			<div className="row border" style={{ backgroundColor: "white" }}>
				<div className="col-4">
					<div className="row h-100  align-items-center">
						<div className="col-12 d-flex justify-content-between ">
							<span className="justify-content-center">
								<b>{getUserInfo().Username}</b>
							</span>

							<i className="fa fa-pencil-square-o justify-content-center" style={{ fontSize: "30px", cursor: "pointer" }} onClick={openUserSelectModal} />
						</div>
					</div>
				</div>
				<div className="col-8 border-left">
					<UserChatHeader />
				</div>
			</div>
			<div className="row border-right border-left border-bottom" style={{ backgroundColor: "white" }}>
				<div className="col-4" style={{ overflowX: "hidden", minHeight: "780px", maxHeight: "780px" }}>
					<ConversationsSidebar />
				</div>
				{messageState.selectUserModal.selectedUser.Id === "" ? (
					<div className="col-8 align-items-center border-left d-flex flex-column" style={{ minHeight: "780px", maxHeight: "780px" }}>
						<FrontPageChat openUserSelectModal={openUserSelectModal} />
					</div>
				) : (
					<div className="col-8 border-left d-flex flex-column" style={{ minHeight: "780px", maxHeight: "780px" }}>
						<Chat />
						<ChatForm />
					</div>
				)}
			</div>
		</React.Fragment>
	);
};

export default UsersConversations;
