import React, { useContext } from "react";
import PostComments from "./PostComments";
import PostHeader from "./PostHeader";
import PostImageSlider from "./PostImageSlider";
import PostInformation from "./PostInformation";
import PostInteraction from "./PostInteraction";
import { postService } from "../services/PostService";
import { PostContext } from "../contexts/PostContext";
import { UserContext } from "../contexts/UserContext";
import PostLikesModal from "./modals/PostLikesModal";
import { modalConstants } from "../constants/ModalConstants";
import PostDislikesModal from "./modals/PostDislikesModal";
import ViewPostModal from "./modals/ViewPostModal";
import { getUserInfo } from "../helpers/auth-header";

const Post = ({ post }) => {
	const { dispatch } = useContext(PostContext);
	const { userState } = useContext(UserContext);

	const LikePost = (postId) => {
		postService.likePost(postId, getUserInfo(), dispatch);
	};

	const UnlikePost = (postId) => {
		postService.unlikePost(postId, getUserInfo(), dispatch);
	};

	const DislikePost = (postId) => {
		postService.dislikePost(postId, dispatch);
	};

	const UndislikePost = (postId) => {
		postService.undislikePost(postId, dispatch);
	};

	const showAddToCollectionModal = (postId) => {
		dispatch({ type: modalConstants.OPEN_ADD_TO_COLLECTION_MODAL, postId });
	};

	const addToDefaultCollection = (postId) => {
		let collectionDTO = {
			postId,
			collectionName: "",
		};
		postService.addPostToCollection(collectionDTO, dispatch);
	};

	const deleteFromCollections = (postId) => {
		postService.deletePostFromCollection(postId, dispatch);
	};

	const showLikedByModal = () => {
		dispatch({ type: modalConstants.SHOW_POST_LIKED_BY_DETAILS, LikedBy: post.LikedBy });
	};

	const showDislikesModal = () => {
		dispatch({ type: modalConstants.SHOW_POST_DISLIKES_MODAL, Dislikes: post.DislikedBy });
	};

	const postComment = (comment) => {
		if (comment.length >= 1) {
			let postDTO = {
				PostId: post.Id,
				Content: comment,
			};

			postService.commentPost(postDTO, dispatch);
		}
	};

	const viewAllComments = () => {
		dispatch({ type: modalConstants.SHOW_VIEW_POST_MODAL, post });
	};

	const handleOpenOptionsModal = () => {
		dispatch({ type: modalConstants.SHOW_POST_OPTIONS_MODAL, post });
	};

	return (
		<React.Fragment>
			<div className="d-flex flex-column mt-4 mb-4">
				<div className="card">
					<PostHeader username={post.UserInfo.Username} image={post.UserInfo.ImageURL} openOptionsModal={handleOpenOptionsModal} />
					<div className="card-body p-0">
						<PostImageSlider media={post.Media} />
						<PostInteraction
							post={post}
							LikePost={LikePost}
							DislikePost={DislikePost}
							UnlikePost={UnlikePost}
							UndislikePost={UndislikePost}
							showCollectionModal={showAddToCollectionModal}
							addToDefaultCollection={addToDefaultCollection}
							deleteFromCollections={deleteFromCollections}
						/>
						<div className="pl-3 pr-3 pb-2">
							<PostInformation
								username={post.UserInfo.Username}
								likes={post.LikedBy.length}
								dislikes={post.DislikedBy.length}
								description={post.Description}
								showLikedByModal={showLikedByModal}
								showDislikesModal={showDislikesModal}
							/>
						</div>
						<PostComments comments={post.Comments} postComment={postComment} viewAllComments={viewAllComments} />
						<PostLikesModal />
						<PostDislikesModal />
						<ViewPostModal />
					</div>
				</div>
			</div>
		</React.Fragment>
	);
};

export default Post;
