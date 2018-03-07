function login() {
    var username = $("input[name='username']").val();
    var passwd = $("input[name='passwd']").val();
    $.ajax({
        url: "/user/login/",
        method: "post",
        data: {"username": username, "passwd": passwd},
        success: function (data) {
            var result = JSON.parse(data);
            if(result.Status == "SUCCESS"){
                window.location.href="/index"
            }else{
                $("._error_account_error").show();
            }
        }
    })
}