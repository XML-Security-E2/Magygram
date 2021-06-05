import React, { useState, useEffect } from "react";
import { userService } from "../services/UserService";
import Axios from "axios";
import { authHeader } from "../helpers/auth-header";
import { colorConstants } from "../constants/ColorConstants";
import Header from "../components/Header";
import StoryHighlights from "../components/StoryHighlights";
import StoryContextProvider from "../contexts/StoryContext";
import UserStoriesModal from "../components/modals/UserStoriesModal";

const ProfilePage = () => {
	// const history = useHistory();
	// const navStyle = { height: "50px", borderBottom: "1px solid rgb(200,200,200)" };
	// const inputStyle = { border: "1px solid rgb(200,200,200)", color: "rgb(210,210,210)", textAlign: "center" };
	// const iconStyle = { fontSize: "30px", margin: "0px", marginLeft: "13px" };
	const iconStyle1 = { fontSize: "30px", margin: "0px", marginLeft: "200px" };
	//const imgStyle = { left: "0", width: "30px", height: "30px", marginLeft: "13px", borderWidth: "1px", borderStyle: "solid" };
	const imgProfileStyle = { left: "20", width: "150px", height: "150px", marginLeft: "100px", borderWidth: "1px", borderStyle: "solid" };
	const nameStyle = { left: "20", marginLeft: "13px" };
	const editStyle = { color: "black", left: "20", marginLeft: "13px", marginRight: "13px", borderWidth: "1px", borderStyle: "solid" };
	const sectionStyle = { left: "20", marginLeft: "100px" };
	const [name, setName] = useState("");
	const [username, setUsername] = useState("");
	const [bio, setBio] = useState("");
	const [img, setImg] = useState("");

	const handleLogout = () => {
		userService.logout();
	};

    useEffect(() => {
        Axios.get(`/api/users/logged`, { validateStatus: () => true, headers: authHeader() })
			.then((res) => {
				console.log(res.data);
                setUsername(res.data.username);
                if(res.data.imageUrl == "")
                    setImg("assets/img/profile.jpg");
                else
                setImg(res.data.imageUrl);
                
                Axios.get(`/api/users/` + res.data.id, { validateStatus: () => true, headers: authHeader() })
                    .then((res) => {
                        
				        console.log(res.data);
                        setName(res.data.Name);
                        setBio(res.data.Bio);

                    })
                    .catch((err) => {console.log(err);});

			})
			.catch((err) => {
				console.log(err);
			});
	});

	const handleProfile = () => {
		window.location = `#/profile`;

		Axios.get(`/api/users/logged`, { validateStatus: () => true, headers: authHeader() })
			.then((res) => {
                setUsername(res.data.username);
                
                Axios.get(`/api/users/` + res.data.id, { validateStatus: () => true, headers: authHeader() })
                    .then((res) => {
                        setName(res.data.Name);
                        setBio(res.data.Bio);

                    })
                    .catch((err) => {console.log(err);});

			})
			.catch((err) => {
				console.log(err);
			});
	};

	const handleSettings = () => {
		alert("TOD1O");
	};

	return (
		<React.Fragment>
			<Header />
			<StoryContextProvider>
				<div>
					<div className="mt-4">
						<div className="container d-flex justify-content-center">
							<div className="col-12">
								<div className="row">
									<div className="col-12">
										<nav className="navbar navbar-light navbar-expand-md navigation-clean" style={{ backgroundColor: colorConstants.COLOR_BACKGROUND }}>
											<div className="flexbox-container">
												<div>
													<img className="rounded-circle dropdown-toggle" style={imgProfileStyle} src={img} alt="" />
												</div>
												<section style={sectionStyle}>
													<div className="flexbox-container">
														<div>
															<h2>{username}</h2>
														</div>
														<div>
															<a style={editStyle} href="#/edit-profile" tabindex="0">
																Edit Profile
															</a>
														</div>
													</div>
													<div className="flexbox-container">
														<div>0 posts</div>
														<div style={nameStyle}>0 followers</div>
														<div style={nameStyle}>0 followings</div>
													</div>
													<br />
													<div>
														<h4>{name}</h4>
													</div>
													<div>{bio}</div>
												</section>
											</div>
										</nav>
										<nav>
											<StoryHighlights />
										</nav>
										<hr />
										<nav>
											<div className="container">
												<i className="fa fa-th" style={iconStyle1} /> POSTS
												<i className="fa fa-bookmark" style={iconStyle1} /> SAVED
												<i className="fa fa-tag" style={iconStyle1} /> TAGGED
											</div>
										</nav>
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>
				<UserStoriesModal />
			</StoryContextProvider>
		</React.Fragment>
	);
};

export default ProfilePage;
