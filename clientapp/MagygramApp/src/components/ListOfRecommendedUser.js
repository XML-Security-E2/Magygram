import React, {useContext, useEffect} from "react";
import { UserContext } from "../contexts/UserContext";
import RecommendedUser from "../components/RecommendedUser"
import { userService } from "../services/UserService";

const ListOfRecommendedUser = () => {
    const { userState, dispatch } = useContext(UserContext);

    useEffect(() => {
		const getFollowRecommendationHandler = async () => {
			await userService.getFollowRecommendationHandler(dispatch);
		};
		getFollowRecommendationHandler();
	}, [dispatch]);

	return (
		<React.Fragment> 
            <div>
                {userState.followRecommendationInfo.recommendUserInfo.slice(0,20).map((recommendedUser) => {
                    return (
                        <RecommendedUser userInfo= {recommendedUser}/>);
                })}
            </div>
        
		</React.Fragment>
	);
};

export default ListOfRecommendedUser;
