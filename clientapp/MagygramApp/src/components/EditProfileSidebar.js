const EditProfileSidebar = ({show, handleEditProfile}) => {
	return (
		<nav style={show? {backgroundColor: "lightgray"}:{backgroundColor: "white"}} className="nav flex-column mt-3">
			<button type="button" className="btn btn-link btn-fw nav-link active" onClick={handleEditProfile}>
				Edit profile
			</button>
		</nav>
	);
};

export default EditProfileSidebar;
