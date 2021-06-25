import React, {useContext, useEffect, useCallback, useState} from "react";
import { AdminContext } from "../contexts/AdminContext";
import {requestsService} from "../services/RequestsService"
import { confirmAlert } from 'react-confirm-alert';
import { PostContext } from "../contexts/PostContext";
import { hasPermissions } from "../helpers/auth-header";
import { postService } from "../services/PostService";
import { adminConstants } from "../constants/AdminConstants";
import 'react-confirm-alert/src/react-confirm-alert.css'
import FailureAlert from "./FailureAlert";
import SuccessAlert from "./SuccessAlert";
import ViewPostModal from "./modals/ViewPostModal";
import Axios from "axios";
import { authHeader } from "../helpers/auth-header";
import ReportedPost from "./ReportedPost";
import ReportedStory from "./ReportedStory";

const AdminReportRequestTab = () => {
	const { state,dispatch } = useContext(AdminContext);
	const { statePost,dispatchPost } = useContext(PostContext);

	const handleVisitProfile = (userId) => {
		window.location = "#/profile?userId=" + userId;
	}

	const approveVerificationRequest = (requestId)=>{
		confirmAlert({
			message: 'Are you sure to do this?',
			buttons: [
			  {
				label: 'Yes',
				onClick: () => {

                    Axios.put(`/api/report/${requestId}/delete`, {}, { validateStatus: () => true, headers: authHeader() })
                    .then((res) => {
                        console.log(res);
                        if (res.status === 200) {
                            console.log("Report request has been deleted");
                            window.location.reload()
                        } else {
                            console.log("ERROR")
                        }
                    })
                    .catch((err) => {
                        console.log("ERROR")
                    });
            
                }
			  },
			  {
				label: 'No',
			  }
			]
		  });
	}

    
	const handleViewStory = (requestId)=>{

    }
    
	const handleViewPost = async (postId) => {
	
		await postService.findPostById(postId, dispatchPost);
	};

    useEffect(() => {
		const getReportRequestsHandler = async () => {
			await requestsService.getAllReportRequest(dispatch)
		};
		getReportRequestsHandler();
	}, [dispatch]);

	return (
		<React.Fragment>
			<div hidden={!state.activeTab.contentReportShow}>
                 <div className="mt-3">
					<SuccessAlert
						hidden={!state.verificationRequests.showSuccessMessage}
						header="Success"
						message={state.verificationRequests.successMessage}
						handleCloseAlert={() => dispatch({ type: adminConstants.APPROVE_VERIFICATION_REQUEST_REQUEST })}
					/>
					<FailureAlert
						hidden={!state.verificationRequests.showError}
						header="Error"
						message={state.verificationRequests.errorMessage}
						handleCloseAlert={() => dispatch({ type: adminConstants.APPROVE_VERIFICATION_REQUEST_REQUEST })}
					/>
				</div>
                {state.reportRequests.requests!==null ?
					<table hidden={state.reportRequests.requests===null} className="table mt-5" style={{width:"100%"}}>
						<tbody>
							{state.reportRequests.requests.map(request => 
								<tr id={request.Id} key={request.Id} >
									<td >
										<div><b>REPORT TYPE:</b> {request.ContentType}</div>
									</td>
									<td className="text-center">
										<div>
											<a hidden={(request.ContentType === "POST") || (request.ContentType === "STORY")} onClick={() => handleVisitProfile(request.ContentId)} class="link-primary">Visit profile</a>
										</div>
                                        <div  hidden={(request.ContentType === "USER") || (request.ContentType === "POST")}>
                                            <ReportedStory  storyId={request.ContentId}/>
                                        </div>
                                        <div hidden={(request.ContentType === "USER") || (request.ContentType === "STORY")} >
                                            <ReportedPost  id={request.ContentId}/>
                                        </div>
                                	</td>
									<td className="text-right">
										<div className="mt-2">
											<button style={{height:'40px'},{verticalAlign:'center'}} className="btn btn-outline-secondary" type="button" onClick={()=>approveVerificationRequest(request.Id)}><i className="icofont-subscribe mr-1"></i>Delete report</button>
										                                  
										</div>
									</td>
								</tr>
							)}		
                            
                			<ViewPostModal/>					
						</tbody>
					</table> : 
					<div>
						<div className="col-12 mt-5 d-flex justify-content-center text-secondary" >
							<h3>Report request not exist</h3>
						</div>
					</div>
				}
				
			</div>
            
		</React.Fragment>

	);
};

export default AdminReportRequestTab;
