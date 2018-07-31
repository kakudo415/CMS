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
		if (isLocal(ev.target.href)) {
			ajaxGET(ev.target.href);
		}
	};

	const mouseDown = (ev) => {
		if (ev.button == 0) {
			if (isLocal(ev.target.href)) {
				let count = 0;
				let timer = setInterval(() => {
					if (Ajax.Path == ev.target.href && Ajax.Data != null) {
						movePage();
						clearTimeout(timer);
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

	const isLocal = (URL) => {
		return URL.startsWith("http://" + document.domain) || URL.startsWith("https://" + document.domain) || URL.startsWith("//" + document.domain) || URL.startsWith("/");
	};

	for (let tag of document.getElementsByTagName("a")) {
		tag.onmouseover = mouseOver;
		tag.onmousedown = mouseDown;
		tag.onclick = () => { return false };
	}
};