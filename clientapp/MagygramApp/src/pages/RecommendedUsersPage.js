import React, {useContext} from "react";
import HeaderWrapper from "../components/HeaderWrapper";
import UserContextProvider, { UserContext } from "../contexts/UserContext";
import ListOfRecommendedUser from "../components/ListOfRecommendedUser";

const RecommendedUsersPage = () => {

	return (
		<React.Fragment>
            <UserContextProvider>
                <div className="container-scroller">
                    <div className="container-fluid ">
                        <HeaderWrapper />
                        <div className="main-panel">
                            <div className="container">
                                <div class="mt-4">
                                    <div class="container d-flex justify-content-center">
                                        <div class="col-7">
											<div class="row">
												<div class="col-10">
                                                    <label className="mt-4">Suggested</label>
                                                    <ListOfRecommendedUser/>
												</div>
											</div>
										</div> 
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </UserContextProvider>
		</React.Fragment>
	);
};

export default RecommendedUsersPage;
