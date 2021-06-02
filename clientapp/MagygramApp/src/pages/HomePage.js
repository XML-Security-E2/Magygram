import React, { useState } from "react";
import { userService } from "../services/UserService";
import Axios from "axios";
import { config } from "../config/config";
import Timeline from "../components/Timeline";
import PostContextProvider from "../contexts/PostContext";
import StoryContextProvider from "../contexts/StoryContext";
import CreateStoryModal from "../components/modals/CreateStoryModal";
import StoryButton from "../components/StoryButton";
import { useHistory } from "react-router-dom";
import { authHeader } from "../helpers/auth-header";
import Header from "../components/Header";
import Storyline from "../components/Storyline"

const HomePage = () => {
	const history = useHistory();
	const navStyle = { height: "50px", borderBottom: "1px solid rgb(200,200,200)" };
	const inputStyle = { border: "1px solid rgb(200,200,200)", color: "rgb(210,210,210)", textAlign: "center" };
	const iconStyle = { fontSize: "30px", margin: "0px", marginLeft: "13px" };
	const imgStyle = { left: "0", width: "30px", height: "30px", marginLeft: "13px", borderWidth: "1px", borderStyle: "solid" };
	const [name, setName] = useState("");

	const handleLogout = () => {
		userService.logout();
	};

	const handleProfile = () => {

		let path = `/profile`; 
		history.push(path);

		Axios.get(`${config.API_URL}/users/logged`,{
			validateStatus: () => true,
			headers: { Authorization: authHeader()}
		})
			.then((res) => {
				console.log(res.data);
				setName(res.data.Name);
			})
			.catch((err) => {
				console.log(err);});
	};

	const handleSettings = () => {
		alert("TOD1O");
	};

	return (
		<React.Fragment>
			<Header/>
			<StoryContextProvider>
				<CreateStoryModal />
				<div>
					<div class="mt-4">
						<div class="container d-flex justify-content-center">
							<div class="col-9">
								<div class="row">
									<div class="col-8">
										<Storyline/>
										<PostContextProvider>
											<Timeline />
										</PostContextProvider>
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>
			</StoryContextProvider>
		</React.Fragment>
	);
};

export default HomePage
