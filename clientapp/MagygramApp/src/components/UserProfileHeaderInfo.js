import { useContext, useEffect } from "react";
import { Link } from "react-router-dom";
import { colorConstants } from "../constants/ColorConstants";
import { UserContext } from "../contexts/UserContext";
import { userService } from "../services/UserService";
import FollowingUsersModal from "./modals/FollowingUsersModal";

const UserProfileHeaderInfo = ({ userId }) => {
	const { userState, dispatch } = useContext(UserContext);

	const sectionStyle = { left: "20", marginLeft: "100px" };
	const imgProfileStyle = { left: "20", width: "150px", height: "150px", marginLeft: "100px", borderWidth: "1px", borderStyle: "solid" };
	const editStyle = { color: "black", left: "20", marginLeft: "13px", marginRight: "13px", borderWidth: "1px", borderStyle: "solid" };

	useEffect(() => {
		const getProfileHandler = async () => {
			await userService.getUserProfileByUserId(userId, dispatch);
		};
		getProfileHandler();
	}, [userId]);

	const getFollowersHandler = async () => {
		await userService.findAllFollowedUsers(userId, dispatch);
	};

	const getFollowingHandler = async () => {
		await userService.findAllFollowingUsers(userId, dispatch);
	};

	return (
		<nav className="navbar navbar-light navbar-expand-md navigation-clean" style={{ backgroundColor: colorConstants.COLOR_BACKGROUND }}>
			<div className="flexbox-container">
				<div>
					<img className="rounded-circle dropdown-toggle" style={imgProfileStyle} src={userState.userProfile.user.imageUrl} alt="" />
				</div>
				<section style={sectionStyle}>
					<div className="flexbox-container">
						<div>
							<h2>{userState.userProfile.user.username}</h2>
						</div>
						<div>
							<Link hidden={userId !== localStorage.getItem("userId")} style={editStyle} to="/edit-profile" tabindex="0">
								Edit Profile
							</Link>
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
						<h4>
							{userState.userProfile.user.name} {userState.userProfile.user.surname}
						</h4>
					</div>
					<div>{userState.userProfile.user.bio}</div>
					<div>
						<a href={userState.userProfile.user.website}>{userState.userProfile.user.website}</a>
					</div>
				</section>
			</div>
			<FollowingUsersModal />
		</nav>
	);
};

export default UserProfileHeaderInfo;
