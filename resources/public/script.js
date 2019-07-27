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
            document.getElementById("titleImage").src = e.target.result;
            //document.getElementById("imagePreview").style.display = "block"
        };
        reader.readAsDataURL(input.files[0]);
    }
}

function showLogin() {
    document.getElementById("showLogin").style.display = "none";
    document.getElementById("login").style.display = "inline-block";
}

function preview() {
    document.getElementById("motto").innerText = document.getElementsByName("theme")[0].value;
    var inputDate = new Date(document.getElementsByName("date")[0].value);
    document.getElementById("date").innerText = ('0' + inputDate.getDate()).slice(-2) + "." + ('0' + (inputDate.getMonth() + 1)).slice(-2) + "." + inputDate.getFullYear() + " - " + ('0' + inputDate.getHours()).slice(-2) + ":" + ('0' + inputDate.getMinutes()).slice(-2);
    document.getElementById("vorspeiseText").innerText = document.getElementsByName("starter")[0].value;
    document.getElementById("hauptgangText").innerText = document.getElementsByName("main-dish")[0].value;
    document.getElementById("nachspeiseText").innerText = document.getElementsByName("dessert")[0].value;
    document.getElementById("infoText").innerText = document.getElementsByName("info")[0].value;
}

function updateClosingDate() {
    var tzoffset = (new Date()).getTimezoneOffset() * 60000;
    var inputDate = new Date(document.getElementsByName("date")[0].value);
    console.log(new Date(inputDate - tzoffset - (24 * 60 * 60 * 1000)).toISOString().slice(0, -8));
    document.forms.createEventForm.closingDate.value = new Date(inputDate - tzoffset - (24 * 60 * 60 * 1000)).toISOString().slice(0, -8);
}

document.addEventListener('DOMContentLoaded', testCookie);
