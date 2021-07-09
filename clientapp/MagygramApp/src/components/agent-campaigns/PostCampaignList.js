import React, { useContext, useEffect } from "react";
import { PostContext } from "../../contexts/PostContext";
import { postService } from "../../services/PostService";

const PostCampaignList = () => {
	const { postState, dispatch } = useContext(PostContext);

	const getPostDetailsHandler = async (postId) => {
		await postService.findPostById(postId, dispatch);
	};

	useEffect(() => {
		const getPostsHandler = async () => {
			await postService.findAllUsersCampaignPosts(dispatch);
		};
		getPostsHandler();
	}, []);

	return (
		<React.Fragment>
			<h3 className="text-dark mt-5">Post campaigns</h3>
			<div className="row ">
				{postState.agentCampaignPosts !== null && postState.agentCampaignPosts.length > 0 ? (
					postState.agentCampaignPosts.map((post) => {
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
						<h4>User don't have active post campaigns</h4>
					</div>
				)}
			</div>
		</React.Fragment>
	);
};

export default PostCampaignList;
