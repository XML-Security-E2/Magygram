import React, { useState } from "react";

const MediaInputField = () => {
	const [image, setImage] = useState("");

	const [showedImages, setShowedImages] = useState([]);

	const imgRef = React.createRef();

	const onImageChange = (e) => {
		setImage(e.target.files[0]);

		if (e.target.files && e.target.files[0]) {
			let img = e.target.files[0];
			let imgCpy = [...showedImages];
			imgCpy.push(URL.createObjectURL(img));
			setShowedImages(imgCpy);
		}
	};

	const handleImageDeselect = () => {
		//dispatch({ type: objectConstants.OBJECT_IMAGE_DESELECTED });
	};

	return (
		<React.Fragment>
			<div className="row">
				<div className="col-3 box mb-4">
					<input type="file" ref={imgRef} style={{ display: "none" }} name="image" accept="image/png, image/jpeg video/mp4" onChange={onImageChange} />

					<button type="button" disabled={showedImages.length > 5} onClick={() => imgRef.current.click()} className="btn btn-outline-secondary rounded-lg btn-icon w-100 h-100">
						<i className="mdi mdi-plus w-50 h-50"></i>
					</button>
				</div>
				{showedImages.map((showedImage) => {
					return (
						<div className="col-3 box mb-4 container-img">
							<img src={showedImage} className="img-fluid rounded-lg w-100 h-100" alt="" />
							<div className="overlay-img rounded">
								<button
									hidden={false}
									className="btn btn-outline-dark  btn-icon icon-img btn-rounded"
									data-toggle="tooltip"
									title="Delete image"
									onClick={() => imgRef.current.click()}
								>
									<i class="mdi mdi-close text-danger"></i>
								</button>
							</div>
						</div>
					);
				})}
			</div>
		</React.Fragment>
	);
};

export default MediaInputField;
