import { useContext } from "react";
import { modalConstants } from "../constants/ModalConstants";
import { StoryContext } from "../contexts/StoryContext";
import { hasRoles } from "../helpers/auth-header";

const StoryButton = () => {
	const { dispatch } = useContext(StoryContext);

	const handleOpenStory = () => {
		if (hasRoles(["agent"])) {
			dispatch({ type: modalConstants.OPEN_CREATE_AGENT_STORY_MODAL });
		} else {
			dispatch({ type: modalConstants.OPEN_CREATE_STORY_MODAL });
		}
	};

	return (
		<li class="list-inline-item">
			<button onClick={handleOpenStory} className="btn p-0 m-0">
				<div className="d-flex flex-column align-items-center">
					<div className="btn-secondary rounded-circle overflow-hidden d-flex justify-content-center align-items-center border border-danger story-profile-photo-add-story">
						<i className="mdi mdi-plus w-50 h-50"></i>
					</div>
					<small>Add story</small>
				</div>
			</button>
		</li>
	);
};

export default StoryButton;
