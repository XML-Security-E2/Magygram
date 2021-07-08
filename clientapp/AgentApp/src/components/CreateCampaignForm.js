import React, { forwardRef, useImperativeHandle, useState } from "react";

const CreateCampaignForm = forwardRef((props, ref) => {
	const [minAge, setMinAge] = useState(16);
	const [maxAge, setMaxAge] = useState(23);
	const [minDisplaysForRepeatedly, setMinDisplaysForRepeatedly] = useState(1);

	const [displayTime, setDisplayTime] = useState(12);
	const [checkedOnce, setCheckedOnce] = useState(true);
	const [checkedPost, setCheckedPost] = useState(true);
	const [gender, setGender] = useState("ANY");

	const [startDate, setStartDate] = useState(new Date());
	const [exposeOnceDate, setExposeOnceDate] = useState(new Date());
	const [endDate, setEndDate] = useState(new Date(new Date().getTime() + 24 * 60 * 60 * 1000));

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
			checkedPost,
			gender,
			startDate,
			endDate,
			exposeOnceDate,
			minDisplaysForRepeatedly,
		},
	}));

	return (
		<div className="mt-4">
			<div className="row">
				<div className="form-group">
					<h5>Campaign type</h5>

					<div className="form-check">
						<input className="form-check-input" type="radio" name="exampleRadios11" id="exampleRadios11" value="post" checked={checkedPost} onChange={() => setCheckedPost(!checkedPost)} />
						<label className="form-check-label" for="exampleRadios11">
							Post
						</label>
					</div>
					<div className="form-check">
						<input
							className="form-check-input"
							type="radio"
							name="exampleRadios21"
							id="exampleRadios21"
							value="story"
							checked={!checkedPost}
							onChange={() => setCheckedPost(!checkedPost)}
						/>
						<label className="form-check-label" for="exampleRadios21">
							Story
						</label>
					</div>
				</div>
			</div>
			<div className="row">
				<div className="form-group">
					<h5>Campaign information</h5>

					<label>Campaign frequency</label>

					<div className="form-check">
						<input className="form-check-input" type="radio" name="exampleRadios" id="exampleRadios1" value="once" checked={checkedOnce} onChange={() => setCheckedOnce(!checkedOnce)} />
						<label className="form-check-label" for="exampleRadios1">
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
						<label className="form-check-label" for="exampleRadios2">
							Multiple days campaign
						</label>
					</div>
				</div>
			</div>
			<div className="row" hidden={!checkedOnce}>
				<div className="form-group w-100">
					<label>Exposure time</label>
					<div className="form-row d-flex align-items-center">
						<div className="col">
							<input type="number" min="0" max="24" value={displayTime} className="form-control w-40" onChange={(e) => setDisplayTime(e.target.value)} />
						</div>
						<div className="col">h</div>
					</div>
				</div>
			</div>
			<div className="row" hidden={checkedOnce}>
				<div className="form-group">
					<label>Exposure dates</label>

					<div className="form-row">
						<div className="col">
							Start date
							<input
								type="date"
								className="form-control"
								value={startDate.toISOString().split("T")[0]}
								min={new Date().toISOString().split("T")[0]}
								onChange={(e) => handleStartDateChange(e)}
							/>
						</div>
						<div className="col">
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
					<div className="form-group mt-2">
						<div className="form-row">
							<div className="col">
								Times to expose per day
								<input type="number" min="1" value={minDisplaysForRepeatedly} className="form-control" onChange={(e) => setMinDisplaysForRepeatedly(e.target.value)} />
							</div>
						</div>
					</div>
				</div>
			</div>
			<div className="row" hidden={!checkedOnce}>
				<div className="form-group">
					<label>Exposure date</label>
					<div className="form-row">
						<div className="col">
							<input
								type="date"
								className="form-control"
								value={exposeOnceDate.toISOString().split("T")[0]}
								min={new Date().toISOString().split("T")[0]}
								onChange={(e) => setExposeOnceDate(new Date(e.target.value))}
							/>
						</div>
					</div>
				</div>
			</div>

			<div className="row">
				<div className="form-group">
					<h5>Target group</h5>

					<label>Target gender</label>

					<div className="form-group">
						<div className="form-check form-check-inline">
							<input className="form-check-input" type="radio" name="exampleRadios1" id="exampleRadios6" value="ANY" checked={gender === "ANY"} onChange={() => setGender("ANY")} />
							<label className="form-check-label" for="exampleRadios6">
								Any
							</label>
						</div>
						<div className="form-check form-check-inline ml-2">
							<input className="form-check-input" type="radio" name="exampleRadios1" id="exampleRadios5" value="MALE" checked={gender === "MALE"} onChange={() => setGender("MALE")} />
							<label className="form-check-label" for="exampleRadios5">
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
							<label className="form-check-label" for="exampleRadios4">
								Female
							</label>
						</div>
					</div>
				</div>
			</div>

			<div className="row">
				<div className="form-group">
					<label>Target age</label>

					<div className="form-row">
						<div className="col">
							Min age
							<input type="number" min="16" max="110" value={minAge} className="form-control" onChange={(e) => setMinAge(e.target.value)} />
						</div>
						<div className="col">
							Max age
							<input type="number" min="16" max="110" value={maxAge} className="form-control" onChange={(e) => setMaxAge(e.target.value)} />
						</div>
					</div>
				</div>
			</div>
			<button type="button" className="btn btn-primary" onClick={props.handleSubmit}>
				Create
			</button>
		</div>
	);
});

export default CreateCampaignForm;
