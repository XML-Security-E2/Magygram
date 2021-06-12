import { useContext, useState } from "react";
import { Button, Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { storyConstants } from "../../constants/StoryConstants";
import { StoryContext } from "../../contexts/StoryContext";
import { storyService } from "../../services/StoryService";
import CreateHighlightNameForm from "../CreateHighlightNameForm";
import FailureAlert from "../FailureAlert";
import UserStorySelectableList from "../UserStorySelectableList";

const UserStoriesModal = () => {
	const { storyState, dispatch } = useContext(StoryContext);
	const [selectedStoryIds, setSelectedStoryIds] = useState([]);
	const [highlightName, setHighlightName] = useState("");

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_STORY_SELECT_HIGHLIGHTS_MODAL });
	};

	const handleShowNameInput = () => {
		if (selectedStoryIds.length > 0) dispatch({ type: storyConstants.SHOW_HIGHLIGHTS_NAME_INPUT });
		else dispatch({ type: storyConstants.SHOW_HIGHLIGHTS_MODAL_ERROR_MESSAGE, errorMessage: "You must select story" });
	};

	const handleSubmit = (e) => {
		e.preventDefault();

		let highlight = {
			storyIds: selectedStoryIds,
			name: highlightName,
		};

		console.log(highlight);
		storyService.createHighlight(highlight, dispatch);
		setSelectedStoryIds([]);
	};

	return (
		<Modal show={storyState.highlights.showModal} size="lg" aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Header closeButton>
				<Modal.Title id="contained-modal-title-vcenter">Select stories</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				<FailureAlert
					hidden={!storyState.highlights.showError}
					header="Error"
					message={storyState.highlights.errorMessage}
					handleCloseAlert={() => dispatch({ type: storyConstants.HIDE_HIGHLIGHTS_MODAL_ERROR_MESSAGE })}
				/>
				<div className="row">
					<UserStorySelectableList selectedStoryIds={selectedStoryIds} setSelectedStoryIds={setSelectedStoryIds} />
				</div>
			</Modal.Body>
			<Modal.Footer>
				<CreateHighlightNameForm handleSubmit={handleSubmit} setHighlightName={setHighlightName} />

				<Button hidden={storyState.highlights.showHighlightsName} onClick={handleShowNameInput}>
					Next
				</Button>
			</Modal.Footer>
		</Modal>
	);
};

export default UserStoriesModal;
