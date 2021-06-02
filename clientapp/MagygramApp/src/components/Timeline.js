import React, { useContext, useEffect } from "react";
import { PostContext } from "../contexts/PostContext";
import Post from "./Post";
import "react-image-gallery/styles/css/image-gallery.css";
import { postService } from "../services/PostService";

const Timeline = () => {
	const { postState, dispatch } = useContext(PostContext);

	useEffect(() => {
		const getPostsHandler = async () => {
			await postService.findPostsForTimeline(dispatch);
		};
		getPostsHandler();
	}, [dispatch]);

	return (
		<React.Fragment>
			{postState.timeline.posts.map((post) => {
				console.log(post);
				return <Post post={post} />;
			})}
		</React.Fragment>
	);
};

export default Timeline;
