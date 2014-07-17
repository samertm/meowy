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

$("button.adddata").each(function(i) {
    $(this).attr("num", i)
});


$("button.delete").each(function(i) {
    $(this).attr("num", i)
});

$("button.promote").each(function(i) {
    $(this).attr("num", i+1) // i+1 because there is no promote button for the first value
    console.log(this)
});

$("button").click(function() {
    var toSend = ""
    var url = ""
    if ($(this).hasClass("delete")) {
        toSend = "delete=" + $(this).attr("num");
        url = "thing/delete";
    } else if ($(this).hasClass("promote")) {
        toSend = "promote=" + $(this).attr("num");
        url = "thing/promote";
    } else if ($(this).hasClass("adddata")) {
        
        // toSend = "id=" + $(this).attr("num");
        //     "data=" + $(this).attr("data");
        // url = "thing/adddata";
    } else if (toSend == "" || url == "") {
        return;
    }
    sendData(url, toSend);
});

// reloads on successful send
function sendData(url, toSend) {
    $.ajax({
        url: url,
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

// turn all links into hyperlinks
$(".thing").each(function() {
    re = new RegExp("(https?://[^ \n!]+)")
    newText = $(this).html().replace(re, "<a href='$1'>$1</a>")
    $(this).html(newText)
});
