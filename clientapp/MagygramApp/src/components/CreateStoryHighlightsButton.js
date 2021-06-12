import { useContext } from "react";
import { modalConstants } from "../constants/ModalConstants";
import { StoryContext } from "../contexts/StoryContext";
import { storyService } from "../services/StoryService";

const CreateStoryHighlightsButton = ({ userId }) => {
	const { dispatch } = useContext(StoryContext);

	const getUserStoriesHandler = async () => {
		await storyService.findAllUserStories(dispatch);
	};

	const handleOpenStory = () => {
		dispatch({ type: modalConstants.SHOW_STORY_SELECT_HIGHLIGHTS_MODAL });
		getUserStoriesHandler();
	};
	return (
		<li class="list-inline-item" hidden={userId !== localStorage.getItem("userId")}>
			<button onClick={handleOpenStory} className="btn p-0 m-0">
				<div className="d-flex flex-column align-items-center">
					<div className="btn-outline-secondary rounded-circle overflow-hidden d-flex justify-content-center align-items-center border border-danger story-profile-photo-add-story">
						<i className="mdi mdi-plus w-50 h-50"></i>
					</div>
					<small>New</small>
				</div>
			</button>
		</li>
	);
};

export default CreateStoryHighlightsButton;
