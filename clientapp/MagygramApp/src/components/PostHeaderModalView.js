import React from "react";

const PostHeaderModalView = ({username, image}) => {
	const imgStyle = {"transform":"scale(1.5)","width":"100%","position":"absolute","left":"0"};

	return (
        <React.Fragment>
                <div className="d-flex flex-row align-items-center">
                    <div className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border border-danger post-profile-photo mr-3">
                        <img src={image} alt="..." style={imgStyle}/>
                    </div>
                    <div className="align-items-top">
                        <div className="font-weight-bold">{username}</div>
                        <div className="font-weight">Ovde staviti lokaciju</div>
                    </div>
                </div>
        </React.Fragment>
	);
};

export default PostHeaderModalView;
