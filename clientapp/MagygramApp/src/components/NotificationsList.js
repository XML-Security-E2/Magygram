import React, { useContext } from "react";
import { NotificationContext } from "../contexts/NotificationContext";

const NotificationsList = () => {
	const { notificationState } = useContext(NotificationContext);
	const imgStyle = { transform: "scale(1.5)", width: "100%", position: "absolute", left: "0" };
	const visitedStoryClassName = "rounded-circle overflow-hidden d-flex justify-content-center align-items-center border border-danger story-profile-photo-visited ml-2";

	return (
		<React.Fragment>
			<li className="mb-3">
				<b className="ml-2">Notifications</b>
			</li>
			{notificationState.notifications !== null &&
				notificationState.notifications.map((notification) => {
					return (
						<li key={notification.id}>
							<div className="row d-flex justify-content-center align-items-center">
								<div className="col-2">
									<div className={visitedStoryClassName}>
										<img src={notification.imageUrl} alt="" style={imgStyle} />
									</div>
								</div>
								<div className="col-7">
									<span className="ml-1">{notification.note}</span>
								</div>
								<div className="col-3 text-right">
									<span style={{ fontSize: "0.8em", color: "gray" }} className="mr-3">
										<b>
											{Math.abs(new Date() - new Date(notification.time)) > 1000 * 60 * 60 * 24
												? new Date(notification.time).getDay() + "/" + new Date(notification.time).getUTCMonth()
												: new Date(notification.time).getHours() + ":" + new Date(notification.time).getMinutes()}
										</b>
									</span>
								</div>
							</div>
							<hr />
						</li>
					);
				})}
		</React.Fragment>
	);
};

export default NotificationsList;
