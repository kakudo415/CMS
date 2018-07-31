"use strict";

// if (history.pushState && history.state !== undefined)

window.onload = () => {
	let contentHTML = document.getElementById("content");

	let Ajax = {
		XHR: new XMLHttpRequest(),
		Path: "",
		Time: Date.now(),
		Data: ""
	};

	const mouseOver = (ev) => {
		ajaxGET(ev.target.href);
	};

	const mouseDown = (ev) => {
		let count = 0;
		let timer = setInterval(() => {
			if (Ajax.Path == ev.target.href) {
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

	const ajaxGET = async (url) => {
		if (Ajax.Path != url && (Ajax.Time + 100) < Date.now()) {
			await Ajax.XHR.abort();
			Ajax.XHR.responseType = "document";
			Ajax.XHR.onload = (ev) => {
				Ajax.Path = url;
				Ajax.Time = Date.now();
				Ajax.Data = ev.target.responseXML;
			};
			Ajax.XHR.open("GET", url + "?e=e", true);
			Ajax.XHR.send();
		}
	};

	const movePage = () => {
		document.title = Ajax.Data.title;
		contentHTML.innerHTML = Ajax.Data.body.innerHTML;
		history.pushState(null, null, Ajax.Path);
	};

	for (let tag of document.getElementsByTagName("a")) {
		tag.onmouseover = mouseOver;
		tag.onmousedown = mouseDown;
		tag.onclick = () => { return false };
	}
};
