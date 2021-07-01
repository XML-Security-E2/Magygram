import React, { useContext, useEffect } from "react";
import { StoryContext } from "../../contexts/StoryContext";
import { storyService } from "../../services/StoryService";
import StoryCampaignPreview from "./StoryCampaignPreview";

const StoryCampaignList = () => {
	const { storyState, dispatch } = useContext(StoryContext);

	useEffect(() => {
		const getStoriesHandler = async () => {
			await storyService.findAllUsersCampaignStories(dispatch);
		};
		getStoriesHandler();
	}, []);

	return (
		<React.Fragment>
			{storyState.agentCampaignStories !== null && <h3 className="text-dark">Story campaigns</h3>}

			<div className="row ">
				{storyState.agentCampaignStories !== null ? (
					storyState.agentCampaignStories.map((story) => {
						return (
							<div className="col-3">
								<StoryCampaignPreview story={story} storyId={story.Id} />
							</div>
						);
					})
				) : (
					<div className="col-12 mt-5 d-flex justify-content-center text-secondary">
						<h4>User don't have active story campaigns</h4>
					</div>
				)}
			</div>
		</React.Fragment>
	);
};

export default StoryCampaignList;
