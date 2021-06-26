import { useContext } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import Stories from "react-insta-stories";
import { MessageContext } from "../../contexts/MessageContext";

const StoryMessageModal = () => {
	const { messageState, dispatch } = useContext(MessageContext);

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_STORY_MESSAGE_MODAL });
	};

	const onAllStoriesEnd = () => {
		//alert(test)
		dispatch({ type: modalConstants.HIDE_STORY_MESSAGE_MODAL });
	};

	return (
		<Modal size="md" show={messageState.storyModal.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Body className="modal-body-remove-margins modal-content-remove-margins">
				<Stories currentIndex={0} width="100%" stories={messageState.storyModal.stories} onAllStoriesEnd={onAllStoriesEnd} />
			</Modal.Body>
		</Modal>
	);
};

export default StoryMessageModal;
