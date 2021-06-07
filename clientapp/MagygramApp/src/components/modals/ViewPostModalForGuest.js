import { useContext } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { PostContext } from "../../contexts/PostContext";
import  PostHeader  from "../../components/PostHeader"
import  PostImageSlider  from "../../components/PostImageSlider"

const ViewPostModalForGuest = () => {
	const { postState, dispatch } = useContext(PostContext);

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_VIEW_POST_FOR_GUEST_MODAL, post: postState.viewPostModalForGuest.post });
	};

	return (
		<Modal size="md" show={postState.viewPostModalForGuest.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Body>
                    <div className="card">
                        <PostHeader username={postState.viewPostModalForGuest.post.UserInfo.Username} image={postState.viewPostModalForGuest.post.UserInfo.ImageURL} />
                        <div className="card-body p-0">
                            <PostImageSlider media={postState.viewPostModalForGuest.post.Media} />
                        </div>
                        <div className="pl-3 pr-3 pb-2 pt-3">
                            <strong className="d-block">{postState.viewPostModalForGuest.post.UserInfo.Username}</strong>
                            <p className="d-block mb-1">{postState.viewPostModalForGuest.post.Description}</p>
                        </div>
                    </div>
			</Modal.Body>
		</Modal>
	);
};

export default ViewPostModalForGuest;
