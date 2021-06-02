import React from "react";

const CollectionButton = ({ collectionName, media }) => {
	return (
		<React.Fragment>
			<div className="row w-100 " style={{ borderRadius: "5px", backgroundColor: "lightgray", cursor: "pointer" }}>
				<div className="col-6  box-coll  p-1">
					{media.length > 0 ? (
						media[0].media.mediaType === "IMAGE" ? (
							<img src={media[0].media.url} className="img-fluid rounded" alt="col" />
						) : (
							<video className="img-fluid rounded-lg w-100">
								<source src={media[0].media.url} type="video/mp4" />
							</video>
						)
					) : (
						<img src="" alt="" className="img-fluid rounded" />
					)}
				</div>
				<div className="col-6 box-coll  p-1">
					{media.length > 1 ? (
						media[1].media.mediaType === "IMAGE" ? (
							<img src={media[1].media.url} className="img-fluid rounded" alt="col" />
						) : (
							<video className="img-fluid rounded-lg w-100">
								<source src={media[1].media.url} type="video/mp4" />
							</video>
						)
					) : (
						<img src="" alt="" className="img-fluid rounded w-100" />
					)}
				</div>
				<div className="col-6 box-coll  p-1">
					{media.length > 2 ? (
						media[2].media.mediaType === "IMAGE" ? (
							<img src={media[2].media.url} className="img-fluid rounded" alt="col" />
						) : (
							<video className="img-fluid rounded-lg w-100">
								<source src={media[2].media.url} type="video/mp4" />
							</video>
						)
					) : (
						<img src="" alt="" className="img-fluid rounded" />
					)}
				</div>
				<div className="col-6 box-coll  p-1">
					{media.length > 3 ? (
						media[3].media.mediaType === "IMAGE" ? (
							<img src={media[3].media.url} className="img-fluid rounded" alt="col" />
						) : (
							<video className="img-fluid rounded-lg w-100">
								<source src={media[3].media.url} type="video/mp4" />
							</video>
						)
					) : (
						<img src="" alt="" className="img-fluid rounded" />
					)}
				</div>
			</div>
			<div className="row w-100 d-flex justify-content-center">{collectionName}</div>
		</React.Fragment>
	);
};

export default CollectionButton;
