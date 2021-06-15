import { useContext } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { PostContext } from "../../contexts/PostContext";
import { authHeader } from "../../helpers/auth-header";
import Axios from "axios";

const OptionsModal = () => {
	const { postState, dispatch } = useContext(PostContext);

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_POST_OPTIONS_MODAL });
	};

	const handleOpenPostEditModal = () => {
		dispatch({ type: modalConstants.SHOW_POST_EDIT_MODAL });
	};

	const handleReportPost = () => {
		
		console.log(postState.editPost.post.id)

		var reportDTO = {
			contentId: postState.editPost.post.id,
			contentType: "POST"
		}

		Axios.post(`/api/report`, reportDTO, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				console.log(res.data);
				dispatch({ type: modalConstants.HIDE_POST_OPTIONS_MODAL });
				alert("You have successfully reported this post!")
			} else {
				console.log("Error");
			}
		})
		.catch((err) => {
			console.log(err)
		});
		
	};

	return (
		<Modal show={postState.postOptions.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Header closeButton>
				<Modal.Title id="contained-modal-title-vcenter">Options</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				<div className="row">
					<button type="button" className="btn btn-link btn-fw text-secondary w-100 border-0" onClick={handleOpenPostEditModal}>
						Edit
					</button>
				</div>
				<div className="row">
					<button type="button" className="btn btn-link btn-fw text-secondary w-100 border-0" onClick={handleReportPost}>
						Report
					</button>
				</div>
				<div className="row">
					<button type="button" className="btn btn-link btn-fw text-danger w-100 border-0">
						Delete?
					</button>
				</div>
			</Modal.Body>
		</Modal>
	);
};

export default OptionsModal;
