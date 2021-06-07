import React, { useState, useContext, useEffect } from "react";
import { userService } from "../services/UserService";
import Axios from "axios";
import { authHeader } from "../helpers/auth-header";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import Header from "../components/Header";
import PostContextProvider from "../contexts/PostContext";

const EditProfile = () => {
	// history = useHistory();
	const navStyle = { height: "50px", borderBottom: "1px solid rgb(200,200,200)" };
	const inputStyle = { border: "1px solid rgb(200,200,200)", color: "rgb(210,210,210)", textAlign: "center" };
	const iconStyle = { fontSize: "30px", margin: "0px", marginLeft: "13px" };
	//const iconStyle1 = { fontSize: "30px", margin: "0px", marginLeft: "200px" };
	const imgStyle = { left: "0", width: "30px", height: "30px", marginLeft: "13px", borderWidth: "1px", borderStyle: "solid" };
	// const imgProfileStyle = { left: "20", width: "150px", height: "150px", marginLeft: "100px", borderWidth: "1px", borderStyle: "solid" };
	// const nameStyle = { left: "20",  marginLeft: "13px"}
	// const editStyle = {color: "black", left: "20",  marginLeft: "13px",marginRight: "13px", borderWidth: "1px", borderStyle: "solid" }
	const sectionStyle = { left: "20", marginLeft: "100px" };

	const [email, setEmail] = useState("");
	const [id, setId] = useState("");
	const [name, setName] = useState("");
	const [surname, setSurname] = useState("");
	const [username, setUsername] = useState("");
	const [bio, setBio] = useState("");
	const [img, setImg] = useState("");
	const [website, setWebsite] = useState("");
	const [number, setNumber] = useState("");
	const [gender, setGender] = useState("");

	const [emailInput, setEmailInput] = useState("");
	const [nameInput, setNameInput] = useState("");
	const [surnameInput, setSurnameInput] = useState("");
	const [usernameInput, setUsernameInput] = useState("");
	const [bioInput, setBioInput] = useState("");
	//const [imgInput, setImgInput] = useState("");
	const [websiteInput, setWebsiteInput] = useState("");
	const [numberInput, setNumberInput] = useState("");
	const [genderInput, setGenderInput] = useState("");

	useEffect(() => {
		Axios.get(`/api/users/logged`, { validateStatus: () => true, headers: authHeader() })
			.then((res) => {
				setId(res.data.id);
				if (res.data.imageUrl == "") setImg("assets/img/profile.jpg");
				else setImg(res.data.imageUrl);

				Axios.get(`/api/users/` + res.data.id, { validateStatus: () => true, headers: authHeader() })
					.then((res) => {
						console.log(res.data);
						setName(res.data.Name);
						setSurname(res.data.Surname);
						setUsername(res.data.Username);
						setBio(res.data.Bio);
						setEmail(res.data.Email);
						setWebsite(res.data.Website);
						setNumber(res.data.Number);
						setGenderInput(res.data.Gender);
					})
					.catch((err) => {
						console.log(err);
					});
			})
			.catch((err) => {
				console.log(err);
			});
	});

	const handleSettings = () => {
		alert("TOD1O");
	};

	const handleSubmit = (e) => {
		e.preventDefault();
		var n = name;
		var sur = surname;
		var un = username;
		var e = email;
		var w = website;
		var num = num;
		var b = bio;
		var g = genderInput;

		if (gender == "") {
			g = "MALE";
		} else {
			g = gender;
		}

		if (nameInput == "") {
			n = name;
		} else {
			n = nameInput;
		}
		if (numberInput == "") {
			num = number;
		} else {
			num = numberInput;
		}
		if (bioInput == "") {
			b = bio;
		} else {
			b = bioInput;
		}
		if (surnameInput == "") {
			sur = surname;
		} else {
			sur = surnameInput;
		}

		if (usernameInput == "") {
			un = username;
		} else {
			un = usernameInput;
		}

		if (emailInput == "") {
			e = email;
		} else {
			e = emailInput;
		}

		if (websiteInput == "") {
			w = website;
		} else {
			w = websiteInput;
		}
		let user = {
			id,
			nameInput: n,
			surnameInput: sur,
			usernameInput: un,
			emailInput: e,
			websiteInput: w,
			numberInput: num,
			bioInput: b,
			gender: g,
		};
		console.log(user);

		Axios.post(`/api/users/edit`, user, { validateStatus: () => true })
			.then((res) => {
				window.location = "#/profile";
			})
			.catch((err) => {
				console.log(err);
			});
	};

	const handleLogout = () => {
		userService.logout();
	};

	const handleProfile = () => {
		window.location = `#/profile`;
	};

	const handleChange = (e) => {
		setName(e.target.value);
	};

	return (
		<React.Fragment>
			<PostContextProvider>
				<Header />
			</PostContextProvider>
			<div>
				<br />
				<br />
				<div className="container">
					<form method="post" onSubmit={handleSubmit}>
						<div>
							<h2 className="text-center">
								<strong>Edit</strong> profile
							</h2>
							<br />
							<div className="form-group">
								<text>Name</text>
								<input className="form-control" type="text" name="name" placeholder={name} value={nameInput} onChange={(e) => setNameInput(e.target.value)} />
							</div>
							<div className="form-group">
								<text>Surname</text>
								<input className="form-control" type="text" name="surname" placeholder={surname} value={surnameInput} onChange={(e) => setSurnameInput(e.target.value)} />
							</div>
							<div className="form-group">
								<text>Email</text>
								<input className="form-control" type="email" name="email" placeholder={email} value={emailInput} onChange={(e) => setEmailInput(e.target.value)} />
							</div>

							<div className="form-group">
								<text>Username</text>
								<input className="form-control" type="username" name="username" placeholder={username} value={usernameInput} onChange={(e) => setUsernameInput(e.target.value)} />
							</div>

							<div className="form-group">
								<text>Website</text>
								<input className="form-control" type="text" name="websiteInput" placeholder={website} value={websiteInput} onChange={(e) => setWebsiteInput(e.target.value)} />
							</div>

							<div className="form-group">
								<text>Bio</text>
								<input className="form-control" type="text" name="bioInput" placeholder={bio} value={bioInput} onChange={(e) => setBioInput(e.target.value)} />
							</div>

							<div className="form-group">
								<text>Number</text>
								<input className="form-control" type="text" name="numberInput" placeholder={number} value={numberInput} onChange={(e) => setNumberInput(e.target.value)} />
							</div>
							<br />
							<div class="flexbox-container">
								<div>Gender:</div>
								<div>{genderInput}</div>
								<div style={sectionStyle}>
									<select id="dropdown" onChange={(e) => setGender(e.target.value)}>
										<option value="MALE"> Male</option>
										<option value="FEMALE"> Female</option>
									</select>
								</div>
							</div>
							<br />
							<div className="form-group">
								<input className="btn btn-primary btn-block" type="submit" value="Save" />
							</div>
						</div>
					</form>
				</div>
			</div>
		</React.Fragment>
	);
};

export default EditProfile;
