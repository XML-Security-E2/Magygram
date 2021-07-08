import React, {useContext, useEffect, useCallback, useState} from "react";
import { UserContext } from "../contexts/UserContext";
import {requestsService} from "../services/RequestsService"
import { confirmAlert } from 'react-confirm-alert';
import 'react-confirm-alert/src/react-confirm-alert.css'
import Axios from "axios";
import { authHeader } from "../helpers/auth-header";
import { userService } from "../services/UserService";
import StorySliderForAdmin from "./modals/StorySliderForAdmin";
import ViewPostModal from "./modals/ViewPostModal";
import PostContextProvider from "../contexts/PostContext";
import StoryContextProvider from "../contexts/StoryContext";
import ReportedPost from "./ReportedPost";
import ReportedStory from "./ReportedStoryInfluencer";
import { postService } from "../services/PostService";

const InfluencerCampagns = ({show}) => {

	const { userState, dispatch } = useContext(UserContext);

    useEffect(() => {
		var username = localStorage.getItem('username');
		console.log(localStorage.getItem('username'))
		const getReportRequestsHandler = async () => {
			await userService.getCampaignsForUser(username, dispatch)
		};
		getReportRequestsHandler();
	}, [dispatch]);

	
	const deleteVerificationRequest = (requestId)=>{
		confirmAlert({
			message: 'Are you sure to do this?',
			buttons: [
			  {
				label: 'Yes',
				onClick: () => {
                    Axios.put(`/api/requests/campaign/${requestId}/delete`, {}, { validateStatus: () => true, headers: authHeader() })
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
	const acceptVerificationRequest = (requestId, id,agentId)=>{
		
			

			let influencerPostDTO = {
				postIdInfl: requestId,
				userId: localStorage.getItem("userId"),
				username: localStorage.getItem("username"),
			};
			Axios.put(`/api/posts/campaign/influencer`, influencerPostDTO, { validateStatus: () => true, headers: authHeader() })
				.then((res) => {
					console.log(res.data)
					let influencerDTO = {
						contentId: res.data,
						ownerId: agentId,
						campaignType: "POST",
						userId: localStorage.getItem("userId"),
						username: localStorage.getItem("username"),
					};
					Axios.post(`/api/ads/campaign/create/influencer`,influencerDTO, { validateStatus: () => true, headers: authHeader() })
						.then((res) => {
							console.log(res)
							 Axios.put(`/api/requests/campaign/${id}/delete`, {}, { validateStatus: () => true, headers: authHeader() })
								.then((res) => {
									console.log(res);
									if (res.status === 200) {
										window.location = "#/profile?userId=" + localStorage.getItem("userId");
									} else {
										console.log("ERROR")
									}
								})
								.catch((err) => {
									console.log("ERROR")
								});
						})
						.catch((err) => {
							console.log(err);
						});
					
				})
				.catch((err) => {
					console.log(err);
				});
	

	}
	const acceptVerificationRequestStory = (requestId,id,agentId)=>{
		
		let influencerPostDTO = {
			postIdInfl: requestId,
			userId: localStorage.getItem("userId"),
			username: localStorage.getItem("username"),
		};
		Axios.put(`/api/story/campaign/influencer`, influencerPostDTO, { validateStatus: () => true, headers: authHeader() })
			.then((res) => {
				console.log(res)
				let influencerDTO = {
					contentId: res.data,
					ownerId: agentId,
					campaignType: "STORY",
					userId: localStorage.getItem("userId"),
					username: localStorage.getItem("username"),
				};
		
				Axios.post(`/api/ads/campaign/create/influencer`,influencerDTO, { validateStatus: () => true, headers: authHeader() })
					.then((res) => {
						console.log(res)
							 Axios.put(`/api/requests/campaign/${id}/delete`, {}, { validateStatus: () => true, headers: authHeader() })
								.then((res) => {
									console.log(res);
									if (res.status === 200) {
										window.location = "/#";
									} else {
										console.log("ERROR")
									}
								})
								.catch((err) => {
									console.log("ERROR")
								});
						
					})
					.catch((err) => {
						console.log(err);
					});
				
			})
			.catch((err) => {
				console.log(err);
			});

	}

	return (
		<React.Fragment>
		<PostContextProvider>
			<StoryContextProvider>
			<div hidden={!show} >
            {userState.campaigns!==null ?
					<table hidden={userState.campaigns===null} className="table mt-5" style={{width:"100%"}}>
						<tbody>
							{userState.campaigns.map(request => 
								<tr id={request.Id} key={request.Id} >
									<td >
										<div><b>Price:</b> {request.Price}</div>
									</td>
									<td className="text-center">
                                        <div  hidden={(request.ContentType === "POST")}>
                                            <ReportedStory  storyId={request.ContentId} />
                                        </div>
                                        <div hidden={ (request.ContentType === "STORY")} >
                                            <ReportedPost  id={request.ContentId}/>
                                        </div>
                                	</td>
									<td className="text-right">
										<div className="mt-2">
											<button style={{height:'40px'},{verticalAlign:'center'}} className="btn btn-outline-secondary" type="button" onClick={()=>deleteVerificationRequest(request.Id)}><i className="icofont-subscribe mr-1"></i>Delete request</button>
										                                  
										</div>

										<div hidden={(request.ContentType === "STORY")} className="mt-2">
											<button style={{height:'40px'},{verticalAlign:'center'}} className="btn btn-outline-secondary" type="button" onClick={()=>acceptVerificationRequest(request.ContentId,request.Id, request.AgentId)}><i className="icofont-subscribe mr-1"></i>Accept post</button>
										                                  
										</div>
										<div hidden={(request.ContentType === "POST")} className="mt-2">
											<button style={{height:'40px'},{verticalAlign:'center'}} className="btn btn-outline-secondary" type="button" onClick={()=>acceptVerificationRequestStory(request.ContentId,request.Id,  request.AgentId)}><i className="icofont-subscribe mr-1"></i>Accept story</button>
										                                  
										</div>
									</td>
								</tr>
							)}		
                            					
						</tbody>
					</table> : 
					<div>
						<div className="col-12 mt-5 d-flex justify-content-center text-secondary" >
							<h3 hidden={localStorage.getItem('category') !== "INFLUENCER" }>Campaign request not exist</h3>
						</div>
						<div className="col-12 mt-5 d-flex justify-content-center text-secondary" >
							<h3 hidden={localStorage.getItem('category') === "INFLUENCER" }> You need to be influencer for this option!</h3>
						</div>
					</div>
				}
			</div>
            	
			<StorySliderForAdmin/>
			<ViewPostModal/>	
			</StoryContextProvider>
			</PostContextProvider>
		</React.Fragment>

	);
};

export default InfluencerCampagns;
