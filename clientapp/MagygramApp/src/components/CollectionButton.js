import React from "react";

const CollectionButton = ({ collectionName, media }) => {
	return (
		<React.Fragment>
			<div className="row w-100 " style={{ borderRadius: "5px", backgroundColor: "lightgray", cursor: "pointer" }}>
				<div className="col-6  box-coll  p-1">
					{media.length > 0 ? <img src={media[0].media.url} className="img-fluid rounded" alt="col" /> : <img src="" alt="" className="img-fluid rounded" />}
				</div>
				<div className="col-6 box-coll  p-1">
					{media.length > 1 ? <img src={media[1].media.url} className="img-fluid rounded" alt="col" /> : <img src="" alt="" className="img-fluid rounded" />}
				</div>
				<div className="col-6 box-coll  p-1">
					{media.length > 2 ? <img src={media[2].media.url} className="img-fluid rounded" alt="col" /> : <img src="" alt="" className="img-fluid rounded" />}
				</div>
				<div className="col-6 box-coll  p-1">
					{media.length > 3 ? <img src={media[3].media.url} className="img-fluid rounded" alt="col" /> : <img src="" alt="" className="img-fluid rounded" />}
				</div>
			</div>
			<div className="row w-100 d-flex justify-content-center">{collectionName}</div>
		</React.Fragment>
	);
};

export default CollectionButton;
