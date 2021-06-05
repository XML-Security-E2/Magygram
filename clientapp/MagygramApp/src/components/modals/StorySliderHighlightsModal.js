import { useContext } from "react";
import { Modal } from "react-bootstrap";
import { StoryContext } from "../../contexts/StoryContext";
import { modalConstants } from "../../constants/ModalConstants";
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import Stories from "react-insta-stories";

const StorySliderHighlightsModal = () => {
	const { storyState, dispatch } = useContext(StoryContext);

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_STORY_SLIDER_HIGHLIGHTS_MODAL });
	};

	const onAllStoriesEnd = () => {
		//alert(test)
		dispatch({ type: modalConstants.HIDE_STORY_SLIDER_HIGHLIGHTS_MODAL });
	};

	return (
		<Modal size="md" show={storyState.highlightsSliderModal.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Body className="modal-body-remove-margins modal-content-remove-margins">
				<Stories currentIndex={0} width="100%" stories={storyState.highlightsSliderModal.highlights} onAllStoriesEnd={onAllStoriesEnd} />
			</Modal.Body>
		</Modal>
	);
};

export default StorySliderHighlightsModal;
