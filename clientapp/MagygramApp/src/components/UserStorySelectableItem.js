import React, { useState } from "react";

const UserStorySelectableItem = ({ time, media, deselectStoryItem, selectStoryItem, id }) => {
	const [selectedStory, setSelectedStory] = useState(false);

	const handleStoryItemClick = () => {
		if (selectedStory) {
			deselectStoryItem(id);
		} else {
			selectStoryItem(id);
		}
		setSelectedStory(!selectedStory);
	};

	return (
		<React.Fragment>
			<div className="container-select-img" onClick={handleStoryItemClick} style={{ cursor: "pointer" }}>
				{media.MediaType === "IMAGE" ? (
					<img src={media.Url} className="img-fluid rounded-lg w-100" alt="" />
				) : (
					<video className="img-fluid box-coll rounded-lg w-100" style={{ objectFit: "cover" }}>
						<source src={media.Url} type="video/mp4" />
					</video>
				)}

				<div className="story-time">{time}</div>
				<div hidden={!selectedStory} className="overlay-select-img rounded" style={{ backgroundColor: "rgba(83, 83, 83, 0.6)" }}>
					<button className="btn icon-select-img">
						<i className="fa fa-check" style={{ color: "white" }}></i>
					</button>
				</div>
			</div>
		</React.Fragment>
	);
};

export default UserStorySelectableItem;
