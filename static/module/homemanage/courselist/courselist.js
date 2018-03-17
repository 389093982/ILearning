$(function () {
    $(".star").raty({
        number: 5, // 多少个星星设置
        hints: ['冷门', '一般', '比较热门', '热门', '非常热门'],
        score: function(){      // 初始值设置
            // return $(this).attr('score') / 2;
            return 5;
        },
        path: "/static/common/raty-2.8.0/lib/images",
        precision: true, //是否包含小数
        readOnly:true
    });


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
            })
        }
    });
});