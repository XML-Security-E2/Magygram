import { useContext } from "react";
import { Modal } from "react-bootstrap";
import { StoryContext } from "../../contexts/StoryContext";
import { modalConstants } from "../../constants/ModalConstants";
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import Stories from 'react-insta-stories';

const StorySliderModal = () => {
	const { storyState, dispatch } = useContext(StoryContext);

    const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_STORY_SLIDER_MODAL, stories: storyState.storySliderModal.stories });
	}

	const onAllStoriesEnd = () => {
		alert('todo on stories end')
	}

	return (
		<Modal  size="md" show={storyState.storySliderModal.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Body className="modal-body-remove-margins modal-content-remove-margins">
				<Stories 
					width="100%" 
					stories={storyState.storySliderModal.stories}
					onAllStoriesEnd={onAllStoriesEnd}/>
			</Modal.Body>
		</Modal>
	);
};

export default StorySliderModal;
