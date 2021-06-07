import React, { useContext, useEffect } from "react";
import { PostContext } from "../contexts/PostContext";
import { postService } from "../services/PostService";
import ViewPostModal from "./modals/ViewPostModal";
import StoryHighlights from "./StoryHighlights";
import UserProfilePosts from "./UserProfilePosts";
import UserProfileSavedPosts from "./UserProfileSavedPosts";

const UserProfileContent = ({ userId }) => {
	const { postState, dispatch } = useContext(PostContext);

	const getPostsHandler = async () => {
		await postService.findAllUserPosts(userId, dispatch);
	};

	const getCollectionsHandler = async () => {
		await postService.findAllUsersProfileCollections(dispatch);
	};

	useEffect(() => {
		getPostsHandler();
	}, [userId]);

	return (
		<React.Fragment>
			{!postState.userProfileContent.showUnauthorizedErrorMessage ? (
				<div>
					<StoryHighlights userId={userId} />
					<hr />
					<nav className="mb-5">
						<div className="container d-flex justify-content-center">
							<button disabled={postState.userProfileContent.showPosts} type="button" onClick={getPostsHandler} className="btn btn-outline-secondary btn-icon-text border-0 mr-2">
								<i className="fa fa-th"></i> POSTS
							</button>
							<button
								hidden={userId !== localStorage.getItem("userId")}
								disabled={postState.userProfileContent.showCollections}
								type="button"
								onClick={getCollectionsHandler}
								className="btn btn-outline-secondary btn-icon-text border-0 mr-2 ml-2"
							>
								<i className="fa fa-bookmark"></i> SAVED
							</button>
							<button hidden={userId !== localStorage.getItem("userId")} type="button" className="btn btn-outline-secondary btn-icon-text border-0 ml-2">
								<i className="fa fa-tag"></i> TAGGED
							</button>
						</div>
					</nav>
					<UserProfilePosts />
					<UserProfileSavedPosts />
					<ViewPostModal />
				</div>
			) : (
				<div className="mt-5 d-flex justify-content-center border" style={{ backgroundColor: "white" }}>
					<div>
						<h3 className="d-flex justify-content-center">This profile is private</h3>
						<p className="text-secondary">Only users who follow this user can see photos and videos he/she uploaded</p>
					</div>
				</div>
			)}
		</React.Fragment>
	);
};

export default UserProfileContent;
