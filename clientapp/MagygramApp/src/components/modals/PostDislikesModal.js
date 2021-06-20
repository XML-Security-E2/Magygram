import { useContext } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { PostContext } from "../../contexts/PostContext";

const PostDislikesModal = () => {
	const imgStyle = { transform: "scale(1.5)", width: "100%", position: "absolute", left: "0" };

	const { postState, dispatch } = useContext(PostContext);

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_POST_DISLIKES_MODAL });
	};

	return (
		<Modal show={postState.postDislikes.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Header closeButton>
				<Modal.Title id="contained-modal-title-vcenter" className="text-center ">
					Likes
				</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				<div className="row">
					<div className="col-md-12">
						<div className="card" style={{ border: "0" }}>
							<div className="card-body">
								{postState.postDislikes.dislikes.map((dislikeUserInfo) => {
									return (
										<div className="card-header">
											<div className="d-flex flex-row align-items-center">
												<div className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border border-danger post-profile-photo mr-3">
													<img src={dislikeUserInfo.ImageURL !== "" ? dislikeUserInfo.ImageURL : "assets/img/profile.jpg"} alt="" style={imgStyle} />
												</div>
												<span className="font-weight-bold">{dislikeUserInfo.Username}</span>
											</div>
										</div>
									);
								})}
							</div>
						</div>
					</div>
				</div>
			</Modal.Body>
		</Modal>
	);
};

export default PostDislikesModal;
