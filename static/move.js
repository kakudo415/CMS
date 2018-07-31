"use strict";

// if (history.pushState && history.state !== undefined)

window.onload = () => {
	let Ajax = {
		XHR: new XMLHttpRequest(),
		Path: "",
		Time: Date.now(),
		Data: ""
	};

	const mouseOver = async (ev) => {
		ajaxGET(ev.target.href);
	};

	const mouseDown = async (ev) => {
		let count = 0;
		let timer = setInterval(() => {
			if (Ajax.Data.length > 0) {
				movePage();
				clearTimeout(timer);
			}
			count++;
			if (count > 300) {
				clearTimeout(timer);
				location.href = Ajax.Path;
			}
		}, 10);
	};

	const ajaxGET = (url) => {
		if (Ajax.Path != url && (Ajax.Time + 100) < Date.now()) {
			Ajax.XHR.abort();
			Ajax.XHR.open("GET", url + "?e=e", true);
			Ajax.XHR.onloadstart = (ev) => {
				Ajax.Data = "";
			};
			Ajax.XHR.onload = (ev) => {
				Ajax.Path = url;
				Ajax.Time = Date.now();
				Ajax.Data = ev.target.responseText;
			};
			Ajax.XHR.send();
		}
	};

	const movePage = () => {
		history.pushState(null, null, Ajax.Path);
	};

	for (let tag of document.getElementsByTagName("a")) {
		tag.onmouseover = mouseOver;
		tag.onmousedown = mouseDown;
		tag.onclick = () => { return false };
	}
};
