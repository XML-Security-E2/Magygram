import React from "react";

const PostInformation = ({likes, username, description}) => {
	return (
        <React.Fragment>
            <strong className="d-block">{likes} likes</strong>
            <strong className="d-block">{username}</strong>
            <p className="d-block mb-1">{description}</p>
        </React.Fragment>
	);
};

export default PostInformation;
