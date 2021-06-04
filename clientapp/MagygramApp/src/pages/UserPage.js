import React, {useState, useEffect} from "react";
import Axios from "axios";
import { authHeader } from "../helpers/auth-header";

const UserPage = (props) => {

    const userId = props.match.params.id;
	const [username, setUsername] = useState("");
	const [logged, setLogged] = useState();

	useEffect(() => {
		Axios.get(`https://localhost:460/api/users/logged`, { validateStatus: () => true, headers: authHeader() })
			.then((res) => {
				console.log(res.data);
                setLogged(res.data.id);
			})
			.catch((err) => {console.log(err);});

        Axios.get(`https://localhost:460/api/users/` + userId, { validateStatus: () => true, headers: authHeader() })
			.then((res) => {
				if (res.status === 200) {
					setUsername(res.data.Username);
				}
			})
			.catch((err) => {console.log(err);});
    
      });

	const handleFollowRequest = () => {

		var followRequest = {subjectId: logged, objectId: userId}

		Axios.post(`https://localhost:463/api/relationship/follow`, followRequest, { validateStatus: () => true })
		.then((res) => {

		})
		.catch((err) => {
			console.log(err);
		});
	};

	return (
		<React.Fragment>
            <label>{username}</label>
			<button onClick={handleFollowRequest}>Follow</button>
		</React.Fragment>
	);
};

export default UserPage