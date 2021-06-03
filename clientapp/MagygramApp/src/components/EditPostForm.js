import React, { useContext, useEffect, useState } from "react";
import { postConstants } from "../constants/PostConstants";
import { PostContext } from "../contexts/PostContext";
import { v4 as uuidv4 } from "uuid";
import FailureAlert from "./FailureAlert";
import PostImageSlider from "./PostImageSlider";
import SuccessAlert from "./SuccessAlert";
import TagsListInput from "./TagsListInput";
import { postService } from "../services/PostService";

const EditPostForm = () => {
	const { postState, dispatch } = useContext(PostContext);

	const [location, setLocation] = useState("");
	const [description, setDescription] = useState("");
	const [tags, setTags] = useState([]);
	const [tagInput, setTagInput] = useState("");

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

	const handleSubmit = (e) => {
		e.preventDefault();

		let tagNames = [];

		tags.forEach((tag) => {
			tagNames.push(tag.EntityDTO.Name);
		});

		let post = {
			id: postState.editPost.post.id,
			location: location,
			description: description,
			tags: tagNames,
		};

		postService.editPost(post, dispatch);
	};

	useEffect(() => {
		setLocation(postState.editPost.post.location);
		setDescription(postState.editPost.post.description);

		let tags = [];
		postState.editPost.post.tags.forEach((tag) => {
			tags.push({ Id: uuidv4(), EntityDTO: { Name: tag } });
		});

		setTags(tags);
	}, [postState.editPost.post]);

	return (
		<React.Fragment>
			<div className="container ">
				<SuccessAlert
					hidden={!postState.editPost.showSuccessMessage}
					header="Success"
					message={postState.editPost.successMessage}
					handleCloseAlert={() => dispatch({ type: postConstants.EDIT_POST_REQUEST })}
				/>
				<FailureAlert
					hidden={!postState.editPost.showError}
					header="Error"
					message={postState.editPost.errorMessage}
					handleCloseAlert={() => dispatch({ type: postConstants.EDIT_POST_REQUEST })}
				/>
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
										<input type="text" className="form-control" value={location} placeholder="Location" onChange={(e) => setLocation(e.target.value)} />
									</div>
								</div>
							</div>
							<hr />
							<div className="form-group">
								<PostImageSlider media={postState.editPost.post.media} />
							</div>
							<hr />

							<div className="form-group">
								<label for="tags">Tag people</label>
								<TagsListInput list={tags} handleItemDelete={handleTagDelete} handleItemAdd={handleAddTag} itemInput={tagInput} setItemInput={setTagInput} />
							</div>

							<div className="form-group">
								<label for="description">Description</label>
								<textarea
									className="form-control"
									id="description"
									value={description}
									rows="3"
									placeholder="Description..."
									onChange={(e) => setDescription(e.target.value)}
								></textarea>
							</div>
							<div className="form-group">
								<button type="submit" className="btn btn-primary float-right">
									Edit
								</button>
							</div>
						</form>
					</div>
				</div>
			</div>
		</React.Fragment>
	);
};

export default EditPostForm;
