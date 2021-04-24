import React, { useContext } from "react";
import { userConstants } from "../constants/UserConstants";
import { UserContext } from "../contexts/UserContext";


const UserActivateRequestPage = (props) => {

    const email = props.match.params.id;
    const { dispatch } = useContext(UserContext);
    const handleSubmit = (e) => {
        e.preventDefault();
        let resendActivationLink = {
			email,
		};

        dispatch({ type: userConstants.RESEND_ACTIVATION_LINK_REQUEST, resendActivationLink });
    };
	return (
            <React.Fragment>
            <section className="login-clean">
                <div className="illustration">
                    <i className="icon ion-ios-navigate"></i>
                </div>
                <form method="post" onSubmit={handleSubmit}>
                    <div>Vas nalog je privremeno blokiran. Ukoliko zelite ponovno aktiviranje naloga pritisnite na dugme ispod nakon cega ce Vam na email: <b>{props.match.params.id}</b> stici aktivacioni link.</div>

                    <div className="form-group">
                        <input className="btn btn-primary btn-block" type="submit" value="Send activation mail" />
                    </div>
                </form>
			</section>
		</React.Fragment>
	);
};

export default UserActivateRequestPage;
