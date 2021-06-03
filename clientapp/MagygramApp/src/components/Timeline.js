import React, { useContext, useEffect } from "react";
import { PostContext } from "../contexts/PostContext";
import Post from "./Post";
import "react-image-gallery/styles/css/image-gallery.css";
import { postService } from "../services/PostService";
import { Link } from "react-router-dom";

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
			<div className="d-flex flex-column mt-4 mb-4">
				<div className="card">
					<Link type="button" className="btn btn-link btn-fw text-secondary w-100 border-0" to="/add-posts">
						Create new post
					</Link>
				</div>
			</div>
			{postState.timeline.posts.map((post) => {
				console.log(post);
				return <Post post={post} />;
			})}
		</React.Fragment>
	);
};

export default Timeline;
