var hori_animax = document.querySelector("div.hori_animax_line");
addEventListener("scroll", function() {
  var max = document.body.scrollHeight - innerHeight;
  var percent = (pageYOffset / max) * 100;
  hori_animax.style.width = percent + "%";
});

var notific = document.querySelector("div.control_panel");
var notific_show = document.getElementById("notific_show_button");
var notific_hide = document.getElementById("notific_hide_button");

notific_show.addEventListener("click", function() {
  notific.style.cssText = "visibility: visible; width: 240px;";
});
notific_hide.addEventListener("click", function() {
  notific.style.cssText = "width: 0px; visibility: hidden;";
});
