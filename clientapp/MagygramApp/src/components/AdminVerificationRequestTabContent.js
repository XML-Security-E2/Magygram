import React, {useContext, useEffect, useCallback, useState} from "react";
import { AdminContext } from "../contexts/AdminContext";
import {requestsService} from "../services/RequestsService"
import { confirmAlert } from 'react-confirm-alert';
import 'react-confirm-alert/src/react-confirm-alert.css'
import ImageViewer from 'react-simple-image-viewer';

const AdminVerificationRequestTabContent = () => {
	const { state,dispatch } = useContext(AdminContext);
	const [isViewerOpen, setIsViewerOpen] = useState(false);
	const [images, setImages]= useState("")

	const handleVisitProfile = (userId) => {
		window.location = "#/profile/" + userId;
	}

	const openImageViewer = useCallback(() => {
		setIsViewerOpen(true);
		setImages(["https://localhost:8085/api/media/5ff2ecfa-85f8-4ca9-8841-47097e1a73b5.jpg"])
	  }, []);

	const closeImageViewer = (image) => {
		setImages([image]);
		setIsViewerOpen(false);
	};

	const handleDownload = async (event,request) => {
		event.preventDefault();
		const response = await fetch(
			request.Document
		);
		if (response.status === 200) {
		  const blob = await response.blob();
		  const url = URL.createObjectURL(blob);
		  const link = document.createElement("a");
		  link.href = url;
		  link.download = request.Name+" "+request.Surname;
		  document.body.appendChild(link);
		  link.click();
		  link.remove();
		  return { success: true };
		}
	}

	const approveVerificationRequest = (requestId)=>{
		confirmAlert({
			message: 'Are you sure to do this?',
			buttons: [
			  {
				label: 'Yes',
				onClick: () => requestsService.rejectVerificationRequest(requestId,dispatch)
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
				onClick: () => requestsService.rejectVerificationRequest(requestId,dispatch)
			  },
			  {
				label: 'No',
			  }
			]
		  });
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
										<a href="#" class="link-primary" onClick={() => openImageViewer(request.Document)}>View document</a>
									</div>
									<div>
										<a href="#" onClick={(event)=>handleDownload(event,request)}  class="link-primary">Download document</a>
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
				{isViewerOpen && (
					<ImageViewer
					src={ images }
					onClose={ closeImageViewer }
					/>
				)}
			</div>
            
		</React.Fragment>

	);
};

export default AdminVerificationRequestTabContent;
