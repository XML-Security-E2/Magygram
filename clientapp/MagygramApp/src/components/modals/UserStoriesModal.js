import { useContext } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { StoryContext } from "../../contexts/StoryContext";
import UserStorySelectableList from "../UserStorySelectableList";

const UserStoriesModal = () => {
	const { storyState, dispatch } = useContext(StoryContext);

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_STORY_SELECT_HIGHLIGHTS_MODAL });
	};
	return (
		<Modal show={storyState.highlights.showModal} size="lg" aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Header closeButton>
				<Modal.Title id="contained-modal-title-vcenter">Select collection</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				<div className="row">
                    <UserStorySelectableList/>
                </div>
			</Modal.Body>
		</Modal>
	);
};

export default UserStoriesModal;
