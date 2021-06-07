import React, { useContext } from "react";
import { Link } from "react-router-dom";
import { modalConstants } from "../constants/ModalConstants";
import { UserContext } from "../contexts/UserContext";
import { userService } from "../services/UserService";

const UserFollowingModalList = () => {
	const { userState, dispatch } = useContext(UserContext);
	const imgStyle = { transform: "scale(1.5)", width: "100%", position: "absolute", left: "0" };
	const visitedStoryClassName = "rounded-circle overflow-hidden d-flex justify-content-center align-items-center border border-danger story-profile-photo-visited";

	const handleProfileClick = () => {
		dispatch({ type: modalConstants.HIDE_FOLLOWING_MODAL });
	};

	const handleUserFollow = (userId) => {
		userService.followUser(userId, dispatch);
	};

	const handleUserUnfollow = (userId) => {
		userService.unfollowUser(userId, dispatch);
	};

	return (
		<React.Fragment>
			{userState.userProfileFollowingModal.userInfos !== null ? (
				userState.userProfileFollowingModal.userInfos.map((user) => {
					return (
						<div key={user.userInfo.id}>
							<div className="row d-flex justify-content-center align-items-center">
								<div className="col-2">
									<div className={visitedStoryClassName}>
										<img src={user.userInfo.imageUrl} alt="..." style={imgStyle} />
									</div>
								</div>
								<div className="col-6">
									<Link type="button" className="float-left btn btn-link btn-fw" style={{ color: "black" }} onClick={handleProfileClick} to={"/profile?userId=" + user.userInfo.id}>
										<b>{user.userInfo.username}</b>
									</Link>
								</div>
								<div className="col-4">
									{user.following ? (
										<button type="button" className="btn btn-secondary" onClick={() => handleUserUnfollow(user.userInfo.id)}>
											Unfollow
										</button>
									) : (
										<button type="button" className="btn btn-secondary" onClick={() => handleUserFollow(user.userInfo.id)}>
											Follow
										</button>
									)}
								</div>
							</div>
							<hr />
						</div>
					);
				})
			) : (
				<div className="mt-5 d-flex justify-content-center ">
					<div>
						<h3 className="d-flex justify-content-center">No users found</h3>
					</div>
				</div>
			)}
		</React.Fragment>
	);
};

export default UserFollowingModalList;
