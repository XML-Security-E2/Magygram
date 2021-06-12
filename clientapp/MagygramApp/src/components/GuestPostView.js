import React from "react";
import PostHeader from "./PostHeader";
import PostImageSlider from "./PostImageSlider";

const GuestPostView = ({ post }) => {
	return (
		<React.Fragment>
			<div className="d-flex flex-column mt-4 mb-4">
				<div className="card">
					<PostHeader username={post.UserInfo.Username} image={post.UserInfo.ImageURL} id={post.UserInfo.Id} />
					<div className="card-body p-0">
						<PostImageSlider media={post.Media} />
					</div>
					<div className="pl-3 pr-3 pb-2 pt-3">
						<strong className="d-block">{post.UserInfo.Username}</strong>
						<p className="d-block mb-1">{post.Description}</p>
					</div>
				</div>
			</div>
		</React.Fragment>
	);
};

export default GuestPostView;
