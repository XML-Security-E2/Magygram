import React, { useContext } from "react";
import { PostContext } from "../contexts/PostContext";
import GuestPostView from "./GuestPostView";

const GuestTimeline = () => {
	const { postState } = useContext(PostContext);

	return (
		<React.Fragment>
			{postState.guestTimeline.posts.map((post) => {
				return <GuestPostView post={post} />;
			})}
		</React.Fragment>
	);
};

export default GuestTimeline;
