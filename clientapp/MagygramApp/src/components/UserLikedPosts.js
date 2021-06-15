import React, { useContext, useEffect } from "react";
import { PostContext } from "../contexts/PostContext";
import { postService } from "../services/PostService";

const UserLikedPosts = () => {
	const { postState, dispatch } = useContext(PostContext);

	const getPostDetailsHandler = async (postId) => {
		await postService.findPostById(postId, dispatch);
	};

	const getPostsHandler = async () => {
		await postService.findAllLikedPosts(dispatch);
	};

    useEffect(() => {
		getPostsHandler();
	}, []);

	return (
		<div className="row ">
			{postState.userLikedPosts !== null ? (
				postState.userLikedPosts.map((post) => {
					return (
						<div className="col-3 mb-4" style={{ cursor: "pointer" }} onClick={() => getPostDetailsHandler(post.id)}>
							{post.media.MediaType === "IMAGE" ? (
								<img src={post.media.Url} className="img-fluid box-coll rounded-lg w-100 " alt="" style={{ objectFit: "cover" }} />
							) : (
								<video className="img-fluid box-coll rounded-lg w-100" style={{ objectFit: "cover" }}>
									<source src={post.media.Url} type="video/mp4" />
								</video>
							)}
						</div>
					);
				})
			) : (
				<div className="col-12 mt-5 d-flex justify-content-center text-secondary">
					<h3>User has not liked posts</h3>
				</div>
			)}
		</div>
	);
};

export default UserLikedPosts;
