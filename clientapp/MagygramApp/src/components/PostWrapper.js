import React, { useContext, useEffect } from "react";
import { PostContext } from "../contexts/PostContext";
import { postService } from "../services/PostService";
import Post from "./Post";

const PostWrapper = ({ postId }) => {
	const { postState, dispatch } = useContext(PostContext);

	useEffect(() => {
		const getPostHandler = async () => {
			await postService.findPostByIdForPage(postId, dispatch);
		};
		getPostHandler();
	}, [dispatch]);

	return (
		<React.Fragment>
			<Post post={postState.postDetailsPage.post} />
		</React.Fragment>
	);
};

export default PostWrapper;
