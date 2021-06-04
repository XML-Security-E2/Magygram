import React, { useContext, useEffect } from "react";
import { colorConstants } from "../constants/ColorConstants";
import { StoryContext } from "../contexts/StoryContext";
import { storyService } from "../services/StoryService";
import CreateStoryHighlightsButton from "./CreateStoryHighlightsButton";
import StorySliderModal from "./modals/StorySliderModal";
import Story from "./Story";

const StoryHighlights = () => {
	const { storyState, dispatch } = useContext(StoryContext);

	useEffect(() => {
		const getStoriesHandler = async () => {
			await storyService.findAllUserStories(dispatch);
		};
		getStoriesHandler();
	}, [dispatch]);

	const openStorySlider = (userId) => {
		storyService.GetStoriesForUser(userId, dispatch);
	};

	return (
		<React.Fragment>
			<div className="card border-0" style={{ backgroundColor: colorConstants.COLOR_BACKGROUND }}>
				<div className="card-body d-flex justify-content-start">
					<ul className="list-unstyled mb-0">
						{storyState.storyline.stories.map((story) => {
							return <Story story={story} openStorySlider={openStorySlider} />;
						})}
						<StorySliderModal />
						<CreateStoryHighlightsButton />
					</ul>
				</div>
			</div>
		</React.Fragment>
	);
};

export default StoryHighlights;
