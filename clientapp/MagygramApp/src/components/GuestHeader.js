import React, { useState } from "react";
import Axios from "axios";
import { useHistory } from "react-router-dom";
import { authHeader } from "../helpers/auth-header";
import AsyncSelect from "react-select/async";

const GuestHeader = () => {
	const history = useHistory();
	const navStyle = { height: "50px", borderBottom: "1px solid rgb(200,200,200)" };
	const iconStyle = { fontSize: "30px", margin: "0px", marginLeft: "13px" };
	const imgStyle = { width: "30px", height: "30px", marginLeft: "13px", borderWidth: "1px", borderStyle: "solid" };

	const [name, setName] = useState("");
	const [img, setImg] = useState("");
	const [currentId, setCurrentId] = useState();
	const [search, setSearch] = useState("");

	const loadOptions = (value, callback) => {
		setTimeout(() => {
			var options;
			Axios.get(`/api/users/search/` + value, { validateStatus: () => true, headers: authHeader() })
			.then((res) => {
				console.log(res.data);
				if (res.status === 200) {
					options = res.data.map(option => ({ value: option.Username, label: option.Username, id: option.Id}))
					callback(options);
				}})
			.catch((err) => {
				console.log(err)
			});
		}, 1000);
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
		if (currentId === option.id) {
			window.location = "#/profile";
		} else {
			window.location = "#/user/" + option.id;
		}
		return false;
	};

    const backToHome = () => {
        window.location = "#/"
    }

    const handleRegistration = () =>{
        window.location = "#/registration"
    }

    const handleLogin = () =>{
        window.location = "#/login"
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
				<div style={{ width: "300px" }}>
					<AsyncSelect defaultOptions loadOptions={loadOptions} onInputChange={onInputChange} onChange={onChange} placeholder="search" inputValue={search} />
				</div>
				<div className="d-xl-flex align-items-xl-center dropdown">
					<i className="fa fa-home" onClick={backToHome()} style={iconStyle} />
					<i className="fa fa-user dropdown-toggle"  data-toggle="dropdown" style={iconStyle} />
					<ul style={{ width: "200px", marginLeft: "15px" }} class="dropdown-menu">
						<li>
							<button onClick={handleRegistration} className="fa fa-user-plus btn shadow-none" >
                                <label className="ml-2">Registruj se</label>
							</button>
						</li>
                        <hr className="solid" />
						<li>
							<button onClick={handleLogin} className="fa fa-sign-in btn shadow-none" >
								<label className="ml-2">Uloguj se</label>
							</button>
						</li>
                    </ul>
            
				</div>
			</div>
		</nav>
	);
};

export default GuestHeader;
