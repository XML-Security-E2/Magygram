import React, { useContext, useEffect } from "react";
import { PostContext } from "../contexts/PostContext";
import GuestPostView from "./GuestPostView";

const GuestTimeline = () => {
	const { postState, dispatch } = useContext(PostContext);

	return (
		<React.Fragment>
			{postState.guestTimeline.posts.map((post) => {
				return <GuestPostView post={post} />;
			})}
		</React.Fragment>
	);
};

export default GuestTimeline;
