var req = new XMLHttpRequest();


var output = document.querySelector("pre.stdout_text");
function stdout() {
	var q = "std";
	var xml = "home?stdout=" + q;
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

setInterval(stdout, 2000)

var kill_but = document.getElementById("kill_button");
var pid = document.getElementById("kill_pid");
var pid_lab = document.getElementById("kill_label");
kill_but.addEventListener("click", function() {
	var q = pid.value;
	var xml = "home?term=" + q;
	console.log(q)
	req.open("GET", xml, true);
	req.addEventListener("load", function(){
	console.log("Ok", req.status);
	})

	req.onreadystatechange = function() {
	if (req.status == 200) {
		var str = req.responseText;
		pid_lab.textContent = str;
		}
	}

	req.send(null);
});


var cmd_but = document.getElementById("cmd_button");
var cmd = document.getElementById("cmd");
var cmd_stdout = document.querySelector("pre.command_text")
cmd_but.addEventListener("click", function() {
	var c = cmd.value;
	var xml = "home?cmd=" + c;
	console.log(c)
	req.open("GET", xml, true);
	req.addEventListener("load", function(){
	console.log("Ok", req.status);
	})

	req.onreadystatechange = function() {
	if (req.status == 200) {
		var str = req.responseText;
		cmd_stdout.textContent = str;
		}
	}

	req.send(null);
});
