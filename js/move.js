"use strict";

if (history.pushState && history.state !== undefined) {
	let contentHTML = document.getElementById("content");
	let Ajax = {
		XHR: new XMLHttpRequest(),
		Path: "",
		Time: Date.now(),
		Data: null
	};

	const isLocal = (uri) => {
		return uri.startsWith("http://" + document.domain) || uri.startsWith("https://" + document.domain) || uri.startsWith("//" + document.domain) || uri.startsWith("/");
	};

	const preGET = async (uri) => {
		if (Ajax.Path !== uri && (Ajax.Time + 100) < Date.now()) {
			await Ajax.XHR.abort();
			Ajax.XHR.responseType = "document";
			Ajax.XHR.onloadstart = (ev) => {
				Ajax.Path = uri;
				Ajax.Time = Date.now();
			};
			Ajax.XHR.onload = (ev) => {
				Ajax.Data = ev.target.responseXML;
			};
			Ajax.XHR.open("GET", uri + "?r=i", true);
			Ajax.XHR.send();
		}
	};

	const mouseOver = (ev) => {
		if (isLocal(ev.target.href)) {
			preGET(ev.target.href);
		}
	};

	const mouseDown = (ev) => {
		if (ev.button === 0) {
			if (isLocal(ev.target.href)) {
				if (Ajax.Path !== ev.target.href) {
					preGET(ev.target.href);
				}
				let count = 0;
				let timer = setInterval(() => {
					if (Ajax.Data !== null) {
						clearTimeout(timer);
						document.title = Ajax.Data.title;
						contentHTML.innerHTML = Ajax.Data.body.innerHTML;
						history.pushState(null, null, Ajax.Path);
						aTagInit();
					}
					count++;
					if (count > 300) {
						clearTimeout(timer);
						location.href = Ajax.Path;
					}
				}, 10);
			} else {
				location.href = ev.target.href;
			}
		}
	};

	const aTagInit = () => {
		for (let tag of document.getElementsByTagName("a")) {
			tag.onmouseover = mouseOver;
			tag.onmousedown = mouseDown;
			tag.onclick = () => { return false };
		}
	};

	aTagInit();
};
