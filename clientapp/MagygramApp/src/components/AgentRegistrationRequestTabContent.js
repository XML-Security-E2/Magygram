import React, {useContext, useEffect, useCallback, useState} from "react";
import { AdminContext } from "../contexts/AdminContext";
import {requestsService} from "../services/RequestsService"
import { confirmAlert } from 'react-confirm-alert';
import 'react-confirm-alert/src/react-confirm-alert.css'
import { adminConstants } from "../constants/AdminConstants";
import FailureAlert from "./FailureAlert";
import SuccessAlert from "./SuccessAlert";
import AddNewAgentModal from "./modals/AddNewAgentModal";
import { modalConstants } from "../constants/ModalConstants";

const AgentRegistrationRequestTabContent = () => {
	const { state,dispatch } = useContext(AdminContext);

	const approveVerificationRequest = (requestId)=>{
		confirmAlert({
			message: 'Are you sure to do this?',
			buttons: [
			  {
				label: 'Yes',
				onClick: () => requestsService.approveAgentRegistrationRequest(requestId,dispatch)
			  },
			  {
				label: 'No',
			  }
			]
		  });
	}

	const rejectVerificationRequest = (requestId)=>{
		confirmAlert({
			message: 'Are you sure to do this?',
			buttons: [
			  {
				label: 'Yes',
				onClick: () => requestsService.rejectAgentRegistrationRequest(requestId,dispatch)
			  },
			  {
				label: 'No',
			  }
			]
		  });
	}

	const addNewAgentHandler = () => {
		dispatch({ type: modalConstants.SHOW_REGISTER_AGENT_MODAL });
	}

	useEffect(() => {
		const getAgentRegistrationRequestsHandler = async () => {
			await requestsService.getAllPendingAgentRegistrationRequest(dispatch)
		};
		getAgentRegistrationRequestsHandler();
	}, [dispatch]);

	return (
		<React.Fragment>
			<div hidden={!state.activeTab.agentRequestsShow}>
				<div className="mt-3">
					<SuccessAlert
						hidden={!state.agentRegistrationRequests.showSuccessMessage}
						header="Success"
						message={state.agentRegistrationRequests.successMessage}
						handleCloseAlert={() => dispatch({ type: adminConstants.APPROVE_AGENT_REGISTRATION_REQUEST_REQUEST })}
					/>
					<SuccessAlert
						hidden={!state.registerNewAgent.showSuccessMessage}
						header="Success"
						message={state.registerNewAgent.successMessage}
						handleCloseAlert={() => dispatch({ type: adminConstants.APPROVE_AGENT_REGISTRATION_REQUEST_REQUEST })}
					/>
					<FailureAlert
						hidden={!state.agentRegistrationRequests.showError}
						header="Error"
						message={state.agentRegistrationRequests.errorMessage}
						handleCloseAlert={() => dispatch({ type: adminConstants.APPROVE_AGENT_REGISTRATION_REQUEST_REQUEST })}
					/>
				</div>
				<div className="d-flex flex-column mt-4 mb-4">
					<div className="card">
						<div type="button" className="btn btn-link btn-fw text-secondary w-100 border-0" onClick={addNewAgentHandler}>
							Add new agent
						</div>
					</div>
				</div>
				{state.agentRegistrationRequests.requests!==null ?
					<table hidden={state.agentRegistrationRequests.requests===null} className="table mt-5" style={{width:"100%"}}>
						<tbody>
							{state.agentRegistrationRequests.requests.map(request => 
								<tr id={request.Id} key={request.Id} >
									<td >
										<div><b>Name:</b> {request.Name}</div>
										<div><b>Surname:</b> {request.Surname}</div>
										<div><b>Username:</b> {request.Username}</div>
                                        <div><b>Website:</b> <a href={"https://"+ request.WebSite} target="_blank">{request.WebSite}</a></div>

									</td>
									<td className="text-right">
										<div className="mt-2">
											<button style={{height:'40px'},{verticalAlign:'center'}} className="btn btn-outline-secondary" type="button" onClick={()=>approveVerificationRequest(request.Id)}><i className="icofont-subscribe mr-1"></i>Accept</button>
											<br></br>
											<button style={{height:'30px'},{verticalAlign:'center'},{marginTop:'2%'}} className="btn btn-outline-secondary"type="button" onClick={()=>rejectVerificationRequest(request.Id)}><i className="icofont-subscribe mr-1"></i>Reject</button>
											<br></br>                                       
										</div>
									</td>
								</tr>
							)}							
						</tbody>
					</table> : 
					<div>
						<div className="col-12 mt-5 d-flex justify-content-center text-secondary" >
							<h3>Agent registration request not exist</h3>
						</div>
					</div>
				}
			</div>
            <AddNewAgentModal/>
		</React.Fragment>

	);
};

export default AgentRegistrationRequestTabContent;
