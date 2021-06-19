import React, {useContext} from "react";
import { UserContext } from "../contexts/UserContext";
import { userService } from "../services/UserService";

const RecommendedUser = ({userInfo}) =>{
    const { dispatch } = useContext(UserContext);
    const imgStyle = { transform: "scale(1.5)", width: "100%", position: "absolute", left: "0" };

    const handleClickOnProfile = (userId) => {
        window.location = "#/profile?userId="+userId;
    }

    const handleClickFollow = (userId) => {
        userService.followRecommendedUser(userId,dispatch)
    }

    return (
		<React.Fragment>
            <div className="d-flex flex-row justify-content-between align-items-center mt-3 mb-3">
                <div className="d-flex flex-row align-items-center">
                    <div onClick={() => handleClickOnProfile(userInfo.Id)} className="rounded-circle overflow-hidden d-flex justify-content-center align-items-center border sugest-profile-photo">
                        <img src={userInfo.ImageURL === "" ? "assets/img/profile.jpg" : userInfo.ImageURL} alt="..." style={imgStyle}/>
                    </div>
                    <strong onClick={() => handleClickOnProfile(userInfo.Id)} className="ml-3 sugest-username">{userInfo.Username}</strong>
                </div>
                {userInfo.SendedRequest  ? 
                    <div style={{color:"green"}}>Request sended</div> 
                    : 
                    <div>
                        {userInfo.Followed? 
                        <div style={{color:"green"}}>Followed</div>:                     
                        <button onClick ={() => handleClickFollow(userInfo.Id)} className="btn btn-primary btn-sm p-0 btn-ig">Follow</button>
                        }
                    </div>
                }
            </div>      	
		</React.Fragment>
	);
};

export default RecommendedUser;
