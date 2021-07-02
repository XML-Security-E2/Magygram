import { useContext, useState } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import Stories from "react-insta-stories";
import { MessageContext } from "../../contexts/MessageContext";

const StoryMessageModal = () => {
	const { messageState, dispatch } = useContext(MessageContext);

	const [website, setWebsite] = useState("");
	const [contentType, setContentType] = useState("");

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_STORY_MESSAGE_MODAL });
	};

	const onStoryStart = (index, story) => {
		console.log(index);
		setContentType(story.contentType);
		setWebsite(story.website);
	};

	const onAllStoriesEnd = () => {
		//alert(test)
		dispatch({ type: modalConstants.HIDE_STORY_MESSAGE_MODAL });
	};

	return (
		<Modal size="md" show={messageState.storyModal.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Body className="modal-body-remove-margins modal-content-remove-margins">
				<div className="d-flex justify-content-start pt-1" style={{ "background-color": "#111111" }}>
					{contentType === "CAMPAIGN" && <span className="text-white ml-3">Sponsored</span>}
				</div>
				<Stories currentIndex={0} width="100%" stories={messageState.storyModal.stories} onAllStoriesEnd={onAllStoriesEnd} onStoryStart={onStoryStart} />
				{contentType === "CAMPAIGN" && (
					<div className="d-flex align-items-center" style={{ "background-color": "#111111" }}>
						<a type="button" className="btn btn-link border-0 text-white" href={"https://" + website} target="_blank">
							Visit {website}
						</a>
					</div>
				)}
			</Modal.Body>
		</Modal>
	);
};

export default StoryMessageModal;
