import { useContext } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import Stories from "react-insta-stories";
import { StoryContext } from "../../contexts/StoryContext";

const StoryCampaignModal = () => {
	const { storyState, dispatch } = useContext(StoryContext);

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_STORY_AGENT_CAMPAIGN_MODAL });
	};

	const onAllStoriesEnd = () => {
		//alert(test)
		dispatch({ type: modalConstants.HIDE_STORY_AGENT_CAMPAIGN_MODAL });
	};

	const handleOpenOptionsModal = () => {
		dispatch({ type: modalConstants.SHOW_STORY_AGENT_OPTIONS_MODAL });
	};

	return (
		<Modal size="md" show={storyState.agentCampaignStoryModal.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Body className="modal-body-remove-margins modal-content-remove-margins">
				<div className="d-flex justify-content-end" style={{ "background-color": "black" }}>
					<button className="btn p-0 mr-3" onClick={handleOpenOptionsModal}>
						<i className="fa fa-ellipsis-h text-white" aria-hidden="true"></i>
					</button>
				</div>
				<Stories currentIndex={0} width="100%" stories={storyState.agentCampaignStoryModal.stories} onAllStoriesEnd={onAllStoriesEnd} />
			</Modal.Body>
		</Modal>
	);
};

export default StoryCampaignModal;
