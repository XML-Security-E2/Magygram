import React, { useContext, useEffect } from "react";
import { StoryContext } from "../contexts/StoryContext";
import { storyService } from "../services/StoryService";
import StorySliderModal from "./modals/StorySliderModal";
import Story from "./Story";
import StoryButton from "./StoryButton";
import MyStoryButton from "../components/MyStoryButton";

const Storyline = () => {
	const { storyState, dispatch } = useContext(StoryContext);

	useEffect(() => {
		const getStoriesHandler = async () => {
			await storyService.findStoriesForStoryline(dispatch);
		};
		getStoriesHandler();
		storyService.HaveActiveStoriesLoggedUser(dispatch)
	}, [dispatch]);

	const openStorySlider = (userId,visited) => {
		storyService.GetStoriesForUser(userId, visited, dispatch);
	};

	return (
		<React.Fragment>
			<div className="card">
				<div className="card-body d-flex justify-content-start">
					<ul className="list-unstyled mb-0">
						<StoryButton />
						<MyStoryButton openStorySlider={openStorySlider} />

						{storyState.storyline.stories !== null &&
							storyState.storyline.stories.map((story) => {
								return <Story story={story} openStorySlider={openStorySlider} />;
						})}
						<StorySliderModal/>
						
					</ul>
				</div>
			</div>
		</React.Fragment>
	);
};

export default Storyline;
