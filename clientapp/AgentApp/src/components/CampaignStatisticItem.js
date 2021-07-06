import React from "react";
import { Link } from "react-router-dom";

const CampaignStatisticItem = ({ statistics }) => {
	return (
		<React.Fragment>
			{statistics.campaignType === "POST" &&
				(statistics.campaignStatus === "REGULAR" ? (
					<div className="col-12">
						<hr />

						<div className="row">
							<div className="col-3">
								{statistics.media.MediaType === "IMAGE" ? (
									<img src={statistics.media.Url} className="img-fluid rounded-lg w-100" alt="" />
								) : (
									<video className="img-fluid box-coll rounded-lg w-100" style={{ objectFit: "cover" }}>
										<source src={statistics.media.Url} type="video/mp4" />
									</video>
								)}
							</div>
							{statistics.frequency === "REPEATEDLY" ? (
								<div className="col-4">
									<div className="row">
										<h5>Campaign info</h5>
									</div>
									<div className="row">
										<b className="mr-1">Campaign type: </b> {statistics.campaignType}
									</div>
									<div className="row d-flex align-items-center">
										<b>Website: </b>
										<button type="button" className="btn btn-link border-0" onClick={() => window.open("https://" + statistics.website, "_blank")}>
											{statistics.website}
										</button>
									</div>

									<div className="row">
										<b className="mr-1">Date from: </b>
										{new Date(statistics.dateFrom).toLocaleDateString("en-US", {
											day: "2-digit",
											month: "2-digit",
											year: "numeric",
										})}
									</div>
									<div className="row">
										<b className="mr-1">Date to: </b>
										{new Date(statistics.dateTo).toLocaleDateString("en-US", {
											day: "2-digit",
											month: "2-digit",
											year: "numeric",
										})}
									</div>
									<div className="row">
										<b className="mr-1">Minimum times to display: </b> {statistics.minDisplaysForRepeatedly}
									</div>
								</div>
							) : (
								<div className="col-4">
									<div className="row">
										<h5>Campaign info</h5>
									</div>

									<div className="row">
										<b className="mr-1">Campaign type: </b> {statistics.campaignType}
									</div>
									<div className="row d-flex align-items-center">
										<b>Website: </b>
										<button type="button" className="btn btn-link border-0" onClick={() => window.open("https://" + statistics.website, "_blank")}>
											{statistics.website}
										</button>
									</div>

									<div className="row">
										<b className="mr-1">Exposure date: </b>
										{new Date(statistics.exposeOnceDate).toLocaleDateString("en-US", {
											day: "2-digit",
											month: "2-digit",
											year: "numeric",
										})}
									</div>
									<div className="row">
										<b className="mr-1">Exposure time: </b>
										{statistics.displayTime} h
									</div>
								</div>
							)}
							<div className="col-3">
								<div className="row">
									<h5>Campaign reach</h5>
								</div>
								<div className="row">
									<b className="mr-1">Website clicks: </b> {statistics.websiteClicks}
								</div>
								<div className="row">
									<b className="mr-1">Daily average views: </b> {statistics.dailyAverage}
								</div>
								<div className="row">
									<b className="mr-1">Campaign views: </b> {statistics.userViews}
								</div>
								<div className="row">
									<b className="mr-1">Likes: </b> {statistics.Likes}
								</div>
								<div className="row">
									<b className="mr-1">Dislikes: </b> {statistics.Dislikes}
								</div>
								<div className="row">
									<b className="mr-1">Comments: </b> {statistics.Comments}
								</div>
							</div>
							<div className="col-2">
								<div className="row">
									<h5>Target group</h5>
								</div>
								<div className="row">
									<b className="mr-1">Min age: </b> {statistics.targetGroup.minAge}
								</div>
								<div className="row">
									<b className="mr-1">Max age: </b> {statistics.targetGroup.maxAge}
								</div>
								<div className="row">
									<b className="mr-1">Gender: </b> {statistics.targetGroup.gender}
								</div>
							</div>
						</div>
					</div>
				) : (
					<div className="col-12">
						<hr />

						<div className="row">
							<div className="col-3">
								{statistics.media.MediaType === "IMAGE" ? (
									<img src={statistics.media.Url} className="img-fluid rounded-lg w-100" alt="" />
								) : (
									<video className="img-fluid box-coll rounded-lg w-100" style={{ objectFit: "cover" }}>
										<source src={statistics.media.Url} type="video/mp4" />
									</video>
								)}
							</div>
							<div className="col-4">
								<div className="row">
									<h5>Campaign info</h5>
								</div>
								<div className="row d-flex align-items-center">
									<b>Website: </b>
									<button type="button" className="btn btn-link border-0" onClick={() => window.open("https://" + statistics.website, "_blank")}>
										{statistics.website}
									</button>
								</div>

								<div className="row">
									<Link className="font-weight-bold btn btn-link" style={{ cursor: "pointer" }} to={"/profile?userId=" + statistics.influencerId}>
										@{statistics.influencerUsername}
									</Link>
								</div>
							</div>

							<div className="col-3">
								<div className="row">
									<h5>Campaign reach</h5>
								</div>
								<div className="row">
									<b className="mr-1">Website clicks: </b> {statistics.websiteClicks}
								</div>
								<div className="row">
									<b className="mr-1">Daily average views: </b> {statistics.dailyAverage}
								</div>
								<div className="row">
									<b className="mr-1">Campaign views: </b> {statistics.userViews}
								</div>
								<div className="row">
									<b className="mr-1">Likes: </b> {statistics.Likes}
								</div>
								<div className="row">
									<b className="mr-1">Dislikes: </b> {statistics.Dislikes}
								</div>
								<div className="row">
									<b className="mr-1">Comments: </b> {statistics.Comments}
								</div>
							</div>
							<div className="col-2">
								<div className="row">
									<h5>Target group</h5>
								</div>
								<div className="row">
									<b className="mr-1">Min age: </b> {statistics.targetGroup.minAge}
								</div>
								<div className="row">
									<b className="mr-1">Max age: </b> {statistics.targetGroup.maxAge}
								</div>
								<div className="row">
									<b className="mr-1">Gender: </b> {statistics.targetGroup.gender}
								</div>
							</div>
						</div>
					</div>
				))}

			{statistics.campaignType === "STORY" &&
				(statistics.campaignStatus === "REGULAR" ? (
					<div className="col-12">
						<hr />

						<div className="row">
							<div className="col-3">
								{statistics.media.MediaType === "IMAGE" ? (
									<img src={statistics.media.Url} className="img-fluid rounded-lg w-100" alt="" />
								) : (
									<video className="img-fluid box-coll rounded-lg w-100" style={{ objectFit: "cover" }}>
										<source src={statistics.media.Url} type="video/mp4" />
									</video>
								)}
							</div>
							{statistics.frequency === "REPEATEDLY" ? (
								<div className="col-4">
									<div className="row">
										<h5>Campaign info</h5>
									</div>
									<div className="row">
										<b className="mr-1">Campaign type: </b> {statistics.campaignType}
									</div>
									<div className="row d-flex align-items-center">
										<b>Website: </b>
										<button type="button" className="btn btn-link border-0" onClick={() => window.open("https://" + statistics.website, "_blank")}>
											{statistics.website}
										</button>
									</div>

									<div className="row">
										<b className="mr-1">Date from: </b>
										{new Date(statistics.dateFrom).toLocaleDateString("en-US", {
											day: "2-digit",
											month: "2-digit",
											year: "numeric",
										})}
									</div>
									<div className="row">
										<b className="mr-1">Date to: </b>
										{new Date(statistics.dateTo).toLocaleDateString("en-US", {
											day: "2-digit",
											month: "2-digit",
											year: "numeric",
										})}
									</div>
									<div className="row">
										<b className="mr-1">Minimum times to display: </b> {statistics.minDisplaysForRepeatedly}
									</div>
								</div>
							) : (
								<div className="col-4">
									<div className="row">
										<h5>Campaign info</h5>
									</div>

									<div className="row">
										<b className="mr-1">Campaign type: </b> {statistics.campaignType}
									</div>
									<div className="row d-flex align-items-center">
										<b>Website: </b>
										<button type="button" className="btn btn-link border-0" onClick={() => window.open("https://" + statistics.website, "_blank")}>
											{statistics.website}
										</button>
									</div>

									<div className="row">
										<b className="mr-1">Exposure date: </b>
										{new Date(statistics.exposeOnceDate).toLocaleDateString("en-US", {
											day: "2-digit",
											month: "2-digit",
											year: "numeric",
										})}
									</div>
									<div className="row">
										<b className="mr-1">Exposure time: </b>
										{statistics.displayTime} h
									</div>
								</div>
							)}
							<div className="col-3">
								<div className="row">
									<h5>Campaign reach</h5>
								</div>
								<div className="row">
									<b className="mr-1">Website clicks: </b> {statistics.websiteClicks}
								</div>
								<div className="row">
									<b className="mr-1">Daily average views: </b> {statistics.dailyAverage}
								</div>
								<div className="row">
									<b className="mr-1">Campaign views: </b> {statistics.userViews}
								</div>
								<div className="row">
									<b className="mr-1">Story views: </b> {statistics.StoryViews}
								</div>
							</div>
							<div className="col-2">
								<div className="row">
									<h5>Target group</h5>
								</div>
								<div className="row">
									<b className="mr-1">Min age: </b> {statistics.targetGroup.minAge}
								</div>
								<div className="row">
									<b className="mr-1">Max age: </b> {statistics.targetGroup.maxAge}
								</div>
								<div className="row">
									<b className="mr-1">Gender: </b> {statistics.targetGroup.gender}
								</div>
							</div>
						</div>
					</div>
				) : (
					<div className="col-12">
						<hr />

						<div className="row">
							<div className="col-3">
								{statistics.media.MediaType === "IMAGE" ? (
									<img src={statistics.media.Url} className="img-fluid rounded-lg w-100" alt="" />
								) : (
									<video className="img-fluid box-coll rounded-lg w-100" style={{ objectFit: "cover" }}>
										<source src={statistics.media.Url} type="video/mp4" />
									</video>
								)}
							</div>
							<div className="col-4">
								<div className="row">
									<h5>Campaign info</h5>
								</div>
								<div className="row d-flex align-items-center">
									<b>Website: </b>
									<button type="button" className="btn btn-link border-0" onClick={() => window.open("https://" + statistics.website, "_blank")}>
										{statistics.website}
									</button>
								</div>

								<div className="row">
									<Link className="font-weight-bold btn btn-link" style={{ cursor: "pointer" }} to={"/profile?userId=" + statistics.influencerId}>
										@{statistics.influencerUsername}
									</Link>
								</div>
							</div>

							<div className="col-3">
								<div className="row">
									<h5>Campaign reach</h5>
								</div>
								<div className="row">
									<b className="mr-1">Website clicks: </b> {statistics.websiteClicks}
								</div>
								<div className="row">
									<b className="mr-1">Daily average views: </b> {statistics.dailyAverage}
								</div>
								<div className="row">
									<b className="mr-1">Campaign views: </b> {statistics.userViews}
								</div>
								<div className="row">
									<b className="mr-1">Story views: </b> {statistics.StoryViews}
								</div>
							</div>
							<div className="col-2">
								<div className="row">
									<h5>Target group</h5>
								</div>
								<div className="row">
									<b className="mr-1">Min age: </b> {statistics.targetGroup.minAge}
								</div>
								<div className="row">
									<b className="mr-1">Max age: </b> {statistics.targetGroup.maxAge}
								</div>
								<div className="row">
									<b className="mr-1">Gender: </b> {statistics.targetGroup.gender}
								</div>
							</div>
						</div>
					</div>
				))}
		</React.Fragment>
	);
};

export default CampaignStatisticItem;
