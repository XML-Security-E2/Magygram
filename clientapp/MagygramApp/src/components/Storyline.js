import React, {useContext, useEffect} from "react";
import { modalConstants } from "../constants/ModalConstants";
import { StoryContext } from "../contexts/StoryContext";
import { storyService } from "../services/StoryService";
import StorySliderModal from "./modals/StorySliderModal";
import Story from "./Story";
import StoryButton from "./StoryButton";

const Storyline = () => {
	const { storyState, dispatch } = useContext(StoryContext);

	useEffect(() => {
		const getStoriesHandler = async () => {
			await storyService.findStoriesForStoryline(dispatch);
		};
		getStoriesHandler();
	}, [dispatch]);

	const openStorySlider = (userId) =>{
		storyService.GetStoryForUser(userId, dispatch)
	}

	return (
        <React.Fragment>
            <div className="card">
				<div className="card-body d-flex justify-content-start">
					<ul className="list-unstyled mb-0">
						<StoryButton/>
						{storyState.storyline.stories.map((story) => {
							return <Story story={story} openStorySlider={openStorySlider}/>; })}
						<StorySliderModal/>

					</ul>
				</div>
			</div>
        </React.Fragment>
	);
};

export default Storyline;
