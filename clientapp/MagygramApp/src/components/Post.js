import React, {useContext} from "react";
import PostComments from "./PostComments";
import PostHeader from "./PostHeader";
import PostImageSlider from "./PostImageSlider";
import PostInformation from "./PostInformation";
import PostInteraction from "./PostInteraction";
import { postService } from "../services/PostService";
import { PostContext } from "../contexts/PostContext";
import PostLikesModal from "./modals/PostLikesModal";
import { modalConstants } from "../constants/ModalConstants";
import PostDislikesModal from "./modals/PostDislikesModal";

const Post = ({post}) => {
	const { dispatch } = useContext(PostContext);

    const LikePost = (postId) => {
        postService.likePost(postId, dispatch)
    }

    const UnlikePost = (postId) => {
        postService.unlikePost(postId, dispatch)
    }

    const DislikePost = (postId) => {
        postService.dislikePost(postId, dispatch)
    }

    const UndislikePost = (postId) => {
        postService.undislikePost(postId, dispatch)
    }

    const showLikedByModal = () =>{
        dispatch({ type: modalConstants.SHOW_POST_LIKED_BY_DETAILS, LikedBy: post.LikedBy })
    }

    const showDislikesModal = () => {
        dispatch({ type: modalConstants.SHOW_POST_DISLIKES_MODAL, Dislikes: post.DislikedBy })

    }

	return (
        <React.Fragment>
            <div className="d-flex flex-column mt-4 mb-4">
                <div className="card">
                    <PostHeader username={post.UserInfo.Username} image={post.UserInfo.ImageURL}/>
                    <div className="card-body p-0">
                        <PostImageSlider media={post.Media}/>
                        <PostInteraction post={post} LikePost={LikePost} DislikePost={DislikePost} UnlikePost={UnlikePost} UndislikePost={UndislikePost}/>
                        <div className="pl-3 pr-3 pb-2">
                        <PostInformation username={post.UserInfo.Username} likes={post.LikedBy.length} dislikes={post.DislikedBy.length} description={post.Description} showLikedByModal={showLikedByModal} showDislikesModal={showDislikesModal}/>
                        </div>
                        <PostComments comments={post.Comments}/>
                        <PostLikesModal/>
                        <PostDislikesModal/>
                    </div>
                    
                </div>
            </div>
        </React.Fragment>
	);
};

export default Post;
