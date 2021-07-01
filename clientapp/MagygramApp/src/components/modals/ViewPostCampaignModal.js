import { useContext } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { PostContext } from "../../contexts/PostContext";
import PostHeaderModalView from "../PostHeaderModalView";
import PostImageSliderModalView from "../PostImageSliderModalView";
import PostCommentsModalView from "../PostCommentsModalView";
import PostInteractionModalView from "../PostInteractionModalView";
import { postService } from "../../services/PostService";
import PostCommentInputModalView from "../PostCommentInputModalView";
import PostLikesAndDislikesModalView from "../PostLikesAndDislikesModalView";
import SearchInfluencerModal from "./SearchInfluencerModal";
import { getUserInfo } from "../../helpers/auth-header";
import { postConstants } from "../../constants/PostConstants";
import CampaignOptionsModal from "./PostCampaignOptionsModal";
import PostCampaignOptionsModal from "./PostCampaignOptionsModal";

const ViewPostCampaignModal = () => {
	const { postState, dispatch } = useContext(PostContext);
	const style = { width: "450px" };

	const LikePost = (postId) => {
		postService.likePost(postId, getUserInfo(), dispatch);
	};

	const UnlikePost = (postId) => {
		postService.unlikePost(postId, getUserInfo(), dispatch);
	};

	const DislikePost = (postId) => {
		postService.dislikePost(postId, getUserInfo(), dispatch);
	};

	const UndislikePost = (postId) => {
		postService.undislikePost(postId, getUserInfo(), dispatch);
	};

	const handleModalClose = () => {
		dispatch({ type: postConstants.SET_CAMPAIGN_POST_DETAILS_REQUEST });
	};

	const postComment = (comment, tags) => {
		if (comment.length >= 1) {
			let postDTO = {
				PostId: postState.viewAgentCampaignPostModal.post.Id,
				Content: comment,
				Tags: tags,
			};

			postService.commentPost(postDTO, dispatch);
		}
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
		dispatch({ type: modalConstants.SHOW_POST_LIKED_BY_DETAILS, LikedBy: postState.viewAgentCampaignPostModal.post.LikedBy });
	};

	const showDislikesModal = () => {
		dispatch({ type: modalConstants.SHOW_POST_DISLIKES_MODAL, Dislikes: postState.viewAgentCampaignPostModal.post.DislikedBy });
	};

	const handleOpenOptionsModal = () => {
		dispatch({ type: modalConstants.SHOW_POST_AGENT_OPTIONS_MODAL });
	};

	const searchInfluencer = () => {
		dispatch({ type: modalConstants.SHOW_SEARCH_INFLUENCER_MODAL, post: postState.viewAgentCampaignPostModal.post });
	};

	const handleRedirect = (userId) => {
		handleModalClose();
		window.location = "#/profile?userId=" + userId;
	};

	return (
		<Modal size="xl" show={postState.viewAgentCampaignPostModal.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Body>
				<div className="d-flex flex-row align-items-top">
					<PostImageSliderModalView media={postState.viewAgentCampaignPostModal.post.Media} />
					<div className="p-2" style={style}>
						<div className="align-top" style={style}>
							<div>
								{postState.viewAgentCampaignPostModal.post.Tags !== null &&
									postState.viewAgentCampaignPostModal.post.Tags.map((tag) => {
										return (
											<button type="button" className="btn btn-light mb-1 ml-1" onClick={() => handleRedirect(tag.Id)}>
												@{tag.Username}
											</button>
										);
									})}
							</div>
							<PostHeaderModalView
								username={postState.viewAgentCampaignPostModal.post.UserInfo.Username}
								image={postState.viewAgentCampaignPostModal.post.UserInfo.ImageURL}
								id={postState.viewAgentCampaignPostModal.post.UserInfo.Id}
								openOptionsModal={handleOpenOptionsModal}
								location={postState.viewAgentCampaignPostModal.post.Location}
							/>
							<hr></hr>
							<div>
								<button style={({ height: "40px" }, { verticalAlign: "center" })} className="btn btn-outline-secondary" type="button" onClick={() => searchInfluencer()}>
									<i className="icofont-subscribe mr-1"></i>Product placement
								</button>
							</div>
							<PostCommentsModalView
								imageUrl={postState.viewAgentCampaignPostModal.post.UserInfo.ImageURL}
								username={postState.viewAgentCampaignPostModal.post.UserInfo.Username}
								description={postState.viewAgentCampaignPostModal.post.Description}
								comments={postState.viewAgentCampaignPostModal.post.Comments}
								handleRedirect={handleRedirect}
							/>
							<hr></hr>
							<div id="viewAgentCampaignPostModalInteraction">
								<PostInteractionModalView
									post={postState.viewAgentCampaignPostModal.post}
									LikePost={LikePost}
									DislikePost={DislikePost}
									UnlikePost={UnlikePost}
									UndislikePost={UndislikePost}
									addToDefaultCollection={addToDefaultCollection}
									deleteFromCollections={deleteFromCollections}
								/>
								<PostLikesAndDislikesModalView
									likes={postState.viewAgentCampaignPostModal.post.LikedBy.length}
									dislikes={postState.viewAgentCampaignPostModal.post.DislikedBy.length}
									showLikedByModal={showLikedByModal}
									showDislikesModal={showDislikesModal}
								/>
								<PostCommentInputModalView postComment={postComment} />
							</div>
						</div>
					</div>
				</div>

				<PostCampaignOptionsModal postId={postState.viewAgentCampaignPostModal.post.Id} />
				<SearchInfluencerModal />
			</Modal.Body>
		</Modal>
	);
};

export default ViewPostCampaignModal;
