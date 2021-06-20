import React, {useContext} from "react";
import { UserContext } from "../contexts/UserContext";

const RegistrationFormTabs = ({onUserRegistrationClick,onAgentRegistrationClick}) => {
	const { userState } = useContext(UserContext);

	return (
		<React.Fragment>
			<ul class="nav nav-tabs">
				<li class="nav-item">
					<a className={userState.registrationTab.showUserRegistrationTab? "nav-item nav-link active":"nav-item nav-link"} onClick={onUserRegistrationClick}>User registration</a>
				</li>
				<li class="nav-item">
					<a className={userState.registrationTab.showAgentRegistrationTab? "nav-item nav-link active":"nav-item nav-link"} onClick={onAgentRegistrationClick}>Agent registration</a>
				</li>
			</ul>
		</React.Fragment>

	);
};

export default RegistrationFormTabs;
