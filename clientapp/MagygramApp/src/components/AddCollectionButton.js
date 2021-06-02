import React, { useState } from "react";

const AddCollectionButton = ({ createCollection }) => {
	const [showInputField, setShowInputField] = useState(false);
	const [collName, setCollName] = useState("");

	const handleCollectionCreateRequest = () => {
		createCollection(collName);
		setShowInputField(false);
		setCollName("");
	};

	return (
		<React.Fragment>
			<div className="row w-100 " style={{ borderRadius: "5px", backgroundColor: "lightgray", cursor: "pointer" }}>
				<div className="col-12  box-coll  p-1">
					{showInputField ? (
						<button disabled={collName.length === 0} type="button" className="btn btn-outline-secondary rounded-lg btn-icon w-100 h-100 border-0" onClick={handleCollectionCreateRequest}>
							<i className="mdi mdi-plus w-50 h-50"></i>
							Add
						</button>
					) : (
						<button type="button" className="btn btn-outline-secondary rounded-lg btn-icon w-100 h-100 border-0" onClick={() => setShowInputField(true)}>
							<i className="mdi mdi-plus w-50 h-50"></i>
						</button>
					)}
				</div>
			</div>
			<div className="row w-100 d-flex justify-content-center">
				<input hidden={!showInputField} type="text" placeholder="Collection" className="form-control" onChange={(e) => setCollName(e.target.value)} />
			</div>
		</React.Fragment>
	);
};

export default AddCollectionButton;
