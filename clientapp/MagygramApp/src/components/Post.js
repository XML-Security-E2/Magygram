import React from "react";
import PostComments from "./PostComments";
import PostHeader from "./PostHeader";
import PostImageSlider from "./PostImageSlider";
import PostInformation from "./PostInformation";
import PostInteraction from "./PostInteraction";

const Post = ({post}) => {
	return (
        <React.Fragment>
            <div className="d-flex flex-column mt-4 mb-4">
                <div className="card">
                    <PostHeader username={post.UserInfo.Username} image={post.UserInfo.ImageURL}/>
                    <div className="card-body p-0">
                        <PostImageSlider media={post.Media}/>
                        <PostInteraction/>
                        <div className="pl-3 pr-3 pb-2">
                        <PostInformation username={post.UserInfo.Username} likes={post.LikedBy.length} description={post.Description}/>
                        </div>
                        <PostComments comments={post.Comments}/>
                    </div>
                </div>
            </div>
        </React.Fragment>
	);
};

export default Post;
