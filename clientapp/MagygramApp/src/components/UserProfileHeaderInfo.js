import { useContext, useEffect } from "react";
import { Link } from "react-router-dom";
import { colorConstants } from "../constants/ColorConstants";
import { UserContext } from "../contexts/UserContext";
import { userService } from "../services/UserService";
import FollowingUsersModal from "./modals/FollowingUsersModal";
import Axios from "axios";
import { authHeader } from "../helpers/auth-header";
import { notificationService } from "../services/NotificationService";
import { NotificationContext } from "../contexts/NotificationContext";
import NotificationSettingsModal from "./modals/NotificationSettingsModal";

const UserProfileHeaderInfo = ({ userId }) => {
	const { userState, dispatch } = useContext(UserContext);
	const ntfxCtx = useContext(NotificationContext);

	const imgProfileStyle = { left: "20", width: "150px", height: "150px", marginLeft: "100px", borderWidth: "1px", borderStyle: "solid" };

	useEffect(() => {
		const getProfileHandler = async () => {
			await userService.getUserProfileByUserId(userId, dispatch);
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

	const reportUser = () => {
		let reportDTO = {
			contentId: userId,
			contentType: "USER",
		};

		Axios.post(`/api/report`, reportDTO, { validateStatus: () => true, headers: authHeader() })
			.then((res) => {
				console.log(res.data);
				if (res.status === 200) {
					alert("You have reported this user successfully");
				} else {
					console.log("err");
				}
			})
			.catch((err) => {
				console.log("err");
			});
	};
	
	const deleteUser = () => {
		console.log(localStorage.getItem("roles"))
		let requestId =  userId;
		Axios.put(`/api/users/${requestId}/delete`, {}, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);
			if (res.status === 200) {
				console.log("User has been deleted");
				alert("You have successfully deleted this user!")
			} else {
				console.log("ERROR")
			}
		})
		.catch((err) => {
			console.log("ERROR")
		});
	}

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
						<div>
							{localStorage.getItem("userId") !== null &&
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
							<button
								hidden={localStorage.getItem("userId") === userId || localStorage.getItem("userId") === null}
								style={{ backgroundColor: "red", borderColor: "red" }}
								type="button"
								className="btn btn-primary ml-2"
								tabindex="0"
								onClick={reportUser}
							>
								Report
							</button>
						</div>
						<div>
							<button
								hidden={!(localStorage.getItem("userId") === userId)}
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
