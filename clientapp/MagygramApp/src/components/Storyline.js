import React, {useContext, useEffect} from "react";
import { StoryContext } from "../contexts/StoryContext";
import { storyService } from "../services/StoryService";
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

	return (
        <React.Fragment>
            <div className="card">
				<div className="card-body d-flex justify-content-start">
					<ul className="list-unstyled mb-0">
						<StoryButton/>
						{storyState.storyline.stories.map((story) => {
							return <Story story={story}/>; })}
					</ul>
				</div>
			</div>
        </React.Fragment>
	);
};

export default Storyline;
