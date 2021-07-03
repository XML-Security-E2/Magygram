import React, { useState, useContext, useRef, useEffect, createRef } from "react";
import { userService } from "../services/UserService";
import { Link } from "react-router-dom";
import AsyncSelect from "react-select/async";
import { searchService } from "../services/SearchService";
import { UserContext } from "../contexts/UserContext";
import { notificationService } from "../services/NotificationService";
import { NotificationContext } from "../contexts/NotificationContext";
import { notificationConstants } from "../constants/NotificationConstants";
import ActivityList from "./ActivityList";
import { hasRoles } from "../helpers/auth-header";

const Header = () => {
	const userCtx = useContext(UserContext);
	const notifyCtx = useContext(NotificationContext);

	const navStyle = { height: "50px", borderBottom: "1px solid rgb(200,200,200)" };
	const iconStyle = { fontSize: "30px", cursor: "pointer" };

	const [search, setSearch] = useState("");
	const [loadData, setLoadData] = useState(true);

	const [rmessages, setRmessages] = useState([]);
	const [notifyMessages, setNotifyMessages] = useState([]);

	const ws = useRef(null);
	const messagesWs = useRef(null);

	useEffect(() => {
		console.log(process.env);
		ws.current = new WebSocket(
			process.env.REACT_APP_STAGE === "prod" ? "ws://localhost:467/ws/notify/" + localStorage.getItem("userId") : "wss://localhost:467/ws/notify/" + localStorage.getItem("userId")
		);
		ws.current.onopen = () => console.log("ws opened");
		ws.current.onclose = () => console.log("ws closed");

		messagesWs.current = new WebSocket(
			process.env.REACT_APP_STAGE === "prod"
				? "ws://localhost:467/ws/notify/messages/" + localStorage.getItem("userId")
				: "wss://localhost:467/ws/notify/messages/" + localStorage.getItem("userId")
		);
		messagesWs.current.onopen = () => console.log("ws opened");
		messagesWs.current.onclose = () => console.log("ws closed");
	}, []);

	useEffect(() => {
		if (!ws.current) return;

		ws.current.onmessage = (evt) => {
			console.log(evt.data);
			let a = [...rmessages];
			a.push(evt.data);

			setRmessages(a);
			notifyCtx.dispatch({ type: notificationConstants.NOTIFICATION_RECEIVED, count: JSON.parse(evt.data).count });
		};
	}, [rmessages]);

	useEffect(() => {
		if (!messagesWs.current) return;

		messagesWs.current.onmessage = (evt) => {
			console.log(evt.data);
			let a = [...notifyMessages];
			a.push(evt.data);

			setNotifyMessages(a);
			notifyCtx.dispatch({ type: notificationConstants.MESSAGE_NOTIFICATION_RECEIVED, count: JSON.parse(evt.data).count });
		};
	}, [notifyMessages]);

	const loadOptions = (value, callback) => {
		if (value.startsWith("#") && value.length >= 2) {
			setTimeout(() => {
				searchService.guestSearchHashtagPosts(value, callback);
			}, 1000);
		} else if (value.startsWith("%") && value.length >= 2) {
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
			window.location = "#/search/location/" + option.value;
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

	const handleCampaignApiToken = () => {
		window.location = "#/campaign-api";

	}

	const campaignOffers = () => {
		window.location = "#/influencer-campagns";
	};

	const backToHome = () => {
		window.location = "#/";
	};

	const handleViewNotifications = () => {
		return new Promise(function () {
			notificationService.viewNotifications(notifyCtx.dispatch);
		});
	};

	const handleLoadActivity = async () => {
		if (loadData) {
			await userService.findAllFollowRequests(userCtx.dispatch);

			await notificationService.getUserNotifiactions(notifyCtx.dispatch).then(handleViewNotifications());
		}
		setLoadData(!loadData);
	};

	return (
		<nav className="navbar navbar-light navbar-expand-md navigation-clean" style={navStyle}>
			<div className="container">
				<div>
					<img onClick={() => backToHome()} src="assets/img/logotest.png" alt="NistagramLogo" />
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
					<Link hidden={hasRoles(["admin"])} className="count-indicator ml-3" to="/chat">
						<i className="la la-wechat shadow-none" style={{ fontSize: "30px", cursor: "pointer", color: "black" }}></i>
						{notifyCtx.notificationState.messageNotificationsNumber > 0 && <span className="count count-varient1">{notifyCtx.notificationState.messageNotificationsNumber}</span>}
					</Link>
					<i hidden={hasRoles(["admin"])} className="la la-compass ml-3" style={iconStyle} />

					<div hidden={hasRoles(["admin"])}>
						<div className="count-indicator ml-3" id="dropdownMenu3" data-toggle="dropdown" onClickCapture={handleLoadActivity}>
							<i className="la la-bell" style={{ fontSize: "30px", cursor: "pointer", color: "black" }}></i>
							{notifyCtx.notificationState.notificationsNumber > 0 && <span className="count count-varient1">{notifyCtx.notificationState.notificationsNumber}</span>}
						</div>
						<ul
							style={{ width: "200px", marginLeft: "15px", minWidth: "370px", height: "auto", maxHeight: "500px", overflowX: "hidden" }}
							className="dropdown-menu  dropdown-menu-right"
							aria-labelledby="dropdownMenu3"
						>
							<ActivityList />
						</ul>
					</div>

					<div>
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
							<li hidden={hasRoles(["admin"])}>
								<Link className="la la-user btn shadow-none" to={"/profile?userId=" + localStorage.getItem("userId")}>
									Profile
								</Link>
							</li>
							<li hidden={!hasRoles(["agent"])}>
								<Link className="la la-volume-up btn shadow-none" to="/campaigns">
									Campaigns
								</Link>
							</li>
							<li hidden={hasRoles(["admin"])}>
								<button className="la la-cog btn shadow-none" onClick={handleSettings}>
									Settings
								</button>
							</li>
							<li hidden={hasRoles(["admin"])}>
								<button className="la la-thumbs-up btn shadow-none" onClick={handleLikedPosts}>
									Liked posts
								</button>
							</li>
							<li hidden={hasRoles(["admin"])}>
								<button className="la la-thumbs-down btn shadow-none" onClick={handleDisikedPosts}>
									Disiked posts
								</button>
							</li>
							<li hidden={!hasRoles(["agent"])}>
								<button className="la la-key btn shadow-none" onClick={handleCampaignApiToken}>
									Campaign API Token
								</button>
							</li>
							<hr hidden={hasRoles(["admin"])} className="solid" />
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
