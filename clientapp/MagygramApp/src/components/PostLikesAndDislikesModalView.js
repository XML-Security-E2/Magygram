import React from "react";

const PostLikesAndDislikesModalView = ({likes,dislikes, showLikedByModal, showDislikesModal}) => {

	return (
        <React.Fragment>
            <strong className="d-block pl-2 ">
                    <label hidden={likes===0} onClick={()=> showLikedByModal()}>{likes} likes    </label>
                    <label hidden={dislikes===0} onClick={()=> showDislikesModal()} className="ml-2">{dislikes} dislikes</label> 
            </strong>
        </React.Fragment>
	);
};

export default PostLikesAndDislikesModalView;
