import React from "react";
import { Modal } from "react-bootstrap";

const VideoViewerModal = ({ handleModalClose, show, video }) => {
	return (
		<Modal show={show} className="video_modal" aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Body>
				<video controls className="img-fluid rounded-lg w-100">
					<source src={video} type="video/mp4" />
				</video>
			</Modal.Body>
		</Modal>
	);
};

export default VideoViewerModal;
