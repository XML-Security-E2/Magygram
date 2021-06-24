import React, { useContext, useEffect } from "react";
import { MessageContext } from "../../contexts/MessageContext";
import { getUserInfo } from "../../helpers/auth-header";
import { getDateTime } from "../../helpers/datetime-helper";

const Chat = () => {
	const { messageState } = useContext(MessageContext);

	useEffect(() => {
		var a = document.querySelector("#conversation");
		console.log(a);
		a.scrollTop = a.scrollHeight - a.clientHeight;
	}, [messageState.showedMessages]);

	return (
		<React.Fragment>
			<div className="row flex-grow-1" id="conversation" style={{ overflowX: "hidden", maxHeight: "95%" }}>
				<div className="col-12">
					<div className="row">
						{messageState.showedMessages.map((message) => {
							if (message.messageFromId === getUserInfo().Id) {
								return (
									<div className="col-12 mt-2">
										<div className="row">
											<div className="col-12 text-break ml-auto" style={{ maxWidth: "50%" }}>
												{message.messageType === "TEXT" && (
													<div className="float-right rounded-lg pl-3 pb-2 pr-3 pt-2" style={{ backgroundColor: "lightgray" }}>
														{message.text}
													</div>
												)}
											</div>
											<div className="col-12">
												<span className="float-right" style={{ fontSize: "0.8em" }}>
													{getDateTime(message.timestamp)}
												</span>
											</div>
										</div>
									</div>
								);
							} else {
								return (
									<div className="col-12 mt-2">
										<div className="row">
											<div className="col-12 text-break mr-auto" style={{ maxWidth: "50%" }}>
												{message.messageType === "TEXT" && <div className="float-left rounded-lg border pl-3 pb-2 pr-3 pt-2">{message.text}</div>}
											</div>
											<div className="col-12">
												<span className="float-left" style={{ fontSize: "0.8em" }}>
													{getDateTime(message.timestamp)}
												</span>
											</div>
										</div>
									</div>
								);
							}
						})}
					</div>
				</div>
			</div>
		</React.Fragment>
	);
};

export default Chat;
