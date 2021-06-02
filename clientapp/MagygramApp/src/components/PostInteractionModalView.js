import React from "react";

const PostInteractionModalView = ({post, LikePost, DislikePost, UnlikePost, UndislikePost}) => {
    const iconStyle = { fontSize: "30px", margin: "0px" };


	return (
        <React.Fragment>
             <div className="d-flex flex-row justify-content-between pb-2 pl-2">
                <ul className="list-inline d-flex flex-row align-items-center m-0">
                    <li hidden={post.Liked || post.Disliked} className="list-inline-item">
                        <button onClick={()=> LikePost(post.Id)} className="btn p-0">
                                <i width="1.6em" height="1.6em" fill="currentColor" className="fa fa-thumbs-o-up" style={iconStyle} />
                        </button>
                    </li>
                    <li hidden={post.Liked || post.Disliked} className="list-inline-item">
                        <button onClick={()=> DislikePost(post.Id)} className="btn p-0">
                                <i width="1.6em" height="1.6em" fill="currentColor" className="fa fa-thumbs-o-down" style={iconStyle} />
                        </button>
                    </li>
                    <li hidden={!post.Liked || post.Disliked} className="list-inline-item">
                        <button onClick={()=> UnlikePost(post.Id)} className="btn p-0">
                                <i width="1.6em" height="1.6em" fill="currentColor" className="fa fa-thumbs-up" style={iconStyle} />
                        </button>
                    </li>
                    <li hidden={post.Liked || !post.Disliked} className="list-inline-item">
                        <button onClick={()=> UndislikePost(post.Id)} className="btn p-0">
                                <i width="1.6em" height="1.6em" fill="currentColor" className="fa fa-thumbs-down" style={iconStyle} />
                        </button>
                    </li>
                </ul>
                <div>
                    <button class="btn p-0">
                        <i width="1.6em" height="1.6em" fill="currentColor" className="la la-save" style={iconStyle} />
                    </button>
                </div>
            </div>
        </React.Fragment>
	);
};

export default PostInteractionModalView;
