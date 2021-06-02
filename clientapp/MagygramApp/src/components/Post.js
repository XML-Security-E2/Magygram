import React, { useContext } from "react";
import PostComments from "./PostComments";
import PostHeader from "./PostHeader";
import PostImageSlider from "./PostImageSlider";
import PostInformation from "./PostInformation";
import PostInteraction from "./PostInteraction";
import { postService } from "../services/PostService";
import { PostContext } from "../contexts/PostContext";
import { modalConstants } from "../constants/ModalConstants";

const Post = ({ post }) => {
	const { dispatch } = useContext(PostContext);

	const LikePost = (postId) => {
		postService.likePost(postId, dispatch);
	};

	const UnlikePost = (postId) => {
		postService.unlikePost(postId, dispatch);
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
		console.log("DEL");
		postService.deletePostFromCollection(postId, dispatch);
	};

	return (
		<React.Fragment>
			<div className="d-flex flex-column mt-4 mb-4">
				<div className="card">
					<PostHeader username={post.UserInfo.Username} image={post.UserInfo.ImageURL} />
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
							<PostInformation username={post.UserInfo.Username} likes={post.LikedBy.length} dislikes={post.DislikedBy.length} description={post.Description} />
						</div>
						<PostComments comments={post.Comments} />
					</div>
				</div>
			</div>
		</React.Fragment>
	);
};

export default Post;
