import React, { useState, useContext } from "react";
import { userService } from "../services/UserService";
import { Link } from "react-router-dom";
import AsyncSelect from "react-select/async";
import { searchService } from "../services/SearchService";
import FollowRequestsList from "./FollowRequestsList";
import { UserContext } from "../contexts/UserContext";
import { hasRoles } from "../helpers/auth-header"

const Header = () => {
	const userCtx = useContext(UserContext);

	const navStyle = { height: "50px", borderBottom: "1px solid rgb(200,200,200)" };
	const iconStyle = { fontSize: "30px", cursor: "pointer" };

	const [search, setSearch] = useState("");

	const loadOptions = (value, callback) => {
		if (value.startsWith("#") && value.length>=2) {
			setTimeout(() => {
				searchService.guestSearchHashtagPosts(value, callback);
			}, 1000);
		} else if (value.startsWith("%") && value.length>=2) {
			setTimeout(() => {
				searchService.guestSearchLocation(value, callback);
			}, 1000);
		} else {
			setTimeout(() => {
				searchService.userSearchUsers(value, callback);
			}, 1000);
		}
	};

	const onInputChange = (inputValue, { action }) => {
		switch (action) {
			case "set-value":
				return;
			case "menu-close":
				setSearch("");
				return;
			case "input-change":
				setSearch(inputValue);
				return;
			default:
				return;
		}
	};

	const onChange = (option) => {
		if (option.searchType === "hashtag") {
			window.location = "#/search/hashtag/" + option.value;
		} else if (option.searchType === "location") {
			window.location = "#/search/location/" +option.value;

		} else {
			window.location = "#/profile?userId=" + option.id;
		}

		return false;
	};

	const handleLogout = () => {
		userService.logout();
	};

	const handleSettings = () => {
		alert("TOD1O");
	};

	const handleLikedPosts = () => {
		window.location = "#/liked";
	};

	const handleDisikedPosts = () => {
		window.location = "#/disliked";
	};

	const backToHome = () => {
        window.location = "#/"
    }
	
	const handleLoadFollowRequests = async () => {
		await userService.findAllFollowRequests(userCtx.dispatch);
	};

	return (
		<nav className="navbar navbar-light navbar-expand-md navigation-clean" style={navStyle}>
			<div className="container">
				<div>
					<img onClick={() =>backToHome()} src="assets/img/logotest.png" alt="NistagramLogo" />
				</div>
				<button className="navbar-toggler" data-toggle="collapse">
					<span className="sr-only">Toggle navigation</span>
					<span className="navbar-toggler-icon"></span>
				</button>
				<div style={{ width: "300px" }}>
					<AsyncSelect defaultOptions loadOptions={loadOptions} onInputChange={onInputChange} onChange={onChange} placeholder="search" inputValue={search} />
				</div>
				<div className="d-flex align-items-center dropdown">
					<i hidden={hasRoles(["admin"])} className="fa fa-home ml-3" style={iconStyle} />
					<i hidden={hasRoles(["admin"])} className="la la-wechat ml-3" style={iconStyle} />
					<i hidden={hasRoles(["admin"])} className="la la-compass ml-3" style={iconStyle} />

					<div hidden={hasRoles(["admin"])}>
						<i className="fa fa-heart-o ml-3" onClick={handleLoadFollowRequests} style={iconStyle} id="dropdownMenu2" data-toggle="dropdown" />

						<ul style={{ width: "200px", marginLeft: "15px", minWidth: "300px" }} className="dropdown-menu" aria-labelledby="dropdownMenu2">
							<li className="mb-3">
								<b className="ml-2">Follow requests</b>
							</li>
							<FollowRequestsList />
						</ul>
					</div>

					<div hidden={hasRoles(["admin"])}>
						<img
							className="rounded-circle overflow-hidden border border-danger header-image-photo dropdown-toggle ml-3"
							style={{ cursor: "pointer" }}
							src={localStorage.getItem("imageURL") !== "" ? localStorage.getItem("imageURL") : "assets/img/profile.jpg"}
							alt=""
							id="dropdownMenu1"
							data-toggle="dropdown"
							aria-haspopup="true"
						/>
						<ul style={{ width: "200px", marginLeft: "15px" }} className="dropdown-menu" aria-labelledby="dropdownMenu1">
							<li>
								<Link className="la la-user btn shadow-none" to={"/profile?userId=" + localStorage.getItem("userId")}>
									Profile
								</Link>
							</li>
							<li>
								<button className="la la-cog btn shadow-none" onClick={handleSettings}>
									Settings
								</button>
							</li>
							<li>
								<button className="la la-thumbs-up btn shadow-none" onClick={handleLikedPosts}>
									Liked posts
								</button>
							</li>
							<li>
								<button className="la la-thumbs-down btn shadow-none" onClick={handleDisikedPosts}>
									Disiked posts
								</button>
							</li>
							<hr className="solid" />
							<li>
								<button className=" btn shadow-none" onClick={handleLogout}>
									Logout
								</button>
							</li>
						</ul>
					</div>
				</div>
			</div>
		</nav>
	);
};

export default Header;
