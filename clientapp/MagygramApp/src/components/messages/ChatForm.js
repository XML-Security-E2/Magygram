import autosize from "autosize";
import { useContext, useEffect, useState } from "react";
import { MessageContext } from "../../contexts/MessageContext";
import { messageService } from "../../services/MessageService";

const ChatForm = () => {
	const { messageState, dispatch } = useContext(MessageContext);

	const [messageText, setMessageText] = useState("");
	const [showSendButton, setShowSendButton] = useState(false);

	const onTextInputChange = (e) => {
		setMessageText(e.target.value);
		if (e.target.value.length > 0) {
			setShowSendButton(true);
		} else {
			setShowSendButton(false);
		}
	};

	const sendMessage = () => {
		let message = {
			messageTo: messageState.selectUserModal.selectedUser.Id,
			messageType: "TEXT",
			text: messageText,
			contentUrl: "",
		};
		console.log(message);
		messageService.sendMessage(message, dispatch);
		setMessageText("");
		setShowSendButton(false);
	};

	useEffect(() => {
		autosize(document.querySelectorAll("textarea"));
	}, []);

	return (
		<div className="row mt-auto d-flex align-items-center border m-1" style={{ borderRadius: "10px" }}>
			<textarea id="textarea" className="form-control border-0 col-10" placeholder="Message..." rows="1" value={messageText} onChange={(e) => onTextInputChange(e)} />
			{showSendButton ? (
				<button type="button" className="btn btn-link  text-primary col-2" onClick={sendMessage}>
					Send
				</button>
			) : (
				<i className="fa fa-picture-o col-2" style={{ fontSize: "30px", cursor: "pointer" }} />
			)}
		</div>
	);
};

export default ChatForm;
