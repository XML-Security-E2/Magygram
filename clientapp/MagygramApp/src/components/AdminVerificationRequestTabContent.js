import React, {useContext, useEffect} from "react";
import { AdminContext } from "../contexts/AdminContext";
import {requestsService} from "../services/RequestsService"

const AdminVerificationRequestTabContent = () => {
	const { state,dispatch } = useContext(AdminContext);

	const handleVisitProfile = (userId) => {
		window.location = "#/profile/" + userId;
	}

	const approveVerificationRequest = (requestId)=>{
		requestsService.approveVerificationRequest(requestId,dispatch)
	}

	const rejectVerificationRequest = (requestId)=>{
		requestsService.rejectVerificationRequest(requestId,dispatch)
	}

	useEffect(() => {
		const getVerificationRequestsHandler = async () => {
			await requestsService.getAllPendingVerificationRequest(dispatch)
		};
		getVerificationRequestsHandler();
	}, [dispatch]);

	return (
		<React.Fragment>
			<div hidden={!state.activeTab.verificationRequestsShow}>
				<table hidden={state.verificationRequests.requests.length===0} className="table mt-5" style={{width:"100%"}}>
					<tbody>
						{state.verificationRequests.requests.map(request => 
							<tr id={request.Id} key={request.Id} >
								<td >
									<div><b>Name:</b> {request.Name}</div>
									<div><b>Surname:</b> {request.Surname}</div>
									<div><b>Category:</b> {request.Category}</div>
								</td>
								<td className="text-center">
									<div>
										<a onClick={() => handleVisitProfile(request.UserId)} class="link-primary">Visit profile</a>
									</div>
									<div>
										<a href="#" class="link-primary">View document</a>
									</div>
									<div>
										<a href="#" class="link-primary">Download document</a>
									</div>
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
				</table>
				<div hidden={state.verificationRequests.requests.length!==0}>
					<div className="col-12 mt-5 d-flex justify-content-center text-secondary" >
						<h3>Verification request not exist</h3>
					</div>
				</div>
			</div>
            
		</React.Fragment>

	);
};

export default AdminVerificationRequestTabContent;
