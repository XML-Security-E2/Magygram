import React from "react";

const CollectionButton = ({ collectionName, media }) => {
	return (
		<React.Fragment>
			<div className="row w-100 " style={{ borderRadius: "5px" }}>
				<div className="col-6 p-1">
					<img src="/api/media/067d9199-5dd2-4285-a5a3-312b35340dcf.jpg" className="img-fluid rounded" />
				</div>
				<div className="col-6 p-1">
					<img src="/api/media/6afd9e6d-c427-476f-a12c-618486e13f05.jpg" className="img-fluid  rounded" />
				</div>
				<div className="col-6 p-1">
					<img src="/api/media/6afd9e6d-c427-476f-a12c-618486e13f05.jpg" className="img-fluid  rounded" />
				</div>
				<div className="col-6 p-1">
					<img src="/api/media/067d9199-5dd2-4285-a5a3-312b35340dcf.jpg" className="img-fluid   rounded" />
				</div>
			</div>
			<div className="row w-100 d-flex justify-content-center">Coll name</div>
		</React.Fragment>
	);
};

export default CollectionButton;
