import React, { useContext } from "react";
import { PostContext } from "../contexts/PostContext";
import { hasPermissions } from "../helpers/auth-header";
import { postService } from "../services/PostService";

const UserProfilePosts = () => {
	const { postState, dispatch } = useContext(PostContext);

	const getPostDetailsHandler = async (postId) => {
		if(localStorage.getItem("userId")===null && !hasPermissions(["visit_private_profiles"])){
			await postService.findPostByIdForGuest(postId, dispatch);
		}else{
			await postService.findPostById(postId, dispatch);
		}
	};

	return (
		<div className="row " hidden={!postState.userProfileContent.showPosts}>
			{postState.userProfileContent.posts !== null ? (
				postState.userProfileContent.posts.map((post) => {
					return (
						<div className="col-4 box" style={{ cursor: "pointer" }} onClick={() => getPostDetailsHandler(post.id)}>
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
					<h3>User has not uploaded yet</h3>
				</div>
			)}
		</div>
	);
};

export default UserProfilePosts;
