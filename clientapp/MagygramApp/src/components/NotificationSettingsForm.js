import { useContext, useEffect, useState } from "react";
import { NotificationContext } from "../contexts/NotificationContext";
import { UserContext } from "../contexts/UserContext";
import { notificationService } from "../services/NotificationService";

const NotificationSettingsForm = () => {
	const { notificationState, dispatch } = useContext(NotificationContext);
	const usrCtx = useContext(UserContext);

	const [notifyPost, setNotifyPost] = useState(notificationState.notificationSettingsModal.settings.notifyPost);
	const [notifyStory, setNotifyStory] = useState(notificationState.notificationSettingsModal.settings.notifyStory);

	useEffect(() => {
		setNotifyPost(notificationState.notificationSettingsModal.settings.notifyPost);
		setNotifyStory(notificationState.notificationSettingsModal.settings.notifyStory);
	}, [notificationState.notificationSettingsModal.settings]);

	const handleSubmit = (e) => {
		e.preventDefault();

		let notificationReq = {
			notifyPost,
			notifyStory,
		};

		notificationService.editProfileNotificationsSettings(notificationReq, usrCtx.userState.userProfile.showedUserId, dispatch);
	};

	return (
		<div>
			<form method="post" onSubmit={handleSubmit}>
				<div>
					<div className="form-group row d-flex align-items-center">
						<label className="col-sm-6 col-form-label">Notify when user publish post</label>
						<div className="col-sm-4 d-flex align-items-center">
							<input type="checkbox" className="mr-1" checked={notifyPost} onChange={() => setNotifyPost(!notifyPost)} />
						</div>
					</div>
					<div className="form-group row d-flex align-items-center">
						<label className="col-sm-6 col-form-label">Notify when user publish story</label>
						<div className="col-sm-4 d-flex align-items-center">
							<input type="checkbox" className="mr-1" checked={notifyStory} onChange={() => setNotifyStory(!notifyStory)} />
						</div>
					</div>

					<div className="form-group">
						<button className="btn btn-primary float-right  mb-2" type="submit">
							Save
						</button>
					</div>
				</div>
			</form>
		</div>
	);
};

export default NotificationSettingsForm;
