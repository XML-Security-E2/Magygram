import React from "react";

const PostInformation = ({likes, dislikes, username, description, showLikedByModal, showDislikesModal}) => {
	return (
        <React.Fragment>
            <strong className="d-block">
                <label hidden={likes===0} onClick={()=> showLikedByModal()}>{likes} likes    </label>
                <label hidden={dislikes===0} onClick={()=> showDislikesModal()} className="ml-2">{dislikes} dislikes</label> 
            </strong>
            <strong className="d-block">{username}</strong>
            <p className="d-block mb-1">{description}</p>
        </React.Fragment>
	);
};

export default PostInformation;
