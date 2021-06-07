import React from "react";
import { useContext } from "react";
import { PostContext } from "../contexts/PostContext";
import { postService } from "../services/PostService";
import CollectionButton from "./CollectionButton";

const UserProfileSavedPosts = () => {
	const { postState, dispatch } = useContext(PostContext);

	const handleShowCollectionPosts = async (Collection) => {
		await postService.findAllPostsFromCollection(Collection, dispatch);
	};

	const getPostDetailsHandler = async (postId) => {
		await postService.findPostById(postId, dispatch);
	};

	return (
		<React.Fragment>
			<div className="row" hidden={!postState.userProfileContent.showCollections}>
				{postState.userProfileContent.collections.length > 0 ? (
					Object.keys(postState.userProfileContent.collections).map((collection) => {
						return (
							<div className="col-4" align="center" onClick={() => handleShowCollectionPosts(collection)}>
								<CollectionButton collectionName={collection} media={postState.userProfileContent.collections[collection]} />
							</div>
						);
					})
				) : (
					<div className="col-12 mt-5 d-flex justify-content-center text-secondary">
						<h3>All collections are empty</h3>
					</div>
				)}
			</div>
			<h4 className="text-secondary" hidden={!postState.userProfileContent.showCollectionPosts}>
				{postState.userProfileContent.selectedCollectionName}
			</h4>

			<div className="row mt-4" hidden={!postState.userProfileContent.showCollectionPosts}>
				{postState.userProfileContent.collectionPosts !== null &&
					postState.userProfileContent.collectionPosts.map((post) => {
						return (
							<div className="col-4 box" style={{ cursor: "pointer" }} onClick={() => getPostDetailsHandler(post.id)}>
								{post.media.mediaType === "IMAGE" ? (
									<img src={post.media.url} className="img-fluid box-coll rounded-lg w-100 " alt="" style={{ objectFit: "cover" }} />
								) : (
									<video className="img-fluid box-coll rounded-lg w-100" style={{ objectFit: "cover" }}>
										<source src={post.media.url} type="video/mp4" />
									</video>
								)}
							</div>
						);
					})}
			</div>
		</React.Fragment>
	);
};

export default UserProfileSavedPosts;
