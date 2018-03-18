$(function () {
    $.ajax({
        url:"/course/queryCourse?filterType=courselist",
        method:"post",
        data:{},
        async: false,   // 执行完了之后在执行 modalEffects.js 中的渲染
        success:function (data) {
            var obj = JSON.parse(data);
            new Vue({
                // 修改 vue 默认分隔符,解决冲突问题
                delimiters: ['[[', ']]'],
                el: '#course_list',
                data: {
                    courses: obj.courses
                }
            });
            renderStar();
        }
    });
});

// html5 实现图片预览功能
$(function () {
    $(".changeImageFile").change(function (e) {
        var file = e.target.files[0] || e.dataTransfer.files[0];
        if (file) {
            var reader = new FileReader();
            reader.onload = function () {
                $(".changeImage").attr("src", this.result);
            }

            reader.readAsDataURL(file);
        }
    });
})

// Jquery 实现文件上传操作
function changeImage(node, courseId) {
    // jquery 对象转 dom 对象,再通过 files[0] 属性获取文件
    // document.getElementById('.changeImageFile').files[0]
    var file = $(node).parent().find(".changeImageFile").get(0).files[0];
    var formData = new FormData();
    formData.append('id',courseId);
    formData.append('file', file);

    $.ajax({
        url: "/course/newcourse/changeImage",
        type: "post",
        data: formData,
        contentType: false,
        processData: false,
        mimeType: "multipart/form-data",
        success: function (data) {
            var obj = JSON.parse(data);
            if(obj.status == "SUCCESS"){
                // 页面刷新
                location.reload();
            }
        },
        error: function (data) {
            console.log(data);
        }
    });
}

// 渲染得分情况
function renderStar() {
    $(".star").raty({
        number: 5, // 多少个星星设置
        hints: ['冷门', '一般', '比较热门', '热门', '非常热门'],
        score: function(){      // 初始值设置
            return $(this).attr('score') / 2;
        },
        path: "/static/common/raty-2.8.0/lib/images",
        precision: true, //是否包含小数
        readOnly:true
    });
}