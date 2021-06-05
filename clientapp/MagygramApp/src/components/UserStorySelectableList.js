import React, { useContext } from "react";
import { StoryContext } from "../contexts/StoryContext";
import UserStorySelectableItem from "./UserStorySelectableItem";

const UserStorySelectableList = ({ selectedStoryIds, setSelectedStoryIds }) => {
	const { storyState } = useContext(StoryContext);

	const handleSelectStoryItem = (storyId) => {
		let pom = [...selectedStoryIds];
		pom.push(storyId);
		setSelectedStoryIds(pom);
	};

	const handleDeselectStoryItem = (storyId) => {
		setSelectedStoryIds(selectedStoryIds.filter((id) => id !== storyId));
	};

	return (
		<React.Fragment>
			<div className="row">
				{storyState.highlights.stories.map((story) => {
					return (
						<div className="col-4" key={story.id}>
							<UserStorySelectableItem time={story.dateTime} media={story.media} selectStoryItem={handleSelectStoryItem} deselectStoryItem={handleDeselectStoryItem} id={story.id} />
						</div>
					);
				})}
			</div>
		</React.Fragment>
	);
};

export default UserStorySelectableList;
