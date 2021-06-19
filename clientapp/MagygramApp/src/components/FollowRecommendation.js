import React,{useContext, useEffect} from "react";
import { UserContext } from "../contexts/UserContext";
import { userService } from "../services/UserService";
import RecommendedUser from "../components/RecommendedUser"

const FollowRecommendation = () =>{
    const { userState, dispatch } = useContext(UserContext);
    const imgStyle = { transform: "scale(1.5)", width: "100%", position: "absolute", left: "0" };

    useEffect(() => {
		const getFollowRecommendationHandler = async () => {
			await userService.getFollowRecommendationHandler(dispatch);
		};
		getFollowRecommendationHandler();
	}, [dispatch]);


    const handleClickOnProfile = () => {
        window.location = "#/profile?userId="+localStorage.getItem("userId");
    }

    const handleSeeAll = () => {
        window.location = "#/recommended-users"
    }

    return (
		<React.Fragment>
			

            <div className="col-4">
                <div className="d-flex flex-row align-items-center">
                    <div className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border sidenav-profile-photo">
                        <img onClick={() => handleClickOnProfile()} src={userState.followRecommendationInfo.imageUrl === "" ? "assets/img/profile.jpg" : userState.followRecommendationInfo.imageUrl} alt="..." style={imgStyle}/>
                    </div>
                    <div className="profile-info ml-3">
                        <span onClick={() => handleClickOnProfile()} className="profile-info-username">{userState.followRecommendationInfo.username}</span>
                        <span className="profile-info-name">{userState.followRecommendationInfo.name} {userState.followRecommendationInfo.surname}</span>
                    </div>
                </div>

                <div className="mt-4">
                    <div className="d-flex flex-row justify-content-between">
                        <small className="text-muted font-weight-normal">Suggestions For You</small>
                        <small onClick={() => handleSeeAll()}>See All</small>
                    </div>
                </div>  

                {userState.followRecommendationInfo.recommendUserInfo.slice(0,5).map((recommendedUser) => {
				    return (
                        <RecommendedUser userInfo= {recommendedUser}/>
                    );
			    })}
            </div>
            
            

            
                    	
		</React.Fragment>
	);
};

export default FollowRecommendation;
