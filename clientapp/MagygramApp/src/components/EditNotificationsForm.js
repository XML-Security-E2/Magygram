import { useContext, useEffect, useState } from "react";
import { UserContext } from "../contexts/UserContext";
import { userService } from "../services/UserService";

const EditNotificationsForm = ({ show }) => {
	const { userState, dispatch } = useContext(UserContext);

	const [notifyLike, setNotifyLike] = useState(userState.userProfile.user.notificationSettings.notifyLike);
	const [notifyDislike, setNotifyDislike] = useState(userState.userProfile.user.notificationSettings.notifyDislike);
	const [notifyComments, setNotifyComments] = useState(userState.userProfile.user.notificationSettings.notifyComments);
	const [notifyPost, setNotifyPost] = useState(userState.userProfile.user.notificationSettings.notifyPost);
	const [notifyStory, setNotifyStory] = useState(userState.userProfile.user.notificationSettings.notifyStory);
	const [notifyFollow, setNotifyFollow] = useState(userState.userProfile.user.notificationSettings.notifyFollow);
	const [notifyFollowRequest, setNotifyFollowRequest] = useState(userState.userProfile.user.notificationSettings.notifyFollowRequest);
	const [notifyAcceptFollowRequest, setNotifyAcceptFollowRequest] = useState(userState.userProfile.user.notificationSettings.notifyAcceptFollowRequest);

	const handleSubmit = (e) => {
		e.preventDefault();

		let notificationReq = {
			notifyLike,
			notifyDislike,
			notifyComments,
			notifyPost,
			notifyStory,
			notifyFollow,
			notifyFollowRequest,
			notifyAcceptFollowRequest,
		};

		userService.editUserNotifications(userState.userProfile.showedUserId, notificationReq, dispatch);
	};

	useEffect(() => {
		setNotifyLike(userState.userProfile.user.notificationSettings.notifyLike);
		setNotifyDislike(userState.userProfile.user.notificationSettings.notifyDislike);
		setNotifyComments(userState.userProfile.user.notificationSettings.notifyComments);
		setNotifyPost(userState.userProfile.user.notificationSettings.notifyPost);
		setNotifyStory(userState.userProfile.user.notificationSettings.notifyStory);
		setNotifyFollow(userState.userProfile.user.notificationSettings.notifyFollow);
		setNotifyFollowRequest(userState.userProfile.user.notificationSettings.notifyFollowRequest);
		setNotifyAcceptFollowRequest(userState.userProfile.user.notificationSettings.notifyAcceptFollowRequest);
	}, [userState.userProfile.user]);

	return (
		<form hidden={!show} method="post" onSubmit={handleSubmit}>
			<div>
				<br />
				<div className="form-group row">
					<label className="col-sm-8 col-form-label">Story notifications</label>
					<div className="col-sm-4">
						<label>
							<input type="checkbox" className="mr-1" checked={notifyStory} onChange={() => setNotifyStory(!notifyStory)} />
						</label>
					</div>
				</div>
				<div className="form-group row d-flex align-items-center">
					<label className="col-sm-8 col-form-label">Post notifications</label>
					<div className="col-sm-4">
						<label>
							<input type="checkbox" className="mr-1" checked={notifyPost} onChange={() => setNotifyPost(!notifyPost)} />
						</label>
					</div>
				</div>
				<div className="form-group row d-flex align-items-center">
					<label className="col-sm-8 col-form-label">Like notifications</label>
					<div className="col-sm-4">
						<label>
							<input type="checkbox" className="mr-1" checked={notifyLike} onChange={() => setNotifyLike(!notifyLike)} />
						</label>
					</div>
				</div>
				<div className="form-group row d-flex align-items-center">
					<label className="col-sm-8 col-form-label">Dislike notifications</label>
					<div className="col-sm-4">
						<label>
							<input type="checkbox" className="mr-1" checked={notifyDislike} onChange={() => setNotifyDislike(!notifyDislike)} />
						</label>
					</div>
				</div>
				<div className="form-group row d-flex align-items-center">
					<label className="col-sm-8 col-form-label">Post comment notifications</label>
					<div className="col-sm-4">
						<label>
							<input type="checkbox" className="mr-1" checked={notifyComments} onChange={() => setNotifyComments(!notifyComments)} />
						</label>
					</div>
				</div>
				<div className="form-group row d-flex align-items-center">
					<label className="col-sm-8 col-form-label">Notify when user follow me</label>
					<div className="col-sm-4">
						<label>
							<input type="checkbox" className="mr-1" checked={notifyFollow} onChange={() => setNotifyFollow(!notifyFollow)} />
						</label>
					</div>
				</div>
				<div className="form-group row d-flex align-items-center">
					<label className="col-sm-8 col-form-label">Notify when receive follow request</label>
					<div className="col-sm-4">
						<label>
							<input type="checkbox" className="mr-1" checked={notifyFollowRequest} onChange={() => setNotifyFollowRequest(!notifyFollowRequest)} />
						</label>
					</div>
				</div>
				<div className="form-group row d-flex align-items-center">
					<label className="col-sm-8 col-form-label">Accepted follow request notifications</label>
					<div className="col-sm-4">
						<label>
							<input type="checkbox" className="mr-1" checked={notifyAcceptFollowRequest} onChange={(e) => setNotifyAcceptFollowRequest(!notifyAcceptFollowRequest)} />
						</label>
					</div>
				</div>
				<div className="form-group">
					<button className="btn btn-primary float-right  mb-2" type="submit">
						Save
					</button>
				</div>
			</div>
		</form>
	);
};

export default EditNotificationsForm;
