import React from "react";

const Highlight = ({ highlight, openHighlightSlider }) => {
	const imgStyle = { transform: "scale(1.5)", width: "100%", position: "absolute", left: "0" };
	const visitedStoryClassName = "rounded-circle overflow-hidden d-flex justify-content-center align-items-center border border-danger story-profile-photo-visited";

	return (
		<React.Fragment>
			<li className="list-inline-item">
				<button onClick={() => openHighlightSlider(highlight.name)} className="btn p-0 m-0">
					<div className="d-flex flex-column align-items-center">
						<div className={visitedStoryClassName}>
							<img src={highlight.url !== "" ? highlight.url : "assets/img/star.jpg"} alt="" style={imgStyle} />
						</div>
						<small>{highlight.name}</small>
					</div>
				</button>
			</li>
		</React.Fragment>
	);
};

export default Highlight;
