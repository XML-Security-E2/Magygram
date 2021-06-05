import React, {useState, useEffect} from "react";
import Axios from "axios";
import { authHeader } from "../helpers/auth-header";

const UserPage = (props) => {

    const userId = props.match.params.id;
	const [username, setUsername] = useState("");
	const [logged, setLogged] = useState();
	const [isFollowed, setIsFollowed] = useState(false);

	useEffect(() => {
		Axios.get(`/api/users/logged`, { validateStatus: () => true, headers: authHeader() })
			.then((res) => {
                setLogged(res.data.id);

				var followRequest = {subjectId: res.data.id, objectId: userId}
				console.log(followRequest);
				Axios.post(`https://localhost:463/api/relationship/is-user-followed`, followRequest, { validateStatus: () => true, headers: authHeader() })
					.then((res) => {
						if (res.status === 200) {
							setIsFollowed(res.data);
						}
					})
				.catch((err) => {console.log(err);});
			})
			.catch((err) => {console.log(err);});

        Axios.get(`/api/users/` + userId, { validateStatus: () => true, headers: authHeader() })
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
			setIsFollowed(true);
		})
		.catch((err) => {
			console.log(err);
		});
	};

	return (
		<React.Fragment>
            <label>{username}</label>
			<button disabled={isFollowed} onClick={handleFollowRequest}>Follow</button>
		</React.Fragment>
	);
};

export default UserPage