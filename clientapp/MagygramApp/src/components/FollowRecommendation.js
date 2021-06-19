import React,{useContext} from "react";
import { UserContext } from "../contexts/UserContext";

const FollowRecommendation = () =>{
    const { userState } = useContext(UserContext);
    const imgStyle = { transform: "scale(1.5)", width: "100%", position: "absolute", left: "0" };

    return (
		<React.Fragment>
			

            <div className="col-4">
                <div className="d-flex flex-row align-items-center">
                    <div className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border sidenav-profile-photo">
                        <img src={userState.userProfile.user.imageUrl === "" ? "assets/img/profile.jpg" : userState.userProfile.user.imageUrl} alt="..." style={imgStyle}/>
                    </div>
                    <div className="profile-info ml-3">
                        <span className="profile-info-username">{userState.userProfile.user.username}</span>
                        <span className="profile-info-name">Nikola Kolovic</span>
                    </div>
                </div>

                <div className="mt-4">
                    <div className="d-flex flex-row justify-content-between">
                        <small className="text-muted font-weight-normal">Suggestions For You</small>
                        <small>See All</small>
                    </div>
                </div>  

                <div className="d-flex flex-row justify-content-between align-items-center mt-3 mb-3">
                    <div className="d-flex flex-row align-items-center">
                        <div className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border sugest-profile-photo">
                            <img src="assets/images/profiles/profile-1.jpg" alt="..." style={imgStyle}/>
						</div>
                        <strong className="ml-3 sugest-username">peraperovic</strong>
                    </div>
                    <button className="btn btn-primary btn-sm p-0 btn-ig ">Follow</button>
                </div>
                <div className="d-flex flex-row justify-content-between align-items-center mt-3 mb-3">
                    <div className="d-flex flex-row align-items-center">
                        <div className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border sugest-profile-photo">
                            <img src="assets/images/profiles/profile-2.jpg" alt="..." style={imgStyle}/>
						</div>
                        <strong className="ml-3 sugest-username">djuraaa</strong>
                    </div>
                    <button className="btn btn-primary btn-sm p-0 btn-ig ">Follow</button>
                </div>
                <div className="d-flex flex-row justify-content-between align-items-center mt-3 mb-3">
                    <div className="d-flex flex-row align-items-center">
                        <div className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border sugest-profile-photo">
                            <img src="assets/images/profiles/profile-3.jpg" alt="..." style={imgStyle}/>
						</div>
                        <strong className="ml-3 sugest-username">djurovic</strong>
                    </div>
                    <button className="btn btn-primary btn-sm p-0 btn-ig ">Follow</button>
                </div>
                <div className="d-flex flex-row justify-content-between align-items-center mt-3 mb-3">
                    <div className="d-flex flex-row align-items-center">
                        <div className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border sugest-profile-photo">
                            <img src="assets/images/profiles/profile-6.jpg" alt="..." style={imgStyle}/>
						</div>
                        <strong className="ml-3 sugest-username">username</strong>
                    </div>
                    <button className="btn btn-primary btn-sm p-0 btn-ig ">Follow</button>
                </div>
                <div className="d-flex flex-row justify-content-between align-items-center mt-3 mb-3">
                    <div className="d-flex flex-row align-items-center">
                        <div className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border sugest-profile-photo">
                            <img src="assets/images/profiles/profile-5.jpg" alt="..." style={imgStyle}/>
						</div>
                        <strong className="ml-3 sugest-username">username2</strong>
                    </div>
                    <button className="btn btn-primary btn-sm p-0 btn-ig ">Follow</button>
                </div>
            </div>
            
            

            
                    	
		</React.Fragment>
	);
};

export default FollowRecommendation;
