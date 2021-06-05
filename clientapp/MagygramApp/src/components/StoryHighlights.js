import React, { useContext, useEffect } from "react";
import { colorConstants } from "../constants/ColorConstants";
import { StoryContext } from "../contexts/StoryContext";
import { storyService } from "../services/StoryService";
import CreateStoryHighlightsButton from "./CreateStoryHighlightsButton";
import Highlight from "./Highlight";
import StorySliderHighlightsModal from "./modals/StorySliderHighlightsModal";

const StoryHighlights = () => {
	const { storyState, dispatch } = useContext(StoryContext);

	useEffect(() => {
		const getHighlightsHandler = async () => {
			await storyService.findAllProfileHighlights(dispatch);
		};
		getHighlightsHandler();
	}, [dispatch]);

	const openHighlightSlider = (name) => {
		storyService.findAllStoriesByHighlightName(name, dispatch);
	};

	return (
		<React.Fragment>
			<div className="card border-0" style={{ backgroundColor: colorConstants.COLOR_BACKGROUND }}>
				<div className="card-body d-flex justify-content-start">
					<ul className="list-unstyled mb-0">
						{storyState.profileHighlights.highlights !== null &&
							storyState.profileHighlights.highlights.map((highlight) => {
								return <Highlight highlight={highlight} openHighlightSlider={openHighlightSlider} />;
							})}
						<StorySliderHighlightsModal />
						<CreateStoryHighlightsButton />
					</ul>
				</div>
			</div>
		</React.Fragment>
	);
};

export default StoryHighlights;
