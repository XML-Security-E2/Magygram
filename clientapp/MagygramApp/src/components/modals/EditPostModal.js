import { useContext } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { PostContext } from "../../contexts/PostContext";
import EditPostForm from "../EditPostForm";

const EditPostModal = () => {
	const { postState, dispatch } = useContext(PostContext);

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_POST_EDIT_MODAL });
	};
	return (
		<Modal show={postState.editPost.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Header closeButton>
				<Modal.Title id="contained-modal-title-vcenter">Edit post</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				<EditPostForm />
			</Modal.Body>
		</Modal>
	);
};

export default EditPostModal;
