import { useContext, useEffect } from "react";
import { Button, Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { postConstants } from "../../constants/PostConstants";
import { PostContext } from "../../contexts/PostContext";
import { postService } from "../../services/PostService";
import AddCollectionButton from "../AddCollectionButton";
import CollectionButton from "../CollectionButton";
import FailureAlert from "../FailureAlert";
import SuccessAlert from "../SuccessAlert";

const AddPostToFavouritesModal = () => {
	const { postState, dispatch } = useContext(PostContext);

	const handleModalClose = () => {
		dispatch({ type: modalConstants.CLOSE_ADD_TO_COLLECTION_MODAL });
	};

	const handleSaveToCollection = (collectionName) => {
		let collectionDTO = {
			postId: postState.addToFavouritesModal.selectedPostId,
			collectionName,
		};

		postService.addPostToCollection(collectionDTO, dispatch);
	};

	const handleCreateCollection = (collectionName) => {
		postService.createCollection(collectionName, dispatch);
	};

	useEffect(() => {
		const getCollections = async () => {
			await postService.findAllUsersCollections(dispatch);
		};
		getCollections();
	}, [postState.addToFavouritesModal.renderCollectionSwitch, dispatch]);

	return (
		<Modal show={postState.addToFavouritesModal.showModal} size="lg" aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Header closeButton>
				<Modal.Title id="contained-modal-title-vcenter">Select collection</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				<SuccessAlert
					hidden={!postState.userCollections.showSuccessMessage}
					header="Success"
					message={postState.userCollections.successMessage}
					handleCloseAlert={() => dispatch({ type: postConstants.ADD_POST_TO_COLLECTION_REQUEST })}
				/>
				<FailureAlert
					hidden={!postState.userCollections.showError}
					header="Error"
					message={postState.userCollections.errorMessage}
					handleCloseAlert={() => dispatch({ type: postConstants.ADD_POST_TO_COLLECTION_REQUEST })}
				/>
				<div className="row">
					{Object.keys(postState.userCollections.collections).map((collection) => {
						return (
							<div className="col-3 " align="center" onClick={() => handleSaveToCollection(collection)}>
								<CollectionButton collectionName={collection} media={postState.userCollections.collections[collection]} />
							</div>
						);
					})}
					<div className="col-3 " align="center">
						<AddCollectionButton createCollection={handleCreateCollection} />
					</div>
				</div>
			</Modal.Body>
			<Modal.Footer>
				<Button onClick={handleModalClose}>Close</Button>
			</Modal.Footer>
		</Modal>
	);
};

export default AddPostToFavouritesModal;
