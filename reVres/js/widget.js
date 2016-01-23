var decl = document.getElementById("decline");
var canl = document.getElementById("cancel");
var widget = document.querySelector("div.accept_widget_cont");
decl.addEventListener("click", function() {
  widget.style.cssText = "display: block;";
});
canl.addEventListener("click", function() {
  widget.style.cssText = "display: none;";
});
