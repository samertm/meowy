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

$("button.delete").each(function(i) {
    $(this).attr("num", i)
});

$("button.promote").each(function(i) {
    $(this).attr("num", i+1) // i+1 because there is no promote button for the first value
    console.log(this)
});

$("button").click(function() {
    if ($(this).hasClass("delete")) {
        var toSend = "delete=" + $(this).attr("num")
        $.ajax({
            url: "thing/delete",
            type: "POST",
            data: toSend,
            success: function() {
                location.reload(true);
            },
            error: function() {
                console.error("errored on ajax request");
            }
        });
    } else if ($(this).hasClass("promote")) {
        var toSend = "promote=" + $(this).attr("num")
        $.ajax({
            url: "thing/promote",
            type: "POST",
            data: toSend,
            success: function() {
                location.reload(true);
            },
            error: function() {
                console.error("errored on ajax request");
            }
        });
    }
});
