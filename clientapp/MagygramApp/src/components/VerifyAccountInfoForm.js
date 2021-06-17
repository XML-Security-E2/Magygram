import React, {useState, useContext} from "react";
import { profileSettingsConstants } from "../constants/ProfileSettingsConstants";
import { ProfileSettingsContext } from "../contexts/ProfileSettingsContext";
import { requestsService }  from "../services/RequestsService"
import FailureAlert from "./FailureAlert";

const VerifyAccontInfoForm = ({show}) => {
    const { profileSettingsState, profileSettingsDispatch } = useContext(ProfileSettingsContext);

    const [category, setCategory] = useState("");
	const [name, setName] = useState("");
	const [surname, setSurname] = useState("");
	const [showForm, setShowForm] = useState(false);
	const [image, setImage] = useState("");
    const [imageName, setImageName] = useState("")
	const imgRef = React.createRef();

    const handleCheckBoxChange = () =>{
        setShowForm(!showForm)
    }

    const onImageChange = (e) =>{ 
        setImage(e.target.files[0]);
        setImageName(e.target.files[0].name)
        if(e.target.files[0].name.length>25){
            setImageName(e.target.files[0].name.substr(0,25)+"...")
        }
    }

    const handleSubmit = (e) =>{
        e.preventDefault();

        let formData = new FormData();
	    formData.append(`images[0]`, image);
	    formData.append(`name`, name);
	    formData.append(`surname`, surname);
	    formData.append(`category`, category);

        if(image!=="" && image !=='Select image'){
            requestsService.createVerificationRequest(formData,profileSettingsDispatch)
        }else{
            setImageName('Select image')
        }


    }

	return (
		<React.Fragment>
            {!profileSettingsState.isUserVerified? <div>
                <div hidden={profileSettingsState.sendedVerifyRequest} className="row">
                    <div className="col-12">
                        <FailureAlert
                            hidden={!profileSettingsState.sendRequest.showError}
                            header="Error"
                            message={profileSettingsState.sendRequest.errorMessage}
                            handleCloseAlert={() => profileSettingsDispatch({ type: profileSettingsConstants.CREATE_VERIFICATION_REQUEST_REQUEST })}
                        />
                    </div>
                </div>
                <form hidden={!show} method="post" onSubmit={handleSubmit}>
                    <div hidden={profileSettingsState.sendedVerifyRequest}>
                        <br />

                        <div className="form-group row ml-4 mt-3">
                            <label>
                                <input
                                    name="verifyAccount"
                                    type="checkbox"
                                    className="mr-1"
                                    checked={showForm}
                                    onChange={handleCheckBoxChange} 
                                />
                                Verify account
                            </label>
                        </div>
                        <div hidden={!showForm}>
                            <div className="form-group row">
                            <input type="file" ref={imgRef} style={{ display: "none" }} name="image" accept="image/png, image/jpeg" onChange={onImageChange} />

                                <label for="name" className="col-sm-3 col-form-label">
                                    <b>Name</b>
                                </label>
                                <div class="col-sm-9">
                                    <input required className="form-control" type="text" id="name" name="name" placeholder="Name" value={name} onChange={(e) => setName(e.target.value)} />
                                </div>
                            </div>
                            <div className="form-group row">
                                <label for="surname" className="col-sm-3 col-form-label">
                                    <b>Surname</b>
                                </label>
                                <div class="col-sm-9">
                                    <input required className="form-control" type="text" id="surname" name="surname" placeholder="Surname" value={surname} onChange={(e) => setSurname(e.target.value)} />
                                </div>
                            </div>
                            <div className="form-group row">
                                <label for="email" className="col-sm-3 col-form-label">
                                    <b>Category</b>
                                </label>
                                <div className="col-sm-9">
                                    <select required id="gender" className="form-control" value={category} onChange={(e) => setCategory(e.target.value)}>
                                        <option value="" disabled>
                                            Select category
                                        </option>
                                        <option value="INFLUENCER"> Influencer</option>
                                        <option value="SPORTS"> Sports</option>
                                        <option value="NEWS/MEDIA"> News/Media</option>
                                        <option value="BUSINESS"> Business</option>
                                        <option value="BRAND"> Brand</option>
                                        <option value="ORGANIZATION"> Organization</option>
                                        <option value="MUSIC"> Music</option>
                                        <option value="ACTOR"> Actor</option>
                                    </select>
                                </div>
                            </div>
                            <div className="form-group row">
                                <label for="surname" className="col-sm-3 col-form-label">
                                    <b>Document</b>
                                </label>
                                <div className="col-sm-9 r">
                                    <label className="mt-2"  style={{width:"150px"}}>{imageName}</label>
                                    <button style={{width: "170px"}} type="button" className="btn form-control float-right" onClick={() => imgRef.current.click()}>
                                        <b className="text-primary">Change document</b>
                                    </button>                            
                                </div>
                            </div>
                            <div className="form-group">
                                <button className="btn btn-primary float-right  mb-2 mt-2" type="submit">
                                    Send request
                                </button>
                            </div>
                        </div>
                    </div>
                    <div className="col-12 mt-5 d-flex justify-content-center text-secondary" >
                        <h3 hidden={!profileSettingsState.sendedVerifyRequest}>The request has been sent</h3>
                    </div>
                </form>
                </div>:
                <div>
                    <div className="col-12 mt-5 d-flex justify-content-center text-secondary" >
                        <h3>Your account is verified</h3>
                    </div>
                </div>}
            
		</React.Fragment>
	);
};

export default VerifyAccontInfoForm;
