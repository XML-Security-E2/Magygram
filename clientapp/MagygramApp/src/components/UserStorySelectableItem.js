import React from "react";

const UserStorySelectableItem = ({ media }) => {
	return (
		<React.Fragment>
			<img src={media.Url} alt="" className="img-fluid rounded-lg" />
		</React.Fragment>
	);
};

export default UserStorySelectableItem;
