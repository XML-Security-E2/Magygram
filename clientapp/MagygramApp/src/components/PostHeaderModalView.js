import React from "react";

const PostHeaderModalView = ({ username, image, location, openOptionsModal }) => {
	const imgStyle = { transform: "scale(1.5)", width: "100%", position: "absolute", left: "0" };

	return (
		<React.Fragment>
			<div className="d-flex flex-row align-items-center justify-content-between">
				<div className="list-inline d-flex flex-row align-items-center m-0">
					<div className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border border-danger post-profile-photo mr-3 ml-2">
						<img src={image} alt="..." style={imgStyle} />
					</div>
					<div className="align-items-top">
						<div className="font-weight-bold">{username}</div>
						<div className="font-weight">{location}</div>
					</div>
				</div>

				<div className="d-flex justify-content-end">
					<button className="btn p-0 mr-3" onClick={openOptionsModal}>
						<i className="fa fa-ellipsis-h" aria-hidden="true"></i>
					</button>
				</div>
			</div>
		</React.Fragment>
	);
};

export default PostHeaderModalView;
