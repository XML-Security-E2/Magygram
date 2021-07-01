import React, { useContext, useRef, useState } from "react";
import { Modal } from "react-bootstrap";
import { storyConstants } from "../../constants/StoryConstants";
import { StoryContext } from "../../contexts/StoryContext";
import { storyService } from "../../services/StoryService";
import AsyncSelect from "react-select/async";
import { searchService } from "../../services/SearchService";
import CreateCampaignForm from "../agent-post-create/CreateCampaignForm";

const CreateAgentStoryModal = () => {
	const { storyState, dispatch } = useContext(StoryContext);
	const [showedMedia, setShowedMedia] = useState("");
	const [showCampaignFields, setShowCampaignFields] = useState(false);

	const [search, setSearch] = useState("");
	const [tags, setTags] = useState([]);

	const imgRef = React.createRef();
	const videoRef = React.createRef();
	const campaignRef = useRef();

	const onImageChange = (e) => {
		if (e.target.files && e.target.files[0]) {
			let img = e.target.files[0];
			let filename = URL.createObjectURL(img);
			let extension = e.target.files[0].name.split(".").pop();

			if (extension === "mp4") {
				setShowedMedia({ src: filename, type: "video", content: e.target.files[0] });
			} else setShowedMedia({ src: filename, type: "image", content: e.target.files[0] });
		}
	};

	const handleImageDeselect = () => {
		setShowedMedia("");
	};

	const handleStoryUpload = () => {
		let tagCollection = [];
		tags.forEach((tag) => {
			tagCollection.push({ username: tag.value, id: tag.id });
		});
		let story = {
			media: showedMedia.content,
			tags: tagCollection,
		};
		console.log(story);
		if (!showCampaignFields) {
			storyService.createStory(story, dispatch);
		} else {
			story.minAge = campaignRef.current.campaignState.minAge;
			story.maxAge = campaignRef.current.campaignState.maxAge;
			story.displayTime = campaignRef.current.campaignState.displayTime;
			if (campaignRef.current.campaignState.checkedOnce) {
				story.frequency = "ONCE";
			} else {
				story.frequency = "REPEATEDLY";
			}
			story.gender = campaignRef.current.campaignState.gender;
			story.startDate = campaignRef.current.campaignState.startDate.getTime();
			story.endDate = campaignRef.current.campaignState.endDate.getTime();

			storyService.createAgentStory(story, dispatch);
		}
	};

	const handleModalClose = () => {
		dispatch({ type: storyConstants.CREATE_AGENT_STORY_REQUEST });
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

	return (
		<Modal show={storyState.createAgentStory.showModal} className="story_modal" aria-labelledby="contained-modal-title-vcenter" onHide={handleModalClose}>
			<Modal.Body>
				<input type="file" ref={imgRef} style={{ display: "none" }} name="image" accept="image/png, image/jpeg, video/mp4" onChange={onImageChange} />
				<div style={{ opacity: showedMedia !== "" ? 1 : 0 }} className="d-flex flex-row-reverse">
					<button
						hidden={false}
						disabled={showedMedia === ""}
						className="btn btn-outline-dark btn-icon btn-rounded border-0"
						data-toggle="tooltip"
						title="Delete image"
						onClick={handleImageDeselect}
					>
						<i className="mdi mdi-close text-danger"></i>
					</button>
				</div>
				<div hidden={showCampaignFields} style={{ height: "85vh", cursor: "pointer" }} className={showCampaignFields ? "" : "d-flex align-items-center justify-content-center"} centered>
					{showedMedia === "" ? (
						<div>
							<div className="row d-flex align-items-center justify-content-center" onClick={() => imgRef.current.click()}>
								<button type="button" className="btn btn-outline-secondary rounded-lg btn-icon">
									<i className="mdi mdi-plus w-50 h-50"></i>
								</button>
							</div>
							<div className="row mt-5" style={{ color: "white", fontSize: "1em" }}>
								Click to add story
							</div>
						</div>
					) : (
						<div>
							<div hidden={showedMedia === ""} className="row d-flex align-items-center justify-content-center">
								{showedMedia.type === "image" ? (
									<img src={showedMedia.src} className="img-fluid rounded-lg w-100 " alt="" />
								) : (
									<video ref={videoRef} controls className="img-fluid rounded-lg w-100">
										<source src={showedMedia.src} type="video/mp4" />
									</video>
								)}
							</div>
							<div hidden={showedMedia === ""} className="row d-flex align-items-center justify-content-center">
								<div style={{ width: "300px", margin: "10px" }}>
									<AsyncSelect isMulti defaultOptions loadOptions={loadOptions} onInputChange={onInputChange} onChange={onChange} placeholder="search" inputValue={search} />
								</div>
								<button type="button" onClick={handleStoryUpload} className="btn btn-outline-secondary btn-icon-text border-0">
									<i className="mdi mdi-upload btn-icon-prepend"></i> Upload
								</button>
								<button type="button" className="btn btn-outline-secondary btn-icon-text float-right" onClick={() => setShowCampaignFields(true)}>
									Campaign <i className="mdi mdi-chevron-right btn-icon-append"></i>
								</button>
							</div>
						</div>
					)}
				</div>

				<div className={!showCampaignFields ? "" : "d-flex row no-gutters align-items-center"} hidden={!showCampaignFields} style={{ height: "85vh" }} centered>
					<div className="col-4">
						<button type="button" hidden={!showCampaignFields} className="btn btn-outline-secondary btn-icon-text float-left" onClick={() => setShowCampaignFields(false)}>
							<i className="mdi mdi-chevron-left btn-icon-append"></i>Regular
						</button>
					</div>

					<div className="col-8">
						<CreateCampaignForm ref={campaignRef} fontColor="white" />
					</div>
					<div className="col-12">
						<button type="button" hidden={!showCampaignFields} className="btn btn-outline-secondary btn-icon-text float-right" onClick={handleStoryUpload}>
							Upload campaign<i className="mdi mdi-chevron-right btn-icon-append"></i>
						</button>
					</div>
				</div>
			</Modal.Body>
		</Modal>
	);
};

export default CreateAgentStoryModal;
