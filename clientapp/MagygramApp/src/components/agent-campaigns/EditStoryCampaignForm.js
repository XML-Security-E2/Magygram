import React, { forwardRef, useContext, useEffect, useImperativeHandle, useState } from "react";
import { StoryContext } from "../../contexts/StoryContext";

const EditStoryCampaignForm = forwardRef((props, ref) => {
	const { storyState } = useContext(StoryContext);

	const [minAge, setMinAge] = useState(storyState.agentCampaignStoryOptionModal.campaign.minAge);
	const [maxAge, setMaxAge] = useState(storyState.agentCampaignStoryOptionModal.campaign.maxAge);
	const [minDisplaysForRepeatedly, setMinDisplaysForRepeatedly] = useState(storyState.agentCampaignStoryOptionModal.campaign.minDisplays);

	const [gender, setGender] = useState(storyState.agentCampaignStoryOptionModal.campaign.gender);

	const [startDate, setStartDate] = useState(new Date(storyState.agentCampaignStoryOptionModal.campaign.startDate));
	const [endDate, setEndDate] = useState(new Date(storyState.agentCampaignStoryOptionModal.campaign.endDate));

	const handleStartDateChange = (e) => {
		let date = new Date(e.target.value);
		setStartDate(date);

		if (date >= endDate) {
			let newDate = new Date(e.target.value);
			newDate.setDate(newDate.getDate() + 1);
			setEndDate(newDate);
		}
	};

	useImperativeHandle(ref, () => ({
		campaignState: {
			minAge,
			maxAge,
			minDisplaysForRepeatedly,
			gender,
			startDate,
			endDate,
		},
	}));

	useEffect(() => {
		setMinAge(storyState.agentCampaignStoryOptionModal.campaign.minAge);
		setMaxAge(storyState.agentCampaignStoryOptionModal.campaign.maxAge);
		setMinDisplaysForRepeatedly(storyState.agentCampaignStoryOptionModal.campaign.minDisplays);
		setGender(storyState.agentCampaignStoryOptionModal.campaign.gender);
		setStartDate(new Date(storyState.agentCampaignStoryOptionModal.campaign.startDate));
		setEndDate(new Date(storyState.agentCampaignStoryOptionModal.campaign.endDate));
	}, [storyState.agentCampaignStoryOptionModal.campaign]);

	return (
		<div className="ml-3 mt-4">
			<div className="row">
				<div className="form-group">
					<label>Exposure dates</label>

					<div className="form-row">
						<div className="col" style={{ color: props.fontColor }}>
							Start date
							<input
								type="date"
								className="form-control"
								value={startDate.toISOString().split("T")[0]}
								min={new Date().toISOString().split("T")[0]}
								onChange={(e) => handleStartDateChange(e)}
							/>
						</div>
						<div className="col" style={{ color: props.fontColor }}>
							End date
							<input
								type="date"
								className="form-control"
								value={endDate.toISOString().split("T")[0]}
								min={new Date(startDate.getTime() + 24 * 60 * 60 * 1000).toISOString().split("T")[0]}
								onChange={(e) => setEndDate(new Date(e.target.value))}
							/>
						</div>
					</div>
				</div>
				<div className="form-group">
					<div className="form-row">
						<div className="col" style={{ color: props.fontColor }}>
							Times to expose per day
							<input type="number" min="1" value={minDisplaysForRepeatedly} className="form-control" onChange={(e) => setMinDisplaysForRepeatedly(e.target.value)} />
						</div>
					</div>
				</div>
			</div>

			<div className="row">
				<div className="form-group">
					<h5 style={{ color: props.fontColor }}>Target group</h5>

					<label style={{ color: props.fontColor }}>Target gender</label>

					<div className="form-group">
						<div className="form-check form-check-inline">
							<input className="form-check-input" type="radio" name="exampleRadios1" id="exampleRadios6" value="ANY" checked={gender === "ANY"} onChange={() => setGender("ANY")} />
							<label className="form-check-label" for="exampleRadios6" style={{ color: props.fontColor }}>
								Any
							</label>
						</div>
						<div className="form-check form-check-inline ml-2">
							<input className="form-check-input" type="radio" name="exampleRadios1" id="exampleRadios5" value="MALE" checked={gender === "MALE"} onChange={() => setGender("MALE")} />
							<label className="form-check-label" for="exampleRadios5" style={{ color: props.fontColor }}>
								Male
							</label>
						</div>
						<div className="form-check form-check-inline ml-2">
							<input
								className="form-check-input"
								type="radio"
								name="exampleRadios1"
								id="exampleRadios4"
								value="FEMALE"
								checked={gender === "FEMALE"}
								onChange={() => setGender("FEMALE")}
							/>
							<label className="form-check-label" for="exampleRadios4" style={{ color: props.fontColor }}>
								Female
							</label>
						</div>
					</div>
				</div>
			</div>

			<div className="row">
				<div className="form-group">
					<label style={{ color: props.fontColor }}>Target age</label>

					<div className="form-row">
						<div className="col" style={{ color: props.fontColor }}>
							Min age
							<input type="number" min="16" max="110" value={minAge} className="form-control" onChange={(e) => setMinAge(e.target.value)} />
						</div>
						<div className="col" style={{ color: props.fontColor }}>
							Max age
							<input type="number" min="16" max="110" value={maxAge} className="form-control" onChange={(e) => setMaxAge(e.target.value)} />
						</div>
					</div>
				</div>
			</div>
		</div>
	);
});

export default EditStoryCampaignForm;
