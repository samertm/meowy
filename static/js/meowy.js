;
// tesing

$("#work").submit(function(e) {
    e.preventDefault();
    $("#err").text("");
    var data = $(this).serializeArray();
    var valid = true;
    for (var i = 0; i < data.length; i++) {
        if (data[i].value == "") {
            valid = false;
        }
    }
    if (!valid) {
        $("#err").text("error: field blank");
        return;
    }
    getSession(); // creates sessionid cookie if it doesn't exist
    var toSend = $(this).serialize();
    $.ajax({
        url: "thing/change",
        type: "POST",
        data: toSend,
        success: function() {
            location.reload(true);
        },
        error: function() {
            console.error("errored on ajax request")
        }
    });
});
