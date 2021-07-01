const EditCampaignsSidebar = ({ show, handleInfluencerCampaigns }) => {
	return (
		<nav style={show ? { backgroundColor: "lightgray" } : { backgroundColor: "white" }} className="nav flex-column">
			<button type="button" className="btn btn-link btn-fw nav-link active" onClick={handleInfluencerCampaigns}>
				Campaigns
			</button>
		</nav>
	);
};

export default EditCampaignsSidebar;
