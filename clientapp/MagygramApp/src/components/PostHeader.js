import React from "react";

const PostHeader = ({ username, image, openOptionsModal }) => {
	const imgStyle = { transform: "scale(1.5)", width: "100%", position: "absolute", left: "0" };

	return (
		<React.Fragment>
			<div className="card-header p-3">
				<div className="d-flex flex-row justify-content-between">
					<ul className="list-inline d-flex flex-row align-items-center m-0">
						<li className="list-inline-item">
							<div className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border border-danger post-profile-photo ">
								<img src={image} alt="..." style={imgStyle} />
							</div>
						</li>
						<li className="list-inline-item">
							<span className="font-weight-bold">{username}</span>
						</li>
					</ul>

				</div>
			</div>
		</React.Fragment>
	);
};

export default PostHeader;
