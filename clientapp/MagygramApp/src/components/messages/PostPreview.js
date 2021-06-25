import React, { useContext, useEffect } from "react";
import { Link } from "react-router-dom";
import { MessageContext } from "../../contexts/MessageContext";
import { messageService } from "../../services/MessageService";

const PostPreview = ({ postId, post }) => {
	const { dispatch } = useContext(MessageContext);
	const imgStyle = { transform: "scale(1.5)", width: "100%", position: "absolute", left: "0" };

	const handlePostRedirect = () => {
		window.location = "#/post?postId=" + postId;
	};

	useEffect(() => {
		messageService.findPostById(postId, dispatch);
	}, [postId]);

	return (
		<div className="border rounded-lg w-100 no-gutters">
			<div className="col-12" hidden={!post.Unauthorized}>
				<div className="p-1">
					This post is unavailable because{" "}
					<a className="text-primary" href={"#/profile?userId=" + post.UserId}>
						{post.Username}
					</a>{" "}
					account is private
				</div>
			</div>
			<div className="col-12" hidden={post.Unauthorized}>
				<div className="d-flex flex-row align-items-center">
					<div className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border message-profile-photo m-2">
						<img style={imgStyle} src={post.UserImageUrl === "" ? "assets/img/profile.jpg" : post.UserImageUrl} alt="" />
					</div>
					<div className="profile-info ml-2">
						<Link className="profile-info-username text-dark" to={"/profile?userId=" + post.UserId}>
							{post.Username}
						</Link>
					</div>
				</div>
			</div>
			<div className="col-12" hidden={post.Unauthorized} onClick={handlePostRedirect}>
				<div style={{ cursor: "pointer" }}>
					{post.MediaType === "IMAGE" ? (
						<img src={post.Url} className="img-fluid box-coll  w-100" alt="" style={{ objectFit: "cover" }} />
					) : (
						<video className="img-fluid box-coll w-100" style={{ objectFit: "cover" }}>
							<source src={post.Url} type="video/mp4" />
						</video>
					)}
				</div>
			</div>
			<div className="col-12" hidden={post.Unauthorized}>
				<div className="m-2">
					<b className="mr-2">{post.Username}</b>
					{post.Description.substring(0, 25)}
				</div>
			</div>
		</div>
	);
};

export default PostPreview;
