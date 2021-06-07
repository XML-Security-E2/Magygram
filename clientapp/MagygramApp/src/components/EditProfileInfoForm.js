import React, { useContext, useEffect } from "react";
import { useState } from "react";
import { UserContext } from "../contexts/UserContext";
import { userService } from "../services/UserService";

const EditProfileInfoForm = () => {
	const imgStyle = { transform: "scale(1.5)", width: "100%", position: "absolute", left: "0" };
	const { userState, dispatch } = useContext(UserContext);

	const imgRef = React.createRef();

	const [email, setEmail] = useState(userState.userProfile.user.email);
	const [name, setName] = useState(userState.userProfile.user.name);
	const [surname, setSurname] = useState(userState.userProfile.user.surname);
	const [username, setUsername] = useState(userState.userProfile.user.username);
	const [bio, setBio] = useState(userState.userProfile.user.bio);
	const [image, setImage] = useState("");
	const [showedImage, setShowedImage] = useState(userState.userProfile.user.imageUrl);
	const [selectedImage, setSelectedImage] = useState(false);

	const [website, setWebsite] = useState(userState.userProfile.user.website);
	const [number, setNumber] = useState(userState.userProfile.user.number);
	const [gender, setGender] = useState(userState.userProfile.user.gender);

	const onImageChange = (e) => {
		setImage(e.target.files[0]);
		setSelectedImage(true);

		if (e.target.files && e.target.files[0]) {
			let img = e.target.files[0];
			setShowedImage(URL.createObjectURL(img));
		}
	};

	const handleImageDeselect = () => {
		setImage("");
		setShowedImage("");
		setSelectedImage(false);
	};

	const handleImageChange = () => {
		userService.editUserImage(localStorage.getItem("userId"), image, dispatch);
	};

	useEffect(() => {
		setName(userState.userProfile.user.name);
		setSurname(userState.userProfile.user.surname);
		setUsername(userState.userProfile.user.username);
		setEmail(userState.userProfile.user.email);
		setWebsite(userState.userProfile.user.website);
		setBio(userState.userProfile.user.bio);
		setGender(userState.userProfile.user.gender);
		setNumber(userState.userProfile.user.number);
		setShowedImage(userState.userProfile.user.imageUrl);
	}, [userState.userProfile.user]);

	useEffect(() => {
		const getProfileHandler = async () => {
			await userService.getUserProfileByUserId(localStorage.getItem("userId"), dispatch);
		};
		getProfileHandler();
	}, [dispatch]);

	const handleSubmit = (e) => {
		e.preventDefault();

		let userRequestDTO = {
			name,
			surname,
			username,
			website,
			bio,
			number,
			gender,
		};

		userService.editUser(localStorage.getItem("userId"), userRequestDTO, dispatch);
	};

	return (
		<React.Fragment>
			<form method="post" onSubmit={handleSubmit}>
				<div>
					<br />

					<div className="form-group row">
						<input type="file" ref={imgRef} style={{ display: "none" }} name="image" accept="image/png, image/jpeg" onChange={onImageChange} />

						<div className="rounded-circle  overflow-hidden d-flex justify-content-center align-items-center border border-danger story-profile-photo-visited ml-5">
							<img src={showedImage === "" ? "assets/img/profile.jpg" : showedImage} alt="" style={imgStyle} />
						</div>
						<div className="col-sm-9">
							<div className="row ml-4">
								<h5>{username}</h5>
							</div>
							<button type="button" hidden={selectedImage} className="btn row btn-link btn-fw ml-2" onClick={() => imgRef.current.click()}>
								<b className="text-primary">Change profile picture</b>
							</button>
							<div className="row ml-2">
								<button hidden={!selectedImage} type="button" className="btn btn-link btn-fw ml-2" style={{ color: "gray" }} onClick={handleImageChange}>
									<b>Upload</b>
								</button>
								<button hidden={!selectedImage} type="button" className="btn btn-link btn-fw ml-2" style={{ color: "gray" }} onClick={handleImageDeselect}>
									<b className="text-danger">Discard</b>
								</button>
							</div>
						</div>
					</div>
					<br />
					<div className="form-group row">
						<label for="name" className="col-sm-3 col-form-label">
							<b>Name</b>
						</label>
						<div class="col-sm-9">
							<input required className="form-control" type="text" id="name" name="name" placeholder="Name" value={name} onChange={(e) => setName(e.target.value)} />
						</div>
					</div>
					<div className="form-group row">
						<label for="surname" className="col-sm-3 col-form-label">
							<b>Surname</b>
						</label>
						<div class="col-sm-9">
							<input required className="form-control" type="text" id="surname" name="surname" placeholder="Surname" value={surname} onChange={(e) => setSurname(e.target.value)} />
						</div>
					</div>
					<div className="form-group row">
						<label for="email" className="col-sm-3 col-form-label">
							<b>Email</b>
						</label>
						<div class="col-sm-9">
							<input required disabled className="form-control" id="email" type="email" name="email" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
						</div>
					</div>

					<div className="form-group row">
						<label for="username" className="col-sm-3 col-form-label">
							<b>Username</b>
						</label>
						<div class="col-sm-9">
							<input
								required
								className="form-control"
								type="username"
								id="username"
								name="username"
								placeholder="Username"
								value={username}
								onChange={(e) => setUsername(e.target.value)}
							/>
						</div>
					</div>

					<div className="form-group row">
						<label for="website" className="col-sm-3 col-form-label">
							<b>Website</b>
						</label>
						<div class="col-sm-9">
							<input className="form-control" type="text" id="webiste" name="websiteInput" placeholder="Website" value={website} onChange={(e) => setWebsite(e.target.value)} />
						</div>
					</div>

					<div className="form-group row">
						<label for="bio" className="col-sm-3 col-form-label">
							<b>Bio</b>
						</label>
						<div class="col-sm-9">
							<input className="form-control" type="text" id="bio" name="bioInput" placeholder="Bio" value={bio} onChange={(e) => setBio(e.target.value)} />
						</div>
					</div>

					<div className="form-group row">
						<label for="website" className="col-sm-3 col-form-label">
							<b>Number</b>
						</label>
						<div class="col-sm-9">
							<input className="form-control" type="text" id="website" name="numberInput" placeholder="Number" value={number} onChange={(e) => setNumber(e.target.value)} />
						</div>
					</div>
					<br />
					<div className="form-group row">
						<label for="gender" className="col-sm-3 col-form-label">
							<b>Gender</b>
						</label>
						<div class="col-sm-9">
							<select id="gender" className="form-control" value={gender} onChange={(e) => setGender(e.target.value)}>
								<option value="" disabled>
									Select gender
								</option>
								<option value="MALE"> Male</option>
								<option value="FEMALE"> Female</option>
							</select>
						</div>
					</div>
					<br />
					<div className="form-group">
						<button className="btn btn-primary float-right  mb-2" type="submit">
							Save
						</button>
					</div>
				</div>
			</form>
		</React.Fragment>
	);
};

export default EditProfileInfoForm;
