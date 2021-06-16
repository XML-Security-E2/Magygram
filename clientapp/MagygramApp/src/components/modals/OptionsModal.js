import { useContext, useEffect } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { PostContext } from "../../contexts/PostContext";
import { UserContext } from "../../contexts/UserContext";
import Axios from "axios";
import { authHeader } from "../../helpers/auth-header";

const OptionsModal = () => {
	const { postState, dispatch } = useContext(PostContext);

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_POST_OPTIONS_MODAL });
	};

	useEffect(() => {
	
	});

	const handleOpenPostEditModal = () => {
		dispatch({ type: modalConstants.SHOW_POST_EDIT_MODAL });
	};

	const handleReportModal = () => {

		let reportDTO = {
			contentId: postState.editPost.post.id,
			contentType: "POST",
		};

		Axios.post(`/api/report`, reportDTO , { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log("blaa");
			if (res.status === 200) {
				alert("You have reported this post successfully")
				dispatch({ type: modalConstants.HIDE_POST_OPTIONS_MODAL });
			} else {
				console.log("err")
			}
		})
		.catch((err) => {
			console.log("err")
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
					<button hidden={localStorage.getItem("userId") === null} type="button" className="btn btn-link btn-fw text-secondary w-100 border-0" onClick={handleReportModal}>
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
