import React, {useState} from "react";
import { hasRoles } from "../helpers/auth-header";

const PostCommentInputModalView = ({postComment}) => {
    const [comment, setComment] = useState("");

	return (
        <React.Fragment>
            <div className="position-relative comment-box pt-2">
                <form hidden={hasRoles(["admin"])}>
                    <input value={comment} onChange={(e) => setComment(e.target.value)} minlength="1" className="w-100 border-0 p-3 input-post" placeholder="Add a comment..."/>
                    <button onClick={() => { setComment(""); postComment(comment)}} className="btn btn-primary position-absolute btn-ig">Post</button>
                </form>
            </div>
        </React.Fragment>
	);
};

export default PostCommentInputModalView;
