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

const StorySliderForAdmin = () => {
	const { storyState, dispatch } = useContext(StoryContext);

	const [tags, setTags] = useState([]);
	const [storyId, setStoryId] = useState("");

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_STORY_SLIDER_MODAL_ADMIN, stories: storyState.storySliderModal.stories });
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
		setStoryId(story.header.storyId)
		console.log(storyState.storySliderModal.userId);
		setTags(story.tags);
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


	return (
		<Modal size="md" show={storyState.storySliderModal.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Body className="modal-body-remove-margins modal-content-remove-margins">
				<div className="d-flex justify-content-end" style={{ "background-color": "black" }}>
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
				<div style={{ "background-color": "black" }}>
					<label className="m-1 text-white">Tagged users: </label>
					{tags !== null &&
						tags.map((tag) => {
							return (
								<button type="button" className="btn btn-dark m-1" onClick={() => handleRedirect(tag.Id)}>
									{tag.Username}
								</button>
							);
						})}
				</div>

				<OptionsModalStory />
			</Modal.Body>
		</Modal>
	);
};

export default StorySliderForAdmin;
