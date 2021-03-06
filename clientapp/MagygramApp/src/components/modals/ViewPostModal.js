import React, { useContext, useEffect, useState } from "react";
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
import OptionsModal from "./OptionsModal";
import SearchInfluencerModal from "./SearchInfluencerModal";
import { getUserInfo, hasRoles } from "../../helpers/auth-header";

const ViewPostModal = () => {
	const { postState, dispatch } = useContext(PostContext);
	const [agent, setAgent] = useState("");
	const style = { width: "450px" };

	const LikePost = (postId) => {
		postService.likePost(postId, getUserInfo(), dispatch);
	};

	useEffect(() => {
		if (hasRoles(["agent"])) setAgent(false);
		else setAgent(true);
	});

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
		dispatch({ type: modalConstants.HIDE_VIEW_POST_MODAL, post: postState.viewPostModal.post });
	};

	const postComment = (comment, tags) => {
		if (comment.length >= 1) {
			let postDTO = {
				PostId: postState.viewPostModal.post.Id,
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
		dispatch({ type: modalConstants.SHOW_POST_LIKED_BY_DETAILS, LikedBy: postState.viewPostModal.post.LikedBy });
	};

	const showDislikesModal = () => {
		dispatch({ type: modalConstants.SHOW_POST_DISLIKES_MODAL, Dislikes: postState.viewPostModal.post.DislikedBy });
	};

	const handleOpenOptionsModal = () => {
		dispatch({ type: modalConstants.SHOW_POST_OPTIONS_MODAL, post: postState.viewPostModal.post });
	};

	const searchInfluencer = () => {
		dispatch({ type: modalConstants.SHOW_SEARCH_INFLUENCER_MODAL, post: postState.viewPostModal.post });
	};

	const handleRedirect = (userId) => {
		handleModalClose();
		window.location = "#/profile?userId=" + userId;
	};

	const handleClickOnWebsite = async () => {
		await postService.clickOnPostCampaignWebsite(postState.viewPostModal.post.Id).then(handleOpenWebsite());
	};

	const handleOpenWebsite = () => {
		return new Promise(function () {
			window.open("https://" + postState.viewAgentCampaignPostModal.post.Website, "_blank");
		});
	};

	return (
		<Modal size="xl" show={postState.viewPostModal.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Body>
				<div className="d-flex flex-row align-items-top">
					<PostImageSliderModalView media={postState.viewPostModal.post.Media} website={postState.viewPostModal.post.Website} contentType={postState.viewPostModal.post.ContentType} />
					<div className="p-2" style={style}>
						<div className="align-top" style={style}>
							<div>
								{postState.viewPostModal.post.Tags !== null &&
									postState.viewPostModal.post.Tags.map((tag) => {
										return (
											<button type="button" className="btn btn-light mb-1 ml-1" onClick={() => handleRedirect(tag.Id)}>
												@{tag.Username}
											</button>
										);
									})}
							</div>
							<PostHeaderModalView
								username={postState.viewPostModal.post.UserInfo.Username}
								image={postState.viewPostModal.post.UserInfo.ImageURL}
								id={postState.viewPostModal.post.UserInfo.Id}
								openOptionsModal={handleOpenOptionsModal}
								location={postState.viewPostModal.post.Location}
							/>
							{postState.viewAgentCampaignPostModal.post.ContentType === "CAMPAIGN" && (
								<div className="border-bottom d-flex align-items-center">
									<span className="text-dark ml-2">Sponsored</span>
									<button type="button" className="btn btn-link border-0" onClick={handleClickOnWebsite}>
										Visit {postState.viewAgentCampaignPostModal.post.Website}
									</button>
								</div>
							)}
							<hr></hr>
							<div>
								{!agent && (
									<button style={({ height: "40px" }, { verticalAlign: "center" })} className="btn btn-outline-secondary" type="button" onClick={() => searchInfluencer()}>
										<i className="icofont-subscribe mr-1"></i>Product placement
									</button>
								)}
							</div>
							<PostCommentsModalView
								imageUrl={postState.viewPostModal.post.UserInfo.ImageURL}
								username={postState.viewPostModal.post.UserInfo.Username}
								description={postState.viewPostModal.post.Description}
								comments={postState.viewPostModal.post.Comments}
								handleRedirect={handleRedirect}
							/>
							<hr></hr>
							<div id="viewPostModalInteraction">
								<PostInteractionModalView
									post={postState.viewPostModal.post}
									LikePost={LikePost}
									DislikePost={DislikePost}
									UnlikePost={UnlikePost}
									UndislikePost={UndislikePost}
									addToDefaultCollection={addToDefaultCollection}
									deleteFromCollections={deleteFromCollections}
								/>
								<PostLikesAndDislikesModalView
									likes={postState.viewPostModal.post.LikedBy.length}
									dislikes={postState.viewPostModal.post.DislikedBy.length}
									showLikedByModal={showLikedByModal}
									showDislikesModal={showDislikesModal}
								/>
								<PostCommentInputModalView postComment={postComment} />
							</div>
						</div>
					</div>
				</div>
				<OptionsModal />
				<SearchInfluencerModal />
			</Modal.Body>
		</Modal>
	);
};

export default ViewPostModal;
