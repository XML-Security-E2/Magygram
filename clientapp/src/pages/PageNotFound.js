import React from "react";

const PageNotFound = () => {
	return (
		<React.Fragment>
			<h1>Not found</h1>
			<div className="illustration">
				<i className="icon ion-ios-navigate"></i>
			</div>
			<div className="text-center mt-5" style={{ fontSize: "6em", color: "#1977cc" }}>
				<b>404</b>
			</div>
			<div className="text-center mt-5" style={{ fontSize: "3em" }}>
				Not Found
			</div>
		</React.Fragment>
	);
};

export default PageNotFound;
