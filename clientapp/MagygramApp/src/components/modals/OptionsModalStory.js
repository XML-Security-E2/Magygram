import { useContext, useEffect } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { StoryContext } from "../../contexts/StoryContext";
import Axios from "axios";
import { authHeader } from "../../helpers/auth-header";

const OptionsModalStory = () => {
	const { storyState, dispatch } = useContext(StoryContext);

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_STORY_OPTIONS_MODAL });
	};

	const handleReportModal = () => {

		let reportDTO = {
			contentId: storyState.storySliderModal.userId,
			contentType: "STORY",
		};

		Axios.post(`/api/report`, reportDTO , { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
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
		<Modal show={storyState.postOptions.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Header closeButton>
				<Modal.Title id="contained-modal-title-vcenter">Options</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				<div className="row">
					<button hidden={localStorage.getItem("userId") === null}  type="button" className="btn btn-link btn-fw text-secondary w-100 border-0" onClick={handleReportModal}>
						Report
					</button>
				</div>
			</Modal.Body>
		</Modal>
	);
};

export default OptionsModalStory;
