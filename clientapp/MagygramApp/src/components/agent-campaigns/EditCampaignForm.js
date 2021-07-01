import React, { forwardRef, useContext, useEffect, useImperativeHandle, useState } from "react";
import { PostContext } from "../../contexts/PostContext";

const EditCampaignForm = forwardRef((props, ref) => {
	const { postState } = useContext(PostContext);

	const [minAge, setMinAge] = useState(postState.agentCampaignPostOptionModal.campaign.minAge);
	const [maxAge, setMaxAge] = useState(postState.agentCampaignPostOptionModal.campaign.maxAge);

	const [displayTime, setDisplayTime] = useState(postState.agentCampaignPostOptionModal.campaign.displayTime);
	const [checkedOnce, setCheckedOnce] = useState(postState.agentCampaignPostOptionModal.campaign.checkedOnce);
	const [gender, setGender] = useState(postState.agentCampaignPostOptionModal.campaign.gender);

	const [startDate, setStartDate] = useState(new Date(postState.agentCampaignPostOptionModal.campaign.startDate));
	const [endDate, setEndDate] = useState(new Date(postState.agentCampaignPostOptionModal.campaign.endDate));

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
			displayTime,
			checkedOnce,
			gender,
			startDate,
			endDate,
		},
	}));

	useEffect(() => {
		setMinAge(postState.agentCampaignPostOptionModal.campaign.minAge);
		setMaxAge(postState.agentCampaignPostOptionModal.campaign.maxAge);
		setDisplayTime(postState.agentCampaignPostOptionModal.campaign.displayTime);
		setCheckedOnce(postState.agentCampaignPostOptionModal.campaign.checkedOnce);
		setGender(postState.agentCampaignPostOptionModal.campaign.gender);
		setStartDate(new Date(postState.agentCampaignPostOptionModal.campaign.startDate));
		setEndDate(new Date(postState.agentCampaignPostOptionModal.campaign.endDate));
	}, [postState.agentCampaignPostOptionModal.campaign]);

	return (
		<div className="ml-3 mt-4">
			<div className="row">
				<div className="form-group">
					<h5 style={{ color: props.fontColor }}>Campaign information</h5>

					<label style={{ color: props.fontColor }}>Campaign frequency</label>

					<div className="form-check">
						<input className="form-check-input" type="radio" name="exampleRadios" id="exampleRadios1" value="once" checked={checkedOnce} onChange={() => setCheckedOnce(!checkedOnce)} />
						<label className="form-check-label" for="exampleRadios1" style={{ color: props.fontColor }}>
							One day campaign
						</label>
					</div>
					<div className="form-check">
						<input
							className="form-check-input"
							type="radio"
							name="exampleRadios"
							id="exampleRadios2"
							value="repeatedly"
							checked={!checkedOnce}
							onChange={() => setCheckedOnce(!checkedOnce)}
						/>
						<label className="form-check-label" for="exampleRadios2" style={{ color: props.fontColor }}>
							Multiple days campaign
						</label>
					</div>
				</div>
			</div>
			<div className="row" hidden={!checkedOnce}>
				<div className="form-group w-100">
					<label style={{ color: props.fontColor }}>Exposure time</label>
					<div className="form-row d-flex align-items-center">
						<div className="col">
							<input type="number" min="0" max="24" value={displayTime} className="form-control w-40" onChange={(e) => setDisplayTime(e.target.value)} />
						</div>
						<div className="col" style={{ color: props.fontColor }}>
							h
						</div>
					</div>
				</div>
			</div>
			<div className="row" hidden={checkedOnce}>
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

export default EditCampaignForm;
