import React, {useContext} from "react";
import { adminConstants } from "../constants/AdminConstants";
import { AdminContext } from "../contexts/AdminContext";

const AdminHomePageTabs = () => {
	const { state, dispatch } = useContext(AdminContext);

	const handleVerificationRequestClick =() => {
		dispatch({type: adminConstants.SHOW_VERIFICATION_REQUEST_TAB })
	}
	
	const handleContentReportClick =() => {
		dispatch({type: adminConstants.SHOW_CONTENT_REPORTS_TAB })
	}
	
	const handleAgentnRequestClick =() => {
		dispatch({type: adminConstants.SHOW_AGENT_REQUEST_TAB })
	}

	return (
		<React.Fragment>
			<nav className="nav nav-tabs  nav-justified">
				<a href="#" className={state.activeTab.verificationRequestsShow? "nav-item nav-link active":"nav-item nav-link"} onClick={handleVerificationRequestClick}>
					<i className="fa fa-check"></i> Verification requests
				</a>
				<a href="#" className={state.activeTab.contentReportShow? "nav-item nav-link active":"nav-item nav-link"} onClick={handleContentReportClick}>
					<i className="bi bi-file-check"></i> Content reports
				</a>
				<a href="#" className={state.activeTab.agentRequestsShow? "nav-item nav-link active":"nav-item nav-link"} onClick={handleAgentnRequestClick}>
					<i className="icon ion-person"></i> Agent requests
				</a>
			</nav>
		</React.Fragment>

	);
};

export default AdminHomePageTabs;
