var output = document.querySelector("div.stdout_text");
var req = new XMLHttpRequest();
function stdout() {
	var q = "std";
	var xml = "view?stdout=" + q;
	console.log(q)
	req.open("GET", xml, true);
	req.addEventListener("load", function(){
	console.log("Ok", req.status);
	})

	req.onreadystatechange = function() {
	if (req.status == 200) {
		var str = req.responseText;
		output.textContent = str;
		}
	}

	req.send(null);
}

setTimeout(stdout, 2000)
