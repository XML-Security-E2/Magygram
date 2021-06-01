import { useContext, useEffect } from "react";
import { Button, Modal } from "react-bootstrap";
import { PostContext } from "../../contexts/PostContext";
import { postService } from "../../services/PostService";
import CollectionButton from "../CollectionButton";

const AddPostToFavouritesModal = () => {
	const { postState, dispatch } = useContext(PostContext);

	useEffect(() => {
		const getCollections = async () => {
			await postService.findAllUsersCollections(dispatch);
		};
		getCollections();
	}, []);

	return (
		<Modal show={true} size="lg" aria-labelledby="contained-modal-title-vcenter" centered>
			<Modal.Header closeButton>
				<Modal.Title id="contained-modal-title-vcenter">Favourites</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				<div className="row ">
					{postState.userCollections.collections.map((collection) => {
						return (
							<div className="col-3 " align="center">
								<CollectionButton />
							</div>
						);
					})}
				</div>
			</Modal.Body>
			<Modal.Footer>
				<Button>Close</Button>
			</Modal.Footer>
		</Modal>
	);
};

export default AddPostToFavouritesModal;
