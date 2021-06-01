import React from "react";

const PostInformation = ({likes, dislikes, username, description}) => {
	return (
        <React.Fragment>
            <strong className="d-block">{likes} likes  and {dislikes} dislikes</strong>
            <strong className="d-block">{username}</strong>
            <p className="d-block mb-1">{description}</p>
        </React.Fragment>
	);
};

export default PostInformation;
