import { useContext } from "react";
import { modalConstants } from "../constants/ModalConstants";
import { StoryContext } from "../contexts/StoryContext";

const StoryButton = () => {
	const { dispatch } = useContext(StoryContext);

	const handleOpenStory = () => {
		dispatch({ type: modalConstants.OPEN_CREATE_STORY_MODAL });
	};

	return (
		<div className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border border-danger post-profile-photo mr-3">
			<button type="button" onClick={handleOpenStory} className="btn btn-secondary rounded-lg btn-icon w-100 h-100">
				<i className="mdi mdi-plus w-50 h-50"></i>
			</button>
		</div>
	);
};

export default StoryButton;
