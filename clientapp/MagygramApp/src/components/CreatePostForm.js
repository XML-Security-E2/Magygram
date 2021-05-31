import React, { useContext, useState } from "react";
import MediaInputField from "./MediaInputField";
import { v4 as uuidv4 } from "uuid";
import TagsListInput from "./TagsListInput";
import { PostContext } from "../contexts/PostContext";
import { postService } from "../services/PostService";

const CreatePostForm = () => {
	const { postState, dispatch } = useContext(PostContext);

	const [location, setLocation] = useState("");
	const [description, setDescription] = useState("");
	const [showedMedia, setShowedMedia] = useState([]);
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
		let postMedia = [];
		let tagNames = [];

		showedMedia.forEach((media) => {
			postMedia.push(media.content);
		});

		tags.forEach((tag) => {
			tagNames.push(tag.EntityDTO.Name);
		});

		let post = {
			location,
			description,
			postMedia: postMedia,
			tags: tagNames[0],
		};

		console.log(post);
		postService.createPost(post, dispatch);
	};

	return (
		<React.Fragment>
			<div className="container ">
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

								<TagsListInput list={tags} handleItemDelete={handleTagDelete} handleItemAdd={handleAddTag} itemInput={tagInput} setItemInput={setTagInput} />
							</div>

							<div className="form-group">
								<label for="description">Description</label>
								<textarea className="form-control" id="description" rows="3" placeholder="Description..." onChange={(e) => setDescription(e.target.value)}></textarea>
							</div>
							<div className="form-group">
								<img src="https://localhost:463/api/media/e5fc6a3a-67fc-4bc8-9740-18eeb1b17942.jpg" />
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
