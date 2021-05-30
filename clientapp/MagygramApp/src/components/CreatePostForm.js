import React, { useState } from "react";
import MediaInputField from "./MediaInputField";
import { v4 as uuidv4 } from "uuid";
import TagsListInput from "./TagsListInput";

const CreatePostForm = () => {
	const [location, setLocation] = useState("");
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

	return (
		<React.Fragment>
			<div className="container ">
				<h3 className="text-center mt-2">Create new post</h3>
				<div className="row">
					<div className="col-12 mt-2">
						<form className="forms-sample" method="put">
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
								<MediaInputField />
							</div>
							<hr />

							<div className="form-group">
								<label for="tags">Tag people</label>

								<TagsListInput list={tags} handleItemDelete={handleTagDelete} handleItemAdd={handleAddTag} itemInput={tagInput} setItemInput={setTagInput} />
							</div>

							<div className="form-group">
								<label for="description">Description</label>
								<textarea className="form-control" id="description" rows="3" placeholder="Description..."></textarea>
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
