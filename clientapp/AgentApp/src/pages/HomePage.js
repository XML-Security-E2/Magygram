import React, {useState} from "react";
import { userService } from "../services/UserService";
import Axios from "axios";
import { config } from "../config/config";

const HomePage = () => {
	const navStyle = { height: "50px", borderBottom: "1px solid rgb(200,200,200)" };
	const inputStyle = { border: "1px solid rgb(200,200,200)", color: "rgb(210,210,210)", textAlign: "center" };
	const iconStyle = { fontSize: "30px", margin: "0px", marginLeft: "13px" };
	const imgStyle = { width: "30px", height: "30px", marginLeft: "13px", borderWidth: "1px", borderStyle: "solid" };

	const [name, setName] = useState('')

	const handleLogout = () => {
		userService.logout();
	}

	const handleProfile = () =>{
		Axios.get(`${config.API_URL}/users/637300ee-4abd-484a-8422-bc87f7cf82ff`, { validateStatus: () => true })
		.then((res) => {
			console.log(res.data)
			setName(res.data.Name)
		})
		.catch((err) => {});
	}

	const handleSettings = () =>{
		alert('TOD1O')
	}

	return (
		<nav className="navbar navbar-light navbar-expand-md navigation-clean" style={navStyle}>
			<div className="container">
				<div>
					<img src="assets/img/logotest.png" alt="NistagramLogo" />
				</div>
				<button className="navbar-toggler" data-toggle="collapse">
					<span className="sr-only">Toggle navigation</span>
					<span className="navbar-toggler-icon"></span>
				</button>
				<div>
					<input type="text" style={inputStyle} placeholder="Search" value="Search" />
				</div>
				<div className="d-xl-flex align-items-xl-center dropdown">
					<i className="fa fa-home" style={iconStyle} />
					<i className="la la-wechat" style={iconStyle} />
					<i className="la la-compass" style={iconStyle} />
					<i className="fa fa-heart-o" style={iconStyle} />
					<img className="rounded-circle dropdown-toggle" data-toggle="dropdown" style={imgStyle} src="assets/img/hitmanImage.jpeg" alt="ProfileImage"/>
					<ul style={{width:'200px', marginLeft:'15px'}} class="dropdown-menu">
						<li><button className="la la-user btn shadow-none" onClick={handleProfile}>  Profile</button></li>
						<li><button className="la la-cog btn shadow-none" onClick={handleSettings}>  Settings</button></li>
						<hr className="solid"/>
						<li><button className=" btn shadow-none" onClick={handleLogout}>  Logout</button></li>
					</ul>
				</div>
				<div>
					{name}
				</div>
			</div>
		</nav>
	);
};

export default HomePage;