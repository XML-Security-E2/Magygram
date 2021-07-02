import React, { useContext, useEffect } from "react";
import { Link } from "react-router-dom";
import { modalConstants } from "../../constants/ModalConstants";
import { MessageContext } from "../../contexts/MessageContext";
import { messageService } from "../../services/MessageService";

const StoryPreview = ({ storyId, story }) => {
	const { dispatch } = useContext(MessageContext);
	const imgStyle = { transform: "scale(1.5)", width: "100%", position: "absolute", left: "0" };

	const handleOpenStoryMessageModal = () => {
		dispatch({ type: modalConstants.SHOW_STORY_MESSAGE_MODAL, story });
	};

	useEffect(() => {
		messageService.findStoryById(storyId, dispatch);
	}, [storyId]);

	return (
		<div className="border rounded-lg w-100 no-gutters">
			<div className="col-12" hidden={!story.Unauthorized}>
				<div className="p-1">
					This story is unavailable because{" "}
					<a className="text-primary" href={"#/profile?userId=" + story.UserId}>
						{story.Username}
					</a>{" "}
					account is private
				</div>
			</div>
			<div className="col-12" hidden={!story.Expired}>
				<div className="p-1">
					This story from{" "}
					<a className="text-primary" href={"#/profile?userId=" + story.UserId}>
						{story.Username}
					</a>{" "}
					is unavailable, it's expired
				</div>
			</div>
			<div className="col-12" hidden={story.Expired || story.Unauthorized}>
				<div className="container-select-img" onClick={handleOpenStoryMessageModal} style={{ cursor: "pointer" }}>
					{story.MediaType === "IMAGE" ? (
						<img src={story.Url} className="img-fluid rounded-lg w-100" alt="" />
					) : (
						<video className="img-fluid box-coll rounded-lg w-100" style={{ objectFit: "cover" }}>
							<source src={story.Url} type="video/mp4" />
						</video>
					)}

					<div className="overlay-select-img rounded d-flex align-items-end">
						<div className="d-flex flex-row align-items-center">
							<div className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border message-profile-photo m-2">
								<img style={imgStyle} src={story.UserImageUrl === "" ? "assets/img/profile.jpg" : story.UserImageUrl} alt="" />
							</div>
							<div className="profile-info ml-2">
								<Link className="profile-info-username text-white" to={"/profile?userId=" + story.UserId}>
									{story.Username}
								</Link>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	);
};

export default StoryPreview;
