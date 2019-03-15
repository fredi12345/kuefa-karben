function confirmDeletion(objectToDelete) {
    return confirm("Möchten Sie " + objectToDelete + " wirklich löschen?");
}

function acceptCookies() {
    document.cookie = "acceptedCookies=true; expires=Sat, 31 Dec 2050 23:59:59; path=/";
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

function previewImage(input) {
    if (input.files && input.files[0]) {
        var reader = new FileReader();
        reader.onload = function (e) {
            document.getElementById("imagePreview").src = e.target.result;
            document.getElementById("imagePreview").style.display = "block"
        };
        reader.readAsDataURL(input.files[0]);
    }
}

function showLogin() {
    document.getElementById("showLogin").style.display = "none";
    document.getElementById("login").style.display = "inline-block";
}

document.addEventListener('DOMContentLoaded', testCookie);
