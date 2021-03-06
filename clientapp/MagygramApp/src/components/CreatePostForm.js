import React, { useContext, useState } from "react";
import MediaInputField from "./MediaInputField";
import { v4 as uuidv4 } from "uuid";
import TagsListInput from "./TagsListInput";
import { PostContext } from "../contexts/PostContext";
import { postService } from "../services/PostService";
import { postConstants } from "../constants/PostConstants";
import SuccessAlert from "./SuccessAlert";
import FailureAlert from "./FailureAlert";
import AsyncSelect from "react-select/async";
import { searchService } from "../services/SearchService";

const CreatePostForm = () => {
	const { postState, dispatch } = useContext(PostContext);

	const [location, setLocation] = useState("");
	const [description, setDescription] = useState("");
	const [showedMedia, setShowedMedia] = useState([]);
	const [tags, setTags] = useState([]);
	const [tagInput, setTagInput] = useState("");
	const [search, setSearch] = useState("");

	const handleAddTag = () => {
		let prom = [...tags];
		//real user
		prom.push({ Id: uuidv4(), EntityDTO: { Name: tagInput } });

		setTags(prom);
		setTagInput("");
	};

	const handleTagDelete = (tagId) => {
		console.log(location);
		let prom = tags.filter((tag) => tag.Id !== tagId);
		setTags(prom);
	};

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
		setTags(option);
		return false;
	};

	const handleSubmit = (e) => {
		e.preventDefault();
		let postMedia = [];
		let tagCollection = [];

		showedMedia.forEach((media) => {
			postMedia.push(media.content);
		});

		tags.forEach((tag) => {
			tagCollection.push({username: tag.value, id: tag.id});
		});

		let post = {
			location,
			description,
			postMedia: postMedia,
			tags: tagCollection,
		};

		console.log(post);
		postService.createPost(post, dispatch);
	};

	return (
		<React.Fragment>
			<div className="container ">
				<SuccessAlert
					hidden={!postState.createPost.showSuccessMessage}
					header="Success"
					message={postState.createPost.successMessage}
					handleCloseAlert={() => dispatch({ type: postConstants.CREATE_POST_REQUEST })}
				/>
				<FailureAlert
					hidden={!postState.createPost.showError}
					header="Error"
					message={postState.createPost.errorMessage}
					handleCloseAlert={() => dispatch({ type: postConstants.CREATE_POST_REQUEST })}
				/>
				<h3 className="text-center mt-4">Create new post</h3>
				<div className="row">
					<div className="col-12 mt-4">
						<form className="forms-sample" method="post" onSubmit={handleSubmit}>
							<div className="form-group row">
								<div className="col-6 float-left">
									<div className="input-group ml-0">
										<div className="input-group-append">
											<button className="btn btn-sm" type="button">
												<i className="mdi mdi-map-marker"></i>
											</button>
										</div>
										<input type="text" className="form-control" placeholder="Location" onChange={(e) => setLocation(e.target.value)} />
									</div>
								</div>
							</div>
							<hr />
							<div className="form-group">
								<MediaInputField showedMedia={showedMedia} setShowedMedia={setShowedMedia} />
							</div>
							<hr />

							<div className="form-group">
								<label for="tags">Tag people</label>

								<AsyncSelect isMulti defaultOptions loadOptions={loadOptions} onInputChange={onInputChange} onChange={onChange} placeholder="search" inputValue={search} />
							</div>

							<div className="form-group">
								<label for="description">Description</label>
								<textarea className="form-control" id="description" rows="3" placeholder="Description..." onChange={(e) => setDescription(e.target.value)}></textarea>
							</div>
							<div className="form-group">
								<button type="submit" className="btn btn-primary float-right">
									Create
								</button>
							</div>
						</form>
					</div>
				</div>
			</div>
		</React.Fragment>
	);
};

export default CreatePostForm;
