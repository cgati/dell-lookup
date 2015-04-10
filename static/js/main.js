document.addEventListener("keydown", doKeyDown)

function doKeyDown(e) {
    if (e.keyCode == 13) {
        var input = document.getElementById("st-search");
        if (input.value !== "") {
            getServiceTag(input.value);
        }
    }
}

function getServiceTag(sts) {
    var request = new XMLHttpRequest();
    request.open('POST', '/dell/test/', true);
    var dump = document.getElementById("content");
    request.onload = function() {
      if (this.status >= 200 && this.status < 400) {
        // Success!
        var resp = this.response;
        dump.textContent = resp;
      } else {
        // We reached our target server, but it returned an error
        dump.textContent = "An error occurred. Status: " + this.status;
      }
    };

    request.onerror = function() {
      // There was a connection error of some sort
      dump.textContent = "A connection error occurred.";
    };

    request.send(sts);

}
