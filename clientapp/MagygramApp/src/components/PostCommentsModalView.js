import React from "react";

const PostCommentsModalView = ({ imageUrl, username, description, comments, handleRedirect }) => {
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
						<div>
							<div className="d-flex flex-row align-items-center pt-3">
								<div className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border border-danger comments-profile-photo mr-3">
									<img src={comment.CreatedBy.ImageURL} alt="..." style={imgStyle} />
								</div>
								<span className="description_label">
									<b className="pr-3">{comment.CreatedBy.Username}</b>
									{comment.Content}
									{comment.Tags.map((tag) => {
									return (
										<button type="button" className="btn btn-light m-1" onClick={() => handleRedirect(tag.Id)}>
											{tag.Username}
										</button>
									);			
								})}
								</span>
							</div>
						</div>
					);
				})}
			</div>
		</React.Fragment>
	);
};

export default PostCommentsModalView;
