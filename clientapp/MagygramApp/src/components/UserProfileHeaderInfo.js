import { useContext, useEffect } from "react";
import { Link } from "react-router-dom";
import { colorConstants } from "../constants/ColorConstants";
import { UserContext } from "../contexts/UserContext";
import { userService } from "../services/UserService";
import FollowingUsersModal from "./modals/FollowingUsersModal";
import Axios from "axios";
import { authHeader } from "../helpers/auth-header";
import { hasRoles } from "../helpers/auth-header";
import { notificationService } from "../services/NotificationService";
import { NotificationContext } from "../contexts/NotificationContext";
import NotificationSettingsModal from "./modals/NotificationSettingsModal";
import { ProfileSettingsContext } from "../contexts/ProfileSettingsContext";
import { modalConstants } from "../constants/ModalConstants";

const UserProfileHeaderInfo = ({ userId }) => {
	const { userState, dispatch } = useContext(UserContext);
	const { profileSettingsState, profileSettingsDispatch } = useContext(ProfileSettingsContext);

	const ntfxCtx = useContext(NotificationContext);

	const imgProfileStyle = { left: "20", width: "150px", height: "150px", marginLeft: "100px", borderWidth: "1px", borderStyle: "solid" };

	useEffect(() => {
		const getProfileHandler = async () => {
			await userService.getUserProfileByUserId(userId, dispatch);
			userService.IsUserVerifiedById(userId, profileSettingsDispatch);
		};
		getProfileHandler();
	}, [userId, dispatch]);

	const getFollowersHandler = async () => {
		await userService.findAllFollowedUsers(userId, dispatch);
	};

	const getFollowingHandler = async () => {
		await userService.findAllFollowingUsers(userId, dispatch);
	};

	const handleUserFollow = () => {
		userService.followUser(userId, dispatch);
	};

	const handleUserUnfollow = () => {
		userService.unfollowUser(userId, dispatch);
	};

	const handleUserMute = () => {
		userService.muteUser(userId, dispatch);
	};

	const handleUserUnmute = () => {
		userService.unmuteUser(userId, dispatch);
	};

	const handleUserBlock = () => {
		userService.blockUser(userId, dispatch);
	};

	const handleUserUnblock = () => {
		userService.unblockUser(userId, dispatch);
	};

	const reportUser = () => {
		dispatch({ type: modalConstants.SHOW_USER_REPORT_MODAL });
	};

	const deleteUser = () => {
		console.log(localStorage.getItem("roles"));
		let requestId = userId;
		Axios.put(`/api/users/${requestId}/delete`, {}, { validateStatus: () => true, headers: authHeader() })
			.then((res) => {
				console.log(res);
				if (res.status === 200) {
					console.log("User has been deleted");
					alert("You have successfully deleted this user!");
				} else {
					console.log("ERROR");
				}
			})
			.catch((err) => {
				console.log("ERROR");
			});
	};

	const handleNotificationsSetttings = async () => {
		await notificationService.getProfileNotificationsSettings(userId, ntfxCtx.dispatch);
	};

	return (
		<nav className="navbar navbar-light navbar-expand-md navigation-clean" style={{ backgroundColor: colorConstants.COLOR_BACKGROUND }}>
			<div className="flexbox-container">
				<div className="mr-5">
					<img
						className="rounded-circle dropdown-toggle"
						style={imgProfileStyle}
						src={userState.userProfile.user.imageUrl === "" ? "assets/img/profile.jpg" : userState.userProfile.user.imageUrl}
						alt=""
					/>
				</div>
				<section className="ml-5">
					<div className="flexbox-container d-flex align-items-center">
						<div>
							<h2>{userState.userProfile.user.username}</h2>
						</div>
						<div className="ml-2" hidden={!profileSettingsState.isUserVerified}>
							<i class="bi bi-patch-check"></i>
							<label className="ml-1">Verified</label>
						</div>
						<div>
							{localStorage.getItem("userId") !== null &&
								userState.userProfile.user.blocked === false &&
								(userState.userProfile.user.sentFollowRequest ? (
									<h5 className="text-secondary ml-2">Following request sent</h5>
								) : (
									userId !== localStorage.getItem("userId") &&
									(localStorage.getItem("userId") !== null && userState.userProfile.user.following ? (
										<button type="button" className="btn btn-outline-secondary ml-2" tabindex="0" onClick={handleUserUnfollow}>
											Unfollow
										</button>
									) : (
										<button type="button" className="btn btn-primary ml-2" tabindex="0" onClick={handleUserFollow}>
											Follow
										</button>
									))
								))}
							{userState.userProfile.user.following && (
								<button type="button" className="btn btn-outline-secondary ml-2 btn-rounded btn-icon" onClick={handleNotificationsSetttings}>
									<i className="mdi mdi-bell"></i>
								</button>
							)}
							<Link
								type="button"
								hidden={userId !== localStorage.getItem("userId")}
								className="btn btn-outline-secondary ml-2"
								style={{ color: "black" }}
								to="/edit-profile"
								tabindex="0"
							>
								Edit Profile
							</Link>
						</div>
						<div>
							{localStorage.getItem("userId") !== null &&
								userId !== localStorage.getItem("userId") &&
								userState.userProfile.user.following &&
								userState.userProfile.user.blocked === false &&
								(userState.userProfile.user.muted ? (
									<button type="button" className="btn btn-outline-secondary ml-2" tabindex="0" onClick={handleUserUnmute}>
										Unmute
									</button>
								) : (
									<button type="button" className="btn btn-primary ml-2" tabindex="0" onClick={handleUserMute}>
										Mute
									</button>
								))}
						</div>
						<div>
							{localStorage.getItem("userId") !== null &&
								userId !== localStorage.getItem("userId") &&
								(userState.userProfile.user.blocked ? (
									<button type="button" className="btn btn-outline-secondary ml-2" tabindex="0" onClick={handleUserUnblock}>
										Unblock
									</button>
								) : (
									<button type="button" className="btn btn-primary ml-2" tabindex="0" onClick={handleUserBlock}>
										Block
									</button>
								))}
						</div>
						<div>
							<button
								hidden={localStorage.getItem("userId") === userId || localStorage.getItem("userId") === null}
								type="button"
								className="btn btn-outline-secondary ml-2"
								tabindex="0"
								onClick={reportUser}
							>
								Report
							</button>
						</div>
						<div>
							<button
								hidden={!hasRoles(["admin"])}
								style={{ backgroundColor: "red", borderColor: "red" }}
								type="button"
								className="btn btn-primary ml-2"
								tabindex="0"
								onClick={deleteUser}
							>
								Delete profile
							</button>
						</div>
					</div>
					<div className="flexbox-container d-flex align-items-center">
						<div>{userState.userProfile.user.postNumber} posts</div>
						{localStorage.getItem("userId") === null ? (
							<div className="ml-3">{userState.userProfile.user.followersNumber} followers</div>
						) : (
							<button type="button" className="btn btn-link btn-fw ml-2" style={{ color: "black" }} onClick={getFollowersHandler}>
								{userState.userProfile.user.followersNumber} followers
							</button>
						)}
						{localStorage.getItem("userId") === null ? (
							<div className="ml-3">{userState.userProfile.user.followingNumber} followings</div>
						) : (
							<button type="button" className="btn btn-link btn-fw" style={{ color: "black" }} onClick={getFollowingHandler}>
								{userState.userProfile.user.followingNumber} followings
							</button>
						)}
					</div>
					<br />
					<div>
						<h5>
							{userState.userProfile.user.name} {userState.userProfile.user.surname}
						</h5>
					</div>
					<div>{userState.userProfile.user.bio}</div>
					<div>
						<a href={userState.userProfile.user.website}>{userState.userProfile.user.website}</a>
					</div>
				</section>
			</div>
			<NotificationSettingsModal />
			<FollowingUsersModal />
		</nav>
	);
};

export default UserProfileHeaderInfo;
