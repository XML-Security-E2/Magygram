import React, { useContext, useState } from "react";
import { Modal } from "react-bootstrap";
import { storyConstants } from "../../constants/StoryConstants";
import { StoryContext } from "../../contexts/StoryContext";
import { storyService } from "../../services/StoryService";
import AsyncSelect from "react-select/async";
import { searchService } from "../../services/SearchService";

const CreateStoryModal = () => {
	const { storyState, dispatch } = useContext(StoryContext);
	const [showedMedia, setShowedMedia] = useState("");
	const [search, setSearch] = useState("");
	const [tags, setTags] = useState([]);

	const imgRef = React.createRef();
	const videoRef = React.createRef();

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
			tagCollection.push({username: tag.value, id: tag.id});
		});
		let story = {
			media: showedMedia.content,
			tags: tagCollection,
		};
		console.log(story);
		storyService.createStory(story, dispatch);
	};

	const handleModalClose = () => {
		dispatch({ type: storyConstants.CREATE_STORY_REQUEST });
	};

	const loadOptions = (value, callback) => {
		setTimeout(() => {
			searchService.userSearchUsers(value, callback);
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
		<Modal show={storyState.createStory.showModal} className="story_modal" aria-labelledby="contained-modal-title-vcenter" onHide={handleModalClose}>
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
				<div style={{ height: "85vh", cursor: "pointer" }} className="d-flex align-items-center justify-content-center" centered>
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
							</div>
						</div>
					)}
				</div>
			</Modal.Body>
		</Modal>
	);
};

export default CreateStoryModal;
