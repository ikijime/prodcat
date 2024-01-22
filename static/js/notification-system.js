htmx.on("showMessage", (e) => {
    var message = e.detail.message
    var messageType = e.detail.messageType

    console.log(messageType)
    
    if (messageType == 'error') {
        Toastify({
            text: message,
            duration: 3000,
            destination: "/",
            newWindow: true,
            close: true,
            gravity: "bottom", // `top` or `bottom`
            position: "right", // `left`, `center` or `right`
            stopOnFocus: true, // Prevents dismissing of toast on hover
            style: {
                background: "linear-gradient(to right, #990000, #550000)",
            },
        onClick: function(){} // Callback after click
        }).showToast()
    }

    if (messageType == 'success') {
        Toastify({
            text: message,
            duration: 3000,
            destination: "/",
            newWindow: true,
            close: true,
            gravity: "bottom", // `top` or `bottom`
            position: "right", // `left`, `center` or `right`
            stopOnFocus: true, // Prevents dismissing of toast on hover
            style: {
                background: "linear-gradient(to right, #009900, #005500)",
            },
        onClick: function(){} // Callback after click
        }).showToast()
    }
});