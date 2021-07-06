import { useContext, useState } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import Stories from "react-insta-stories";
import { StoryContext } from "../../contexts/StoryContext";
import StoryCampaignOptionsModal from "./StoryCampaignOptionsModal";
import { storyService } from "../../services/StoryService";

const StoryCampaignModal = () => {
	const { storyState, dispatch } = useContext(StoryContext);

	const [website, setWebsite] = useState("");
	const [contentType, setContentType] = useState("");
	const [storyId, setStoryId] = useState("");

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_STORY_AGENT_CAMPAIGN_MODAL });
	};

	const onStoryStart = (index, story) => {
		console.log(index);
		console.log(story);

		setContentType(story.contentType);
		setWebsite(story.website);
		setStoryId(story.header.storyId);
	};

	const onAllStoriesEnd = () => {
		//alert(test)
		//dispatch({ type: modalConstants.HIDE_STORY_AGENT_CAMPAIGN_MODAL });
	};

	const handleOpenOptionsModal = () => {
		dispatch({ type: modalConstants.SHOW_STORY_AGENT_OPTIONS_MODAL });
	};

	const handleClickOnWebsite = async () => {
		await storyService.clickOnStoryCampaignWebsite(storyId).then(handleOpenWebsite());
	};

	const handleOpenWebsite = () => {
		return new Promise(function () {
			window.open("https://" + website, "_blank");
		});
	};

	return (
		<Modal size="md" show={storyState.agentCampaignStoryModal.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Body className="modal-body-remove-margins modal-content-remove-margins">
				<div className={contentType === "CAMPAIGN" ? "d-flex justify-content-between pt-1" : "d-flex justify-content-end pt-1"} style={{ "background-color": "#111111" }}>
					{contentType === "CAMPAIGN" && <span className="text-white ml-3">Sponsored</span>}
					<button className="btn p-0 mr-3" onClick={handleOpenOptionsModal}>
						<i className="fa fa-ellipsis-h text-white" aria-hidden="true"></i>
					</button>
				</div>
				<Stories currentIndex={0} width="100%" stories={storyState.agentCampaignStoryModal.stories} onAllStoriesEnd={onAllStoriesEnd} onStoryStart={onStoryStart} />
				{contentType === "CAMPAIGN" && (
					<div className="d-flex align-items-center" style={{ "background-color": "#111111" }}>
						<button type="button" className="btn btn-link text-white border-0" onClick={handleClickOnWebsite}>
							Visit {website}
						</button>
					</div>
				)}
				<StoryCampaignOptionsModal storyId={storyState.agentCampaignStoryModal.storyId} />
			</Modal.Body>
		</Modal>
	);
};

export default StoryCampaignModal;
