document.addEventListener("keydown", doKeyDown);

var button = document.getElementById("st-button");
button.addEventListener('click', function() {
    sendData();
}, false);

function doKeyDown(e) {
    if (e.keyCode == 13) {
        sendData();
    }
}

function sendData() {
    var input = document.getElementById("st-search");
    if (input.value !== "") {
        getServiceTag(input.value);
    }
}

function getServiceTag(sts) {
    var request = new XMLHttpRequest();
    request.open('POST', '/dell/test/', true);
    var dump = document.getElementById("content");
    request.onload = function() {
      if (this.status >= 200 && this.status < 400) {
        // Success!
        renderResult(this.response);
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

function renderResult(data) {
    var content = document.getElementById("content");
    var rendered = document.getElementById("assetList");
    if (rendered !== null) {
        content.removeChild(rendered);
    }
    var dellResult = JSON.parse(data);
    var assetList = document.createElement("div");
    assetList.setAttribute("id", "assetList");
    for (a of dellResult.Assets) {
        var asset = document.createElement("div");
        var assetSpan = document.createElement("span");
        assetSpan.textContent = a.ServiceTag + " - " + a.MachineDescription;
        asset.appendChild(assetSpan);
        var wTable = document.createElement("table");
		wTable.setAttribute("class","warrantyTable");
        var tHead = document.createElement("thead");
        var tHeadRow = document.createElement("tr");
        var thDescription = document.createElement("th");
        thDescription.textContent = "Description";
        var thStart = document.createElement("th");
        thStart.textContent = "Start Date";
        var thEnd = document.createElement("th");
        thEnd.textContent = "End Date";
        var thInWarranty = document.createElement("th");
        thInWarranty.textContent = "Warranty Valid";
        tHeadRow.appendChild(thDescription);
        tHeadRow.appendChild(thStart);
        tHeadRow.appendChild(thEnd);
        tHeadRow.appendChild(thInWarranty);
        tHead.appendChild(tHeadRow);
        wTable.appendChild(tHead);
        var tBody = document.createElement("tbody");
        for (w of a.Warranties.Warranty) {
            var tBodyRow = document.createElement("tr");
            var tdDescription = document.createElement("td");
            tdDescription.textContent = w.ServiceLevelDescription;
            var tdStartDate = document.createElement("td");
            tdStartDate.textContent = w.StartDate;
            var tdEndDate = document.createElement("td");
            tdEndDate.textContent = w.EndDate;
            var tdInWarranty = document.createElement("td");
            tdInWarranty.textContent = w.InWarranty ? "Yes." : "No.";
            tBodyRow.appendChild(tdDescription);
            tBodyRow.appendChild(tdStartDate);
            tBodyRow.appendChild(tdEndDate);
            tBodyRow.appendChild(tdInWarranty);
            tBody.appendChild(tBodyRow);
        }
        wTable.appendChild(tBody);
        asset.appendChild(wTable);
        assetList.appendChild(asset);
        var brk = document.createElement("br");
        assetList.appendChild(brk);
    }
    content.appendChild(assetList);
}
