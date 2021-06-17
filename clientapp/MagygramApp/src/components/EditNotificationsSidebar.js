const EditNotificationsSidebar = ({ show, handleEditNotifications }) => {
	return (
		<nav style={show ? { backgroundColor: "lightgray" } : { backgroundColor: "white" }} className="nav flex-column">
			<button type="button" className="btn btn-link btn-fw nav-link active" onClick={handleEditNotifications}>
				Notifications
			</button>
		</nav>
	);
};

export default EditNotificationsSidebar;
