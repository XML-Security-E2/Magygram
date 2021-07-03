import React, { useContext, useEffect } from "react";
import { UserContext } from "../../contexts/UserContext";
import { userService } from "../../services/UserService";

const CampaignTokenGenerator = () => {
	const {userState, dispatch } = useContext(UserContext);

	useEffect(() => {
		userService.getCampaignAPITokenForAgent(dispatch);
	}, []);

	const generateNewToken = () => {
		userService.generateNewToken(dispatch);
	}

	const deleteToken = () => {
		userService.deleteToken(dispatch);
	}

	return (
		<React.Fragment>
			<div className="form-group row">
                <button className="btn btn-primary float-right" onClick={() => generateNewToken()}>
                    Generate new token
                </button>
				<button className="btn btn-danger float-right ml-3" onClick={() => deleteToken()}>
                    Delete token
                </button>
			</div>
			<div className="form-group row">
				<label for="apitoken" className="col-form-label">
					<b>API Token</b>
				</label>
			</div>
			<div className="form-group row">
			<textarea
          		value={userState.agentCampaignAPITOken}
		  		readOnly
				rows="10"
				cols="40"
        		/>
			</div>
			<div className="form-group row">
                <button className="btn btn-secondary float-right" onClick={() => {navigator.clipboard.writeText(userState.agentCampaignAPITOken)}}>
                    Copy to clipboard
                </button>
			</div>
		</React.Fragment>
	);
};

export default CampaignTokenGenerator;
