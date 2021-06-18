import React, {useState} from "react";
import { hasRoles } from "../helpers/auth-header";

const PostComments = ({comments , postComment ,viewAllComments}) => {
	const [comment, setComment] = useState("");


	return (
        <React.Fragment>
            <div className="pl-3 pr-3 pb-2">
                <button hidden={comments.length<3} onClick={() => viewAllComments()} className="btn p-0">
                    <span className="text-muted">View all {comments.length} comments</span>
                </button>

                {comments.map((comment, i, arr) => {
                    if (arr.length - 1 === i) {
                        return (
                            <div>
                                <div>
                                    <strong className="d-block">{comment.CreatedBy.Username}</strong>
                                    <span>{comment.Content}</span>
                                </div>
                            </div>
                            )
                    } else if(arr.length - 2 === i){
                        return (
                            <div>
                                <div>
                                    <strong className="d-block">{comment.CreatedBy.Username}</strong>
                                    <span>{comment.Content}</span>
                                </div>
                            </div>)

                    }
                        return ("")
                    })}
                
                
            </div>
            <div className="position-relative comment-box">
                <form hidden={hasRoles(["admin"])}>
                    <input value={comment} onChange={(e) => setComment(e.target.value)} minlength="1" className="w-100 border-0 p-3 input-post" placeholder="Add a comment..."/>
                    <button onClick={() => { setComment(""); postComment(comment)}} className="btn btn-primary position-absolute btn-ig">Post</button>
                </form>
            </div>
        </React.Fragment>
	);
};

export default PostComments;
