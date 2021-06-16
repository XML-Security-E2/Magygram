import React, {useContext} from "react";
import { AdminContext } from "../contexts/AdminContext";

const AdminVerificationRequestTabContent = () => {
	const { state } = useContext(AdminContext);

	return (
		<React.Fragment>
            <table hidden={!state.activeTab.verificationRequestsShow} className="table mt-5" style={{width:"100%"}}>
				<tbody>
					<tr id={123} key={123} >
						<td >
							<div><b>Name:</b> Pera</div>
							<div><b>Surname:</b> Perovic</div>
							<div><b>Category:</b> Sport</div>
						</td>
						<td className="text-center">
                            <div>
                                <a href="#" class="link-primary">Visit profile</a>
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
								<button style={{height:'40px'},{verticalAlign:'center'}} className="btn btn-outline-secondary" type="button"><i className="icofont-subscribe mr-1"></i>Accept</button>
								<br></br>
								<button style={{height:'30px'},{verticalAlign:'center'},{marginTop:'2%'}} className="btn btn-outline-secondary"type="button"><i className="icofont-subscribe mr-1"></i>Reject</button>
								<br></br>                                       
							</div>
						</td>
					</tr>
					<tr id={123} key={123} >
						<td >
							<div><b>Name:</b> Pera</div>
							<div><b>Surname:</b> Perovic</div>
							<div><b>Category:</b> Sport</div>
						</td>
						<td className="text-center">
                            <div>
                                <a href="#" class="link-primary">Visit profile</a>
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
								<button style={{height:'40px'},{verticalAlign:'center'}} className="btn btn-outline-secondary" type="button"><i className="icofont-subscribe mr-1"></i>Accept</button>
								<br></br>
								<button style={{height:'30px'},{verticalAlign:'center'},{marginTop:'2%'}} className="btn btn-outline-secondary"type="button"><i className="icofont-subscribe mr-1"></i>Reject</button>
								<br></br>                                       
							</div>
						</td>
					</tr>
                    <tr id={123} key={123} >
						<td >
							<div><b>Name:</b> Pera</div>
							<div><b>Surname:</b> Perovic</div>
							<div><b>Category:</b> Sport</div>
						</td>
						<td className="text-center">
                            <div>
                                <a href="#" class="link-primary">Visit profile</a>
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
								<button style={{height:'40px'},{verticalAlign:'center'}} className="btn btn-outline-secondary" type="button"><i className="icofont-subscribe mr-1"></i>Accept</button>
								<br></br>
								<button style={{height:'30px'},{verticalAlign:'center'},{marginTop:'2%'}} className="btn btn-outline-secondary"type="button"><i className="icofont-subscribe mr-1"></i>Reject</button>
								<br></br>                                       
							</div>
						</td>
					</tr>								
				</tbody>
			</table>
		</React.Fragment>

	);
};

export default AdminVerificationRequestTabContent;
