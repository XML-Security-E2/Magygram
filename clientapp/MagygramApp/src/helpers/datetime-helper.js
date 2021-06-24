export function getDateTime(timestamp) {
	let date = new Date(timestamp);
	let str = "";

	if (Math.abs(new Date() - date) > 1000 * 60 * 60 * 24) {
		let day = date.getDay();
		let month = date.getUTCMonth();

		if (day.toString().length > 1) {
			str += day.toString() + "/";
		} else {
			str += "0" + day.toString() + "/";
		}

		if (month.toString().length > 1) {
			str += month.toString();
		} else {
			str += "0" + month.toString();
		}
	} else {
		let hour = date.getHours();
		let minute = date.getMinutes();

		if (hour.toString().length > 1) {
			str += hour.toString() + ":";
		} else {
			str += "0" + hour.toString() + ":";
		}

		if (minute.toString().length > 1) {
			str += minute.toString();
		} else {
			str += "0" + minute.toString();
		}
	}
	return str;
}
