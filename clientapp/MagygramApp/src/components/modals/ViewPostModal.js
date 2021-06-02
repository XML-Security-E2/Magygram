import { useContext } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { PostContext } from "../../contexts/PostContext";
import PostHeaderModalView from "../PostHeaderModalView";
import PostImageSliderModalView from "../PostImageSliderModalView";
import PostCommentsModalView from "../PostCommentsModalView"
import PostInteractionModalView from "../PostInteractionModalView"
import { postService } from "../../services/PostService";
import PostCommentInputModalView from "../PostCommentInputModalView"
import PostLikesAndDislikesModalView from "../PostLikesAndDislikesModalView"

const ViewPostModal = () => {
	const { postState, dispatch } = useContext(PostContext);
	const style = {"width":"450px"};

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

    const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_VIEW_POST_MODAL, post:postState.viewPostModal.post });
	}

    const postComment = (comment) => {
        if(comment.length>=1){
            let postDTO = {
                PostId : postState.viewPostModal.post.Id,
                Content: comment
            };

            postService.commentPost(postDTO, dispatch)
        }
    }

    const showLikedByModal = () =>{
        dispatch({ type: modalConstants.SHOW_POST_LIKED_BY_DETAILS, LikedBy: postState.viewPostModal.post.LikedBy })
    }

    const showDislikesModal = () => {
        dispatch({ type: modalConstants.SHOW_POST_DISLIKES_MODAL, Dislikes: postState.viewPostModal.post.DislikedBy })

    }

	return (
		<Modal size="xl" show={postState.viewPostModal.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>

			<Modal.Body>
            <div className="d-flex flex-row align-items-top">

                <PostImageSliderModalView media={postState.viewPostModal.post.Media}/>
                <div className="p-2" style={style}>
                    <div className="align-top" style={style}> 
                        <PostHeaderModalView 
                            username={postState.viewPostModal.post.UserInfo.Username} 
                            image={postState.viewPostModal.post.UserInfo.ImageURL}
                        />
                        <hr></hr>
                        <PostCommentsModalView 
                            imageUrl={postState.viewPostModal.post.UserInfo.ImageURL}
                            username={postState.viewPostModal.post.UserInfo.Username}
                            description={postState.viewPostModal.post.Description}
                            comments={postState.viewPostModal.post.Comments}
                        />
                        <hr></hr>
                        <PostInteractionModalView 
                            post={postState.viewPostModal.post} 
                            LikePost={LikePost} 
                            DislikePost={DislikePost} 
                            UnlikePost={UnlikePost} 
                            UndislikePost={UndislikePost} />
                        <PostLikesAndDislikesModalView 
                            likes={postState.viewPostModal.post.LikedBy.length} 
                            dislikes={postState.viewPostModal.post.DislikedBy.length}
                            showLikedByModal={showLikedByModal}
                            showDislikesModal={showDislikesModal}
                        />
                        <PostCommentInputModalView postComment={postComment}/>
                    </div>
                </div>
            </div>

                                        
			</Modal.Body>
		</Modal>
	);
};

export default ViewPostModal;
