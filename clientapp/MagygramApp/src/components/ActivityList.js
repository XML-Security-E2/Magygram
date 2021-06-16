import React from "react";
import FollowRequestsList from "./FollowRequestsList";
import NotificationsList from "./NotificationsList";

const ActivityList = () => {
	return (
		<React.Fragment>
			<FollowRequestsList />
			<NotificationsList />
		</React.Fragment>
	);
};

export default ActivityList;
