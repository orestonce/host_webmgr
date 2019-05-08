function ymd_send_and_reload(urlString) {
    var request = new XMLHttpRequest();
    request.onreadystatechange = function () {
        if (request.readyState === 4) {
            if (request.status === 200) {
                location.reload(true);
            } else {
                alert("State error [" + request.status + "]: " + request.responseText);
            }
        }
    };
    request.open("GET", urlString);
    request.send();
}