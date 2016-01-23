/* show and hide more services from kalibu*/
var apps_show = document.getElementById("apps_button");
var apps_hide = document.getElementById("apps_close_button");
var apps_grid = document.querySelector("div.kalibu_apps");
apps_show.addEventListener("click", function() {
	apps_grid.style.cssText = "visibility: hidden;";
	apps_button.style.cssText = "display: none;";
	apps_hide.style.cssText = "display: block;";
});
apps_hide.addEventListener("click", function() {
	apps_grid.style.cssText = "visibility: visible;";
	apps_hide.style.cssText = "display: none;";
	apps_button.style.cssText = "display: block;";
});
