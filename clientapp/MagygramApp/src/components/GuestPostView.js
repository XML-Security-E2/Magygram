import React, { useContext } from "react";
import PostHeader from "./PostHeader";
import PostImageSlider from "./PostImageSlider";

const GuestPostView = ({ post }) => {

	return (
		<React.Fragment>
			<div className="d-flex flex-column mt-4 mb-4">
				<div className="card">
					<PostHeader username={post.UserInfo.Username} image={post.UserInfo.ImageURL} />
					<div className="card-body p-0">
						<PostImageSlider media={post.Media} />
					</div>
				</div>
			</div>
		</React.Fragment>
	);
};

export default GuestPostView;
