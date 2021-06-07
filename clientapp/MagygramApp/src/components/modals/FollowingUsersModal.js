import { useContext } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { UserContext } from "../../contexts/UserContext";
import UserFollowingModalList from "../UserFollowingModalList";

const FollowingUsersModal = () => {
	const { userState, dispatch } = useContext(UserContext);

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_FOLLOWING_MODAL });
	};

	return (
		<Modal show={userState.userProfileFollowingModal.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Header closeButton>
				<Modal.Title id="contained-modal-title-vcenter">{userState.userProfileFollowingModal.modalHeader}</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				<UserFollowingModalList />
			</Modal.Body>
		</Modal>
	);
};

export default FollowingUsersModal;
