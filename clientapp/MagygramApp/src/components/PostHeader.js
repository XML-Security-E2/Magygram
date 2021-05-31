import React from "react";

const PostHeader = ({username, image}) => {
	const imgStyle = {"transform":"scale(1.5)","width":"100%","position":"absolute","left":"0"};

	return (
        <React.Fragment>
            <div className="card-header p-3" >
                <div className="d-flex flex-row align-items-center">
                    <div className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border border-danger post-profile-photo mr-3">
                        <img src={image} alt="..." style={imgStyle}/>
                    </div>
                    <span className="font-weight-bold">{username}</span>
                </div>
            </div>
        </React.Fragment>
	);
};

export default PostHeader;
