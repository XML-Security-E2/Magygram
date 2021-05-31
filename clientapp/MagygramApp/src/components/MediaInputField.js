import React from "react";
import { v4 as uuidv4 } from "uuid";

const MediaInputField = ({ showedMedia, setShowedMedia }) => {
	const imgRef = React.createRef();
	const videoRef = React.createRef();

	const onImageChange = (e) => {
		if (e.target.files && e.target.files[0]) {
			let img = e.target.files[0];
			let mediaCpy = [...showedMedia];
			let filename = URL.createObjectURL(img);
			let extension = e.target.files[0].name.split(".").pop();

			console.log(extension);
			if (extension === "mp4") {
				mediaCpy.push({ src: filename, type: "video", id: uuidv4(), content: e.target.files[0] });
				console.log(mediaCpy);
			} else mediaCpy.push({ src: filename, type: "image", id: uuidv4(), content: e.target.files[0] });

			setShowedMedia(mediaCpy);
		}
	};

	const handleImageDeselect = (id) => {
		let mediaCpy = showedMedia.filter((media) => media.id !== id);
		setShowedMedia(mediaCpy);
		//dispatch({ type: objectConstants.OBJECT_IMAGE_DESELECTED });
	};

	return (
		<React.Fragment>
			<h5>Choose images</h5>
			<div className="row p-3 ">
				<div className="col-2 box mb-4 ">
					<input type="file" ref={imgRef} style={{ display: "none" }} name="image" accept="image/png, image/jpeg, video/mp4" onChange={onImageChange} />

					<button type="button" disabled={showedMedia.length > 5} onClick={() => imgRef.current.click()} className="btn btn-outline-secondary-whitebckg rounded-lg btn-icon w-100 h-100">
						<i className="mdi mdi-plus w-50 h-50"></i>
					</button>
				</div>
				{showedMedia.map((showedMedia) => {
					return (
						<div className="col-2 box mb-4 container-img" key={showedMedia.id}>
							{showedMedia.type === "image" ? (
								<img src={showedMedia.src} className="img-fluid rounded-lg w-100 " alt="" />
							) : (
								<video ref={videoRef} controls className="img-fluid rounded-lg w-100">
									<source src={showedMedia.src} type="video/mp4" />
								</video>
							)}
							<button
								hidden={false}
								className="btn btn-outline-dark btn-icon icon-img btn-rounded border-0"
								data-toggle="tooltip"
								title="Delete image"
								style={{ zIndex: "101" }}
								onClick={() => handleImageDeselect(showedMedia.id)}
							>
								<i className="mdi mdi-close text-danger"></i>
							</button>
						</div>
					);
				})}
			</div>
		</React.Fragment>
	);
};

export default MediaInputField;
