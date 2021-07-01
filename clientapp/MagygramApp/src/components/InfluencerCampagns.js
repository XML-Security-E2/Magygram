import React, {useContext, useEffect, useCallback, useState} from "react";
import { UserContext } from "../contexts/UserContext";
import {requestsService} from "../services/RequestsService"
import { confirmAlert } from 'react-confirm-alert';
import 'react-confirm-alert/src/react-confirm-alert.css'
import Axios from "axios";
import { authHeader } from "../helpers/auth-header";
import { userService } from "../services/UserService";

const InfluencerCampagns = ({show}) => {

	const { userState, dispatch } = useContext(UserContext);

    useEffect(() => {
		const getReportRequestsHandler = async () => {
			await userService.getCampaignsForUser(dispatch)
		};
		getReportRequestsHandler();
	}, [dispatch]);

	return (
		<React.Fragment>
			<div hidden={!show} >
            {userState.campaigns!==null ?
					<table hidden={userState.campaigns===null} className="table mt-5" style={{width:"100%"}}>
						<tbody>
							{userState.campaigns.map(request => 
								<tr id={request.Id} key={request.Id} >
									<td >
										<div><b>Username:</b> {request.Influencer}</div>
									</td>
                                    <td >
										<div><b>Price:</b> {request.Price}</div>
									</td>
									
								</tr>
							)}		
                            				
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

export default InfluencerCampagns;
