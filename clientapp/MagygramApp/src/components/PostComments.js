import React from "react";

const PostComments = ({comments}) => {

	return (
        <React.Fragment>
            <div className="pl-3 pr-3 pb-2">
                <button className="btn p-0">
                    <span className="text-muted">View all {comments.length} comments</span>
                </button>
                <div>
                    <div>
                        <strong className="d-block">a.7.m3ff</strong>
                        <span>❤️💓💓💓💓💓</span>
                    </div>
                    <div>
                        <strong className="d-block">adri_rez77</strong>
                        <span>Hi</span>
                    </div>
                </div>
                <small className="text-muted">4 HOURS AGO</small>
            </div>
            <div className="position-relative comment-box">
                <form>
                    <input className="w-100 border-0 p-3 input-post" placeholder="Add a comment..."/>
                    <button className="btn btn-primary position-absolute btn-ig">Post</button>
                </form>
            </div>
        </React.Fragment>
	);
};

export default PostComments;
