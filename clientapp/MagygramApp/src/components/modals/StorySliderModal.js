import { useContext, useState } from "react";
import { Modal } from "react-bootstrap";
import { StoryContext } from "../../contexts/StoryContext";
import { modalConstants } from "../../constants/ModalConstants";
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import Stories from "react-insta-stories";
import { storyService } from "../../services/StoryService";
import OptionsModalStory from "../modals/OptionsModalStory";
import { MessageContext } from "../../contexts/MessageContext";
import SendStoryAsMessageModal from "./SendStoryAsMessageModal";

const StorySliderModal = () => {
	const { storyState, dispatch } = useContext(StoryContext);
	const msgCtx = useContext(MessageContext);

	const [tags, setTags] = useState([]);
	const [website, setWebsite] = useState("");
	const [contentType, setContentType] = useState("");

	const [storyId, setStoryId] = useState("");

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_STORY_SLIDER_MODAL, stories: storyState.storySliderModal.stories });
	};

	const handleRedirect = (userId) => {
		handleModalClose();
		window.location = "#/profile?userId=" + userId;
	};

	const onAllStoriesEnd = (test) => {
		//alert(test)
	};

	const onStoryStart = (index, story) => {
		console.log(index);
		console.log(story);
		setStoryId(story.header.storyId);
		setTags(story.tags);

		setContentType(story.contentType);
		setWebsite(story.website);

		setStoryId(story.header.storyId);
		if (!storyState.storySliderModal.visited) {
			if (index + 1 === storyState.storySliderModal.stories.length) {
				storyService.visitedByUser(story.header.storyId, dispatch);
				storyService.findStoriesForStoryline(dispatch);
			} else {
				storyService.visitedByUser(story.header.storyId, dispatch);
			}
		}
	};

	const handleOpenOptionsModal = () => {
		dispatch({ type: modalConstants.SHOW_STORY_OPTIONS_MODAL, storyId: storyId });
	};

	const handleOpenStorySendModal = () => {
		msgCtx.dispatch({ type: modalConstants.SHOW_SEND_STORY_TO_USER_MODAL, storyId });
	};

	return (
		<Modal size="md" show={storyState.storySliderModal.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Body className="modal-body-remove-margins modal-content-remove-margins">
				<div className={contentType === "CAMPAIGN" ? "d-flex justify-content-between pt-1" : "d-flex justify-content-end pt-1"} style={{ "background-color": "#111111" }}>
					{contentType === "CAMPAIGN" && <span className="text-white ml-3">Sponsored</span>}
					<button className="btn p-0 mr-3" onClick={handleOpenOptionsModal}>
						<i className="fa fa-ellipsis-h text-white" aria-hidden="true"></i>
					</button>
				</div>
				<Stories
					currentIndex={storyState.storySliderModal.firstUnvisitedStory}
					width="100%"
					stories={storyState.storySliderModal.stories}
					onStoryStart={onStoryStart}
					onAllStoriesEnd={onAllStoriesEnd}
				/>
				{contentType === "CAMPAIGN" && (
					<div className="d-flex align-items-center" style={{ "background-color": "#111111" }}>
						<a type="button" className="btn btn-link border-0 text-white" href={"https://" + website} target="_blank">
							Visit {website}
						</a>
					</div>
				)}
				<div style={{ "background-color": "#111111" }}>
					<label className="m-1 ml-2 text-white">Tagged users: </label>
					{tags !== null &&
						tags.map((tag) => {
							return (
								<button type="button" className="btn btn-dark m-1" onClick={() => handleRedirect(tag.Id)}>
									{tag.Username}
								</button>
							);
						})}
				</div>

				<div className="d-flex justify-content-end pb-1" style={{ "background-color": "#111111" }}>
					<button className="btn p-0 mr-3" onClick={handleOpenStorySendModal}>
						<i className="fa fa-paper-plane-o text-white" aria-hidden="true" style={{ fontSize: "20px" }}></i>
					</button>
				</div>
				<OptionsModalStory />
				<SendStoryAsMessageModal storyId={storyId} />
			</Modal.Body>
		</Modal>
	);
};

export default StorySliderModal;
