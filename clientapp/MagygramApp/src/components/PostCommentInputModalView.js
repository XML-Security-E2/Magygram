import React, {useState} from "react";
import { hasRoles } from "../helpers/auth-header";
import AsyncSelect from "react-select/async";
import { searchService } from "../services/SearchService";

const PostCommentInputModalView = ({postComment}) => {
    const [comment, setComment] = useState("");
    const [search, setSearch] = useState("");
    const [tags, setTags] = useState([]);
    const [hide, setHide] = useState(true);

    const loadOptions = (value, callback) => {
		setTimeout(() => {
			searchService.userSearchTags(value, callback);
		}, 1000);
	};

	const onInputChange = (inputValue, { action }) => {
		switch (action) {
			case "set-value":
				return;
			case "menu-close":
				setSearch("");
				return;
			case "input-change":
				setSearch(inputValue);
				return;
			default:
				return;
		}
	};

	const onChange = (option) => {
        if (option.length === 0)
            setHide(true);
        let tagCollection = [];
        option.forEach((tag) => {
			tagCollection.push({username: tag.value, id: tag.id});
		});
		setTags(tagCollection);
		return false;
	};

    const updateComment = (value) => {
        if (value.slice(-1) == '@' || tags.length > 0)
            setHide(false);
        else 
            setHide(true);
            
        setComment(value);
    }

	return (
        <React.Fragment>
            <div className="position-relative comment-box pt-2">
                <form hidden={hasRoles(["admin"])}>
                    <input value={comment} onChange={(e) => updateComment(e.target.value)} minlength="1" className="w-100 border-0 p-3 input-post" placeholder="Add a comment..."/>
                    <div hidden={hide}>
                        <AsyncSelect isMulti defaultOptions loadOptions={loadOptions} onInputChange={onInputChange} onChange={onChange} placeholder="users" inputValue={search} />
                    </div>
                    <button onClick={() => { setComment(""); postComment(comment, tags)}} className="btn btn-primary position-absolute btn-ig">Post</button>
                </form>
            </div>
        </React.Fragment>
	);
};

export default PostCommentInputModalView;
