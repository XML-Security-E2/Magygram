const EditPrivacySettingsSidebar = ({show, handleEditPrivacySettings}) => {
    return (
        <nav style={show ? { backgroundColor: "lightgray" } : { backgroundColor: "white" }} className="nav flex-column">
			<button type="button" className="btn btn-link btn-fw nav-link active" onClick={handleEditPrivacySettings}>
				Privacy settings
			</button>
		</nav>
    );
};

export default EditPrivacySettingsSidebar;