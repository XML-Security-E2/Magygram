import React from "react";

const PostCommentsModalView = ({ imageUrl, username, description, comments }) => {
	const imgStyle = { transform: "scale(1.5)", width: "100%", position: "absolute", left: "0" };

	return (
		<React.Fragment>
			<div className="container messagessss">
				<div className="d-flex flex-row align-items-center pt-3">
					<div className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border border-danger comments-profile-photo mr-3">
						<img src={imageUrl === "" ? "assets/img/profile.jpg" : imageUrl} alt="" style={imgStyle} />
					</div>
					<span className="description_label">
						<b className="pr-3">{username}</b>
						{description}
					</span>
				</div>
				{comments.map((comment) => {
					return (
						<div className="d-flex flex-row align-items-center pt-3">
							<div className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border border-danger comments-profile-photo mr-3">
								<img src={comment.CreatedBy.ImageURL} alt="..." style={imgStyle} />
							</div>
							<span className="description_label">
								<b className="pr-3">{comment.CreatedBy.Username}</b>
								{comment.Content}
							</span>
						</div>
					);
				})}
			</div>
		</React.Fragment>
	);
};

export default PostCommentsModalView;
