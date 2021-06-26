import { useContext } from "react";
import { Link } from "react-router-dom";
import { MessageContext } from "../../contexts/MessageContext";

const UserChatHeader = () => {
	const { messageState } = useContext(MessageContext);
	const imgStyle = { transform: "scale(1.5)", width: "100%", position: "absolute", left: "0" };

	return (
		<div className="d-flex flex-row align-items-center">
			<div
				className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border message-profile-photo m-2"
				style={messageState.selectUserModal.selectedUser.Id === "" ? { visibility: "hidden" } : {}}
			>
				<img style={imgStyle} src={messageState.selectUserModal.selectedUser.ImageURL === "" ? "assets/img/profile.jpg" : messageState.selectUserModal.selectedUser.ImageURL} alt="" />
			</div>
			<div className="profile-info ml-2">
				<Link className="profile-info-username text-dark" to={"/profile?userId=" + messageState.selectUserModal.selectedUser.Id}>
					{messageState.selectUserModal.selectedUser.Username}
				</Link>
			</div>
		</div>
	);
};

export default UserChatHeader;
