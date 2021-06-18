import { useContext, useEffect, useState } from "react";
import { UserContext } from "../contexts/UserContext";
import { userService } from "../services/UserService";

const EditNotificationsForm = ({ show }) => {
	const { userState, dispatch } = useContext(UserContext);

	const [notifyLike, setNotifyLike] = useState(userState.userProfile.user.notificationSettings.notifyLike);
	const [notifyDislike, setNotifyDislike] = useState(userState.userProfile.user.notificationSettings.notifyDislike);
	const [notifyComments, setNotifyComments] = useState(userState.userProfile.user.notificationSettings.notifyComments);
	const [notifyFollow, setNotifyFollow] = useState(userState.userProfile.user.notificationSettings.notifyFollow);
	const [notifyFollowRequest, setNotifyFollowRequest] = useState(userState.userProfile.user.notificationSettings.notifyFollowRequest);
	const [notifyAcceptFollowRequest, setNotifyAcceptFollowRequest] = useState(userState.userProfile.user.notificationSettings.notifyAcceptFollowRequest);

	const handleSubmit = (e) => {
		e.preventDefault();

		let notificationReq = {
			notifyLike,
			notifyDislike,
			notifyComments,
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
		setNotifyFollow(userState.userProfile.user.notificationSettings.notifyFollow);
		setNotifyFollowRequest(userState.userProfile.user.notificationSettings.notifyFollowRequest);
		setNotifyAcceptFollowRequest(userState.userProfile.user.notificationSettings.notifyAcceptFollowRequest);
	}, [userState.userProfile.user]);

	return (
		<form hidden={!show} method="post" onSubmit={handleSubmit}>
			<div>
				<br />
				<div className="form-group row d-flex align-items-center">
					<label className="col-sm-4 col-form-label">Like notifications</label>
					<div className="col-sm-2">
						<label>Mute </label>

						<input type="checkbox" className="mr-1" checked={notifyLike == 0} onChange={() => setNotifyLike(0)} />
					</div>
					<div className="col-sm-3">
						<label>People I follow </label>

						<input type="checkbox" className="mr-1" checked={notifyLike == 1} onChange={() => setNotifyLike(1)} />
					</div>
					<div className="col-sm-3">
						<label>Everyone </label>

						<input type="checkbox" className="mr-1" checked={notifyLike == 2} onChange={() => setNotifyLike(2)} />
					</div>
				</div>
				<div className="form-group row d-flex align-items-center">
					<label className="col-sm-4 col-form-label">Dislike notifications</label>
					<div className="col-sm-2">
						<label>Mute </label>

						<input type="checkbox" className="mr-1" checked={notifyDislike == 0} onChange={() => setNotifyDislike(0)} />
					</div>
					<div className="col-sm-3">
						<label>People I follow </label>

						<input type="checkbox" className="mr-1" checked={notifyDislike == 1} onChange={() => setNotifyDislike(1)} />
					</div>
					<div className="col-sm-3">
						<label>Everyone</label>

						<input type="checkbox" className="mr-1" checked={notifyDislike == 2} onChange={() => setNotifyDislike(2)} />
					</div>
				</div>
				<div className="form-group row d-flex align-items-center">
					<label className="col-sm-4 col-form-label">Comment notifications</label>
					<div className="col-sm-2 -flex align-items-center">
						<label>Mute</label>

						<input type="checkbox" className="mr-1" checked={notifyComments == 0} onChange={() => setNotifyComments(0)} />
					</div>
					<div className="col-sm-3 -flex align-items-center">
						<label>People I follow</label>

						<input type="checkbox" className="mr-1" checked={notifyComments == 1} onChange={() => setNotifyComments(1)} />
					</div>
					<div className="col-sm-3 -flex align-items-center">
						<label>Everyone</label>

						<input type="checkbox" className="mr-1" checked={notifyComments == 2} onChange={() => setNotifyComments(2)} />
					</div>
				</div>
				<div className="form-group row d-flex align-items-center">
					<label className="col-sm-6 col-form-label">Notify when user follow me</label>
					<div className="col-sm-4">
						<label>
							<input type="checkbox" className="mr-1" checked={notifyFollow} onChange={() => setNotifyFollow(!notifyFollow)} />
						</label>
					</div>
				</div>
				<div className="form-group row d-flex align-items-center">
					<label className="col-sm-6 col-form-label">Notify when receive follow request</label>
					<div className="col-sm-4">
						<label>
							<input type="checkbox" className="mr-1" checked={notifyFollowRequest} onChange={() => setNotifyFollowRequest(!notifyFollowRequest)} />
						</label>
					</div>
				</div>
				<div className="form-group row d-flex align-items-center">
					<label className="col-sm-6 col-form-label">Accepted follow request notifications</label>
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
