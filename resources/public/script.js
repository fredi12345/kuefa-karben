function confirmDeletion(objectToDelete) {
    return confirm("Möchten Sie " + objectToDelete + " wirklich löschen?");
}

function acceptCookies() {
    document.cookie = "acceptedCookies=true";
    document.getElementById("cookieNotice").style.display = "none";
}

//Show cookie notice if not already accepted
function testCookie() {
    var alreadyAccepted = false;

    var ca = document.cookie.split(';');
    for (var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) === ' ') c = c.substring(1, c.length);
        if (c.indexOf("acceptedCookies=") === 0) alreadyAccepted = true;
    }
    if (!alreadyAccepted) document.getElementById("cookieNotice").style.display = "block";
}

document.addEventListener('DOMContentLoaded', testCookie);
