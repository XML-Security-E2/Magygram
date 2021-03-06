import React, { useContext, useState } from "react";
import { modalConstants } from "../constants/ModalConstants";
import { MessageContext } from "../contexts/MessageContext";
import { hasRoles } from "../helpers/auth-header";
import { postService } from "../services/PostService";

const PostInteraction = ({ post, LikePost, DislikePost, UnlikePost, UndislikePost, showCollectionModal, addToDefaultCollection, deleteFromCollections }) => {
	const { dispatch } = useContext(MessageContext);
	const [showAddToCollection, setShowAddToCollection] = useState(false);
	const iconStyle = { fontSize: "35px", margin: "0px" };

	const handleAddToCollection = (postId) => {
		addToDefaultCollection(postId);
		setShowAddToCollection(true);
		setTimeout(() => setShowAddToCollection(false), 5000);
	};

	const handleDeleteFromToCollection = (postId) => {
		deleteFromCollections(postId);
	};

	const handleClickOnWebsite = async () => {
		await postService.clickOnPostCampaignWebsite(post.Id).then(handleOpenWebsite());
	};

	const handleOpenWebsite = () => {
		return new Promise(function () {
			window.open("https://" + post.Website, "_blank");
		});
	};

	return (
		<React.Fragment>
			<div>
				<div hidden={post.ContentType !== "CAMPAIGN"} className="border-bottom" style={{ cursor: "pointer" }}>
					<div>
						<button type="button" className="btn btn-link border-0" onClick={handleClickOnWebsite} target="_blank">
							Visit {post.Website}
						</button>
					</div>
				</div>
				<div className="d-flex flex-row justify-content-center pl-3 pr-3 pt-3 pb-1">
					<div hidden={!showAddToCollection}>
						<button type="button" className="btn btn-outline-secondary btn-icon-text border-0" onClick={() => showCollectionModal(post.Id)}>
							<i className="mdi mdi-plus btn-icon-prepend"></i>Add to collection
						</button>
					</div>
				</div>
				<div className="d-flex flex-row justify-content-between pl-3 pr-3 pt-3 pb-1">
					<ul className="list-inline d-flex flex-row align-items-center m-0">
						<li hidden={post.Liked || post.Disliked || hasRoles(["admin"])} className="list-inline-item">
							<button onClick={() => LikePost(post.Id)} className="btn p-0">
								<i width="1.6em" height="1.6em" fill="currentColor" className="fa fa-thumbs-o-up" style={iconStyle} />
							</button>
						</li>
						<li hidden={post.Liked || post.Disliked || hasRoles(["admin"])} className="list-inline-item">
							<button onClick={() => DislikePost(post.Id)} className="btn p-0">
								<i width="1.6em" height="1.6em" fill="currentColor" className="fa fa-thumbs-o-down" style={iconStyle} />
							</button>
						</li>
						<li hidden={!post.Liked || post.Disliked || hasRoles(["admin"])} className="list-inline-item">
							<button onClick={() => UnlikePost(post.Id)} className="btn p-0">
								<i width="1.6em" height="1.6em" fill="currentColor" className="fa fa-thumbs-up" style={iconStyle} />
							</button>
						</li>
						<li hidden={post.Liked || !post.Disliked || hasRoles(["admin"])} className="list-inline-item">
							<button onClick={() => UndislikePost(post.Id)} className="btn p-0">
								<i width="1.6em" height="1.6em" fill="currentColor" className="fa fa-thumbs-down" style={iconStyle} />
							</button>
						</li>
						<li hidden={hasRoles(["admin"])} className="list-inline-item">
							<button className="btn p-0" onClick={() => dispatch({ type: modalConstants.SHOW_SEND_POST_TO_USER_MODAL, postId: post.Id })}>
								<i width="1.6em" height="1.6em" fill="currentColor" className="fa fa-paper-plane-o" style={iconStyle} />
							</button>
						</li>
					</ul>
					<div>
						<button hidden={hasRoles(["admin"])} className="btn p-0" onClick={() => (!post.Favourites ? handleAddToCollection(post.Id) : handleDeleteFromToCollection(post.Id))}>
							<i width="1.6em" height="1.6em" fill="black" className={post.Favourites ? "fa fa-bookmark" : "fa fa-bookmark-o"} style={iconStyle} />
						</button>
					</div>
				</div>
			</div>
		</React.Fragment>
	);
};

export default PostInteraction;
