import React, { useState } from "react";
import { Modal } from "react-bootstrap";

const CreateStoryModal = () => {
	const [showedMedia, setShowedMedia] = useState("");
	const imgRef = React.createRef();

	const onImageChange = (e) => {
		if (e.target.files && e.target.files[0]) {
			let img = e.target.files[0];
			let filename = URL.createObjectURL(img);
			let extension = e.target.files[0].name.split(".").pop();

			if (extension === "mp4") {
				setShowedMedia({ src: filename, type: "video", content: e.target.files[0] });
			} else setShowedMedia({ src: filename, type: "image", content: e.target.files[0] });
		}
	};

	const handleImageDeselect = () => {
		setShowedMedia("");
	};

	return (
		<Modal show={true} className="story_modal" aria-labelledby="contained-modal-title-vcenter">
			<Modal.Body>
				<input type="file" ref={imgRef} style={{ display: "none" }} name="image" accept="image/png, image/jpeg, video/mp4" onChange={onImageChange} />
				<div style={{ opacity: showedMedia !== "" ? 1 : 0 }} className="d-flex flex-row-reverse">
					<button
						hidden={false}
						disabled={showedMedia === ""}
						className="btn btn-outline-dark btn-icon btn-rounded border-0"
						data-toggle="tooltip"
						title="Delete image"
						onClick={handleImageDeselect}
					>
						<i className="mdi mdi-close text-danger"></i>
					</button>
				</div>
				<div style={{ height: "85vh", cursor: "pointer" }} className="d-flex align-items-center justify-content-center" centered onClick={() => imgRef.current.click()}>
					{showedMedia === "" ? (
						<div>
							<div className="row d-flex align-items-center justify-content-center">
								<button type="button" className="btn btn-outline-secondary rounded-lg btn-icon">
									<i className="mdi mdi-plus w-50 h-50"></i>
								</button>
							</div>
							<div className="row mt-5" style={{ color: "white", fontSize: "1em" }}>
								Click to add story
							</div>
						</div>
					) : (
						<div hidden={showedMedia === ""} style={{ height: "90vh" }} className="d-flex align-items-center justify-content-center">
							<img src={showedMedia.src} className="img-fluid rounded-lg w-100 " alt="" />
						</div>
					)}
				</div>
			</Modal.Body>
		</Modal>
	);
};

export default CreateStoryModal;
