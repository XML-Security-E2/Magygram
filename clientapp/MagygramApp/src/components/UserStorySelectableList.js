import React, { useContext } from "react";
import { StoryContext } from "../contexts/StoryContext";
import UserStorySelectableItem from "./UserStorySelectableItem";

const UserStorySelectableList = () => {
	const { storyState, dispatch } = useContext(StoryContext);

	return (
		<React.Fragment>
			<div className="row">
				{storyState.highlights.stories.map((story) => {
					return (
						<div className="col-4">
							<UserStorySelectableItem media={story.media} />
						</div>
					);
				})}
			</div>
		</React.Fragment>
	);
};

export default UserStorySelectableList;
