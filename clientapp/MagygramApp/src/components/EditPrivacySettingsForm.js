import { useContext, useEffect, useState } from "react";
import { UserContext } from "../contexts/UserContext";
import { userService } from "../services/UserService";
import EditPrivacySettingsSidebar from "./EditPrivacySettingsSidebar";

const EditPrivacySettingsForm = ({show}) => {
    const { userState, dispatch } = useContext(UserContext);

    const [isPrivate, setIsPrivate] = useState(userState.userProfile.user.privacySettings.isPrivate)
    const [receiveMessages, setReceiveMessages] = useState(userState.userProfile.user.privacySettings.receiveMessages)
    const [isTaggable, setIsTaggable] = useState(userState.userProfile.user.privacySettings.isTaggable)

    useEffect(() => {
        setIsPrivate(userState.userProfile.user.privacySettings.isPrivate);
        setReceiveMessages(userState.userProfile.user.privacySettings.receiveMessages);
        setIsTaggable(userState.userProfile.user.privacySettings.isTaggable)
	}, [userState.userProfile.user]);

    const handleSubmit = (e) => {
		e.preventDefault();

		let privacySettingsReq = {
            isPrivate,
            receiveMessages,
            isTaggable
		};

		userService.editUserPrivacySettings(userState.userProfile.showedUserId, privacySettingsReq, dispatch);
	};

    return (
        <form hidden={!show} method="post" onSubmit={handleSubmit}>
        <div>
            <br />
            <div className="form-group row">
                <label className="col-sm-8 col-form-label">Profile privacy</label>
                <div className="col-sm-4">
                    <label>
                        <input type="checkbox" className="mr-1" checked={isPrivate} onChange={() => setIsPrivate(!isPrivate)} />
                    </label>
                </div>
            </div>
            <div className="form-group row d-flex align-items-center">
                <label className="col-sm-8 col-form-label">Receive messages from users you are not following</label>
                <div className="col-sm-4">
                    <label>
                        <input type="checkbox" className="mr-1" checked={receiveMessages} onChange={() => setReceiveMessages(!receiveMessages)} />
                    </label>
                </div>
            </div>
            <div className="form-group row d-flex align-items-center">
                <label className="col-sm-8 col-form-label">Can other users tag you in posts, stories or messages</label>
                <div className="col-sm-4">
                    <label>
                        <input type="checkbox" className="mr-1" checked={isTaggable} onChange={() => setIsTaggable(!isTaggable)} />
                    </label>
                </div>
            </div>
            <div className="form-group">
                <button className="btn btn-primary float-right  mb-2" type="submit">
                    Save
                </button>
            </div>
        </div>
    </form>
    );
};

export default EditPrivacySettingsForm;