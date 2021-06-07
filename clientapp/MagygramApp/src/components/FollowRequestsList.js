import React, { useContext } from "react";
import { Link } from "react-router-dom";
import { UserContext } from "../contexts/UserContext";
import { userService } from "../services/UserService";

const FollowRequestsList = () => {
	const { userState, dispatch } = useContext(UserContext);
	const imgStyle = { transform: "scale(1.5)", width: "100%", position: "absolute", left: "0" };
	const visitedStoryClassName = "rounded-circle overflow-hidden d-flex justify-content-center align-items-center border border-danger story-profile-photo-visited ml-2";

	const handleAcceptRequest = (userId) => {
		userService.acceptFollowRequest(userId, dispatch);
	};

	return (
		<React.Fragment>
			{userState.userFollowRequests.userInfos !== null ? (
				userState.userFollowRequests.userInfos.map((user) => {
					return (
						<li key={user.userInfo.id}>
							<div className="row d-flex justify-content-center align-items-center">
								<div className="col-2">
									<div className={visitedStoryClassName}>
										<img src={user.userInfo.imageUrl} alt="" style={imgStyle} />
									</div>
								</div>
								<div className="col-6">
									<Link type="button" className="float-left btn btn-link btn-fw" style={{ color: "black" }} to={"/profile?userId=" + user.userInfo.id}>
										<b>{user.userInfo.username}</b>
									</Link>
								</div>
								<div className="col-4">
									{!user.following ? (
										<button type="button" className="btn btn-secondary" onClick={() => handleAcceptRequest(user.userInfo.id)}>
											Accept
										</button>
									) : (
										<div>Following</div>
									)}
								</div>
							</div>
							<hr />
						</li>
					);
				})
			) : (
				<li className="mt-5 d-flex justify-content-center ">
					<div>
						<h4 className="d-flex justify-content-center text-secondary">No follow requests</h4>
					</div>
				</li>
			)}
		</React.Fragment>
	);
};

export default FollowRequestsList;
