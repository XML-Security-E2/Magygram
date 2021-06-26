import React, { useContext, useEffect, useState } from "react";
import { MessageContext } from "../../contexts/MessageContext";
import { getUserInfo } from "../../helpers/auth-header";
import { getDateTime } from "../../helpers/datetime-helper";
import ViewMediaButton from "./ViewMediaButton";
import ImageViewer from "react-simple-image-viewer";
import VideoViewerModal from "../modals/VideoViewerModal";
import { messageService } from "../../services/MessageService";
import PostPreview from "./PostPreview";
import StoryPreview from "./StoryPreview";

const Chat = () => {
	const { messageState, dispatch } = useContext(MessageContext);

	const [showImageViewer, setShowImageViewer] = useState(false);
	const [image, setImage] = useState("");

	const [showVideoViewer, setShowVideoViewer] = useState(false);
	const [video, setVideo] = useState("");

	const openMedia = (messageId, media) => {
		if (media.mediaType == "IMAGE") {
			setShowImageViewer(true);
			setImage([media.url]);
		} else {
			setShowVideoViewer(true);
			setVideo(media.url);
		}

		messageService.viewMediaMessages(messageState.selectedConversationId, messageId, dispatch);
	};

	const openMediaOwner = (media) => {
		if (media.mediaType == "IMAGE") {
			setShowImageViewer(true);
			setImage([media.url]);
		} else {
			setShowVideoViewer(true);
			setVideo(media.url);
		}
	};

	const closeImageViewer = (image) => {
		setImage([image]);
		setShowImageViewer(false);
	};

	const acceptMessageRequest = () => {
		messageService.acceptRequest(messageState.selectedConversationId, dispatch);
	};

	const denyMessageRequest = () => {
		messageService.denyRequest(messageState.selectedConversationId, dispatch);
	};

	const deleteMessageRequest = () => {
		messageService.deleteRequest(messageState.selectedConversationId, dispatch);
	};

	useEffect(() => {
		var a = document.querySelector("#conversation");
		console.log(a);
		a.scrollTop = a.scrollHeight - a.clientHeight;
	}, [messageState.showedMessages]);

	return (
		<React.Fragment>
			<div className="row flex-grow-1" id="conversation" style={{ overflowX: "hidden", maxHeight: "95%" }}>
				<div className="col-12">
					<div className="row">
						{messageState.showedMessages.map((message) => {
							if (message.messageFromId === getUserInfo().Id) {
								return (
									<div className="col-12 mt-2">
										<div className="row">
											{message.messageType === "TEXT" && (
												<div className="col-12 text-break ml-auto" style={{ maxWidth: "50%" }}>
													<div className="float-right rounded-lg pl-3 pb-2 pr-3 pt-2" style={{ backgroundColor: "lightgray" }}>
														{message.text}
													</div>
												</div>
											)}

											{message.messageType === "MEDIA" && (
												<div className="col-12 ml-auto" style={{ maxWidth: "50%" }}>
													{message.media.mediaType === "IMAGE" ? (
														<img
															src={message.media.url}
															className="img-fluid box-coll rounded-lg w-100 "
															alt=""
															onClick={() => openMediaOwner(message.media)}
															style={{ objectFit: "cover", cursor: "pointer" }}
														/>
													) : (
														<video
															className="img-fluid box-coll rounded-lg w-100"
															style={{ objectFit: "cover", cursor: "pointer" }}
															onClick={() => openMediaOwner(message.media)}
														>
															<source src={message.media.url} type="video/mp4" />
														</video>
													)}
												</div>
											)}

											{message.messageType === "POST" && (
												<div className="col-12 ml-auto" style={{ maxWidth: "50%" }}>
													<PostPreview postId={message.contentId} post={message.post} />
												</div>
											)}

											{message.messageType === "STORY" && (
												<div className="col-12 ml-auto" style={{ maxWidth: "50%" }}>
													<StoryPreview storyId={message.contentId} story={message.story} />
												</div>
											)}

											<div className="col-12">
												<span className="float-right" style={{ fontSize: "0.8em" }}>
													{getDateTime(message.timestamp)}
												</span>
											</div>
										</div>
									</div>
								);
							} else {
								return (
									<div className="col-12 mt-2">
										<div className="row">
											{message.messageType === "TEXT" && (
												<div className="col-12 text-break mr-auto" style={{ maxWidth: "50%" }}>
													<div className="float-left rounded-lg border  pl-3 pb-2 pr-3 pt-2">{message.text}</div>
												</div>
											)}

											{message.messageType === "MEDIA" && (
												<div className="col-12 mr-auto" style={{ maxWidth: "50%" }}>
													<ViewMediaButton
														messageId={message.id}
														disabled={message.viewedMedia}
														text={message.media.mediaType === "IMAGE" ? "Photo" : "Video"}
														media={message.media}
														openMedia={openMedia}
													/>
												</div>
											)}

											{message.messageType === "POST" && (
												<div className="col-12 mr-auto" style={{ maxWidth: "50%" }}>
													<PostPreview postId={message.contentId} post={message.post} />
												</div>
											)}

											{message.messageType === "STORY" && (
												<div className="col-12 mr-auto" style={{ maxWidth: "50%" }}>
													<StoryPreview storyId={message.contentId} story={message.story} />
												</div>
											)}

											<div className="col-12">
												<span className="float-left" style={{ fontSize: "0.8em" }}>
													{getDateTime(message.timestamp)}
												</span>
											</div>
										</div>
									</div>
								);
							}
						})}
						<div hidden={!messageState.loadedConversationRequests} className="col-12 mt-2">
							<div className="row">
								<div className="col-12 text-center">
									<button type="button" className="btn btn-outline-secondary border-0" onClick={acceptMessageRequest}>
										Accept request
									</button>
									<button type="button" className="btn btn-outline-secondary border-0" onClick={denyMessageRequest}>
										Deny request
									</button>
									<button type="button" className="btn btn-outline-danger border-0" onClick={deleteMessageRequest}>
										Delete request
									</button>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>

			{showImageViewer && (
				<ImageViewer
					src={image}
					onClose={closeImageViewer}
					backgroundStyle={{
						backgroundColor: "rgba(0,0,0,0.9)",
					}}
				/>
			)}

			<VideoViewerModal
				handleModalClose={() => {
					setShowVideoViewer(false);
					setVideo("");
				}}
				show={showVideoViewer}
				video={video}
			/>
		</React.Fragment>
	);
};

export default Chat;
