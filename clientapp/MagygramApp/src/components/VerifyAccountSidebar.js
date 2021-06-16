const VerifyAccoundSidebar = ({show,handleVerifyAccount}) => {
	return (
		<nav style={ show ? {backgroundColor: "lightgray"}:{backgroundColor: "white"}} className="nav flex-column">
			<button type="button" className="btn btn-link btn-fw nav-link active" onClick={handleVerifyAccount}>
				Verify account
			</button>
		</nav>
	);
};

export default VerifyAccoundSidebar;
