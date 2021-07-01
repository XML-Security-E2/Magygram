import React, { useContext, useRef, useState } from "react";
import MediaInputField from "../MediaInputField";
import { PostContext } from "../../contexts/PostContext";
import { postService } from "../../services/PostService";
import { postConstants } from "../../constants/PostConstants";
import SuccessAlert from "../SuccessAlert";
import FailureAlert from "../FailureAlert";
import AsyncSelect from "react-select/async";
import { searchService } from "../../services/SearchService";
import CreateCampaignForm from "./CreateCampaignForm";

const CreateAgentPostForm = () => {
	const { postState, dispatch } = useContext(PostContext);

	const [location, setLocation] = useState("");
	const [description, setDescription] = useState("");
	const [showedMedia, setShowedMedia] = useState([]);
	const [tags, setTags] = useState([]);
	const [showCampaignFields, setShowCampaignFields] = useState(false);
	const [search, setSearch] = useState("");

	const campaignRef = useRef();

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
			tagCollection.push({ username: tag.value, id: tag.id });
		});

		let post = {
			location,
			description,
			postMedia: postMedia,
			tags: tagCollection,
		};

		console.log(campaignRef.current.campaignState);

		console.log(post);
		if (!showCampaignFields) {
			postService.createPost(post, dispatch);
		} else {
			post.minAge = campaignRef.current.campaignState.minAge;
			post.maxAge = campaignRef.current.campaignState.maxAge;
			post.displayTime = campaignRef.current.campaignState.displayTime;
			if (campaignRef.current.campaignState.checkedOnce) {
				post.frequency = "ONCE";
			} else {
				post.frequency = "REPEATEDLY";
			}
			post.gender = campaignRef.current.campaignState.gender;
			post.startDate = campaignRef.current.campaignState.startDate.getTime();
			post.endDate = campaignRef.current.campaignState.endDate.getTime();

			postService.createPostCampaign(post, dispatch);
		}
	};

	return (
		<React.Fragment>
			<div className="container ">
				<SuccessAlert
					hidden={!postState.createAgentPost.showSuccessMessage}
					header="Success"
					message={postState.createAgentPost.successMessage}
					handleCloseAlert={() => dispatch({ type: postConstants.CREATE_POST_REQUEST })}
				/>
				<FailureAlert
					hidden={!postState.createAgentPost.showError}
					header="Error"
					message={postState.createAgentPost.errorMessage}
					handleCloseAlert={() => dispatch({ type: postConstants.CREATE_POST_REQUEST })}
				/>
				<h3 className="text-center mt-4">Create new post</h3>
				<br />
				<div className="row">
					<div className={showCampaignFields ? "col-6 mt-4" : "col-12 mt-4"}>
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
								<button type="button" hidden={!showCampaignFields} className="btn btn-outline-secondary btn-icon-text float-right" onClick={() => setShowCampaignFields(false)}>
									<i className="mdi mdi-chevron-left btn-icon-append"></i>Regular
								</button>
								<button type="button" hidden={showCampaignFields} className="btn btn-outline-secondary btn-icon-text float-right" onClick={() => setShowCampaignFields(true)}>
									Campaign <i className="mdi mdi-chevron-right btn-icon-append"></i>
								</button>

								<button type="submit" className="btn btn-primary float-left">
									Create post {showCampaignFields && "campaign"}
								</button>
							</div>
						</form>
					</div>
					<div hidden={!showCampaignFields} className="col-6">
						<CreateCampaignForm ref={campaignRef} fontColor="black" />
					</div>
				</div>
			</div>
		</React.Fragment>
	);
};

export default CreateAgentPostForm;
