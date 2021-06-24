import React from "react";

const FrontPageChat = ({ openUserSelectModal }) => {
	return (
		<div className="d-flex align-items-center justify-content-center" style={{ minHeight: "780px", maxHeight: "780px" }}>
			<div>
				<div className="row d-flex align-items-center justify-content-center">
					<span style={{ fontSize: "1.2em", color: "gray" }} className="text-center">
						Send messages, photos or videos to your followers and friends
					</span>
					<button type="button" className="btn btn-primary mt-2" onClick={openUserSelectModal}>
						Send message
					</button>
				</div>
			</div>
		</div>
	);
};

export default FrontPageChat;
