$(function () {

    loadCourse();

});

function bindPopoverEvent() {
    $('#vedio li').each(function (index) {
        $(this).webuiPopover({
            title:function () {
                var courses = $(document).data("courses");
                for(var i=0; i<courses.length; i++){
                    if (courses[i].id == $('#vedio li:eq(' + index + ')').attr("course-id")){
                        var course = courses[i];
                        return course.course_name
                            + "<span style='margin-left: 10px;' class='star' score='" + course.score + "'></span>"
                            + "<span style='float: right;color: red;'>" + course.score +"分</span>";
                    }
                }
                return "";
            },
            content:function () {
                var courses = $(document).data("courses");
                for(var i=0; i<courses.length; i++){
                    if (courses[i].id == $('#vedio li:eq(' + index + ')').attr("course-id")){
                        var course = courses[i];
                        return "<ul><li>"
                            + course.course_status + "："
                            + course.course_number + "集</li><li>类型：<span class='coursetype'>"
                            + course.course_type + "</span></li><li>作者：<span class='courseauthor'>"
                            + course.course_author + "</span></li><li>简介：<span class='courseshortdes'>"
                            + course.course_short_desc + "</span></li><li>"
                            + renderNumber(course.course_number) + "</li></ul>";
                    }
                }
                return "";
            },
            trigger:'hover',
            placement:function () {
                return (index + 1) % 5 == 4 || (index + 1) % 5 == 0 ? "left-bottom" : "right-bottom";
            },
            width:400,
            height:300,
            delay:100,
            onShow: function($element) {
                $element.find('.star').raty({
                    number: 5, // 多少个星星设置
                    hints: ['冷门', '一般', '比较热门', '热门', '非常热门'],
                    score: function(){      // 初始值设置
                        return $(this).attr('score') / 2;
                    },
                    path: "/static/common/raty-2.8.0/lib/images",
                    precision: true, //是否包含小数
                    readOnly:true,
                    // click: function(score, evt) {
                    //     alert('ID: ' + $(this).attr('id') + "\nscore: " + score + "\nevent: " + evt.type);
                    // }
                });
                // 视频集数事件绑定
                $(".jyTable").createTab({marginLeft:-350, time: 10, speed : 'fast'});
            }
        });
    });
}

function loadCourse() {
    $.ajax({
        url:"/course/queryCourse",
        type:"post",
        data:{"offset":15},
        success:function (data) {
            var jsonObj = $.parseJSON(data);
            var courses = jsonObj.courses;
            var html = "";
            for(var i=0; i<courses.length; i++){
                var course = courses[i];
                html += "<li course-id='" + course.id +"'>"
                    + "<a href='#'><img src='" + course.small_image + "'/></a>"
                    + "<dl><dt><a href='#'>" + course.course_name + "</a></dt>"
                    + "<dd>" + course.course_short_desc + "</dd></dl></li>";
            }
            // 拼接 html
            $("#vedio ul").html(html);
            // 缓存数据在 document 上面
            $(document).data("courses", courses);
            // 绑定 popover 事件
            bindPopoverEvent();
        }
    });
}

// 渲染视频集数
function renderNumber(number) {
    var pageSize;
    if(number <=40){
        pageSize = 16;
    }else if(number <= 60){
        pageSize = 24;
    }else {
        pageSize = 32;
    }

    // 向上取整,每页 pageSize 集
    var page = Math.ceil(number / pageSize);

    var funcs = {
        pageHtml:function () {
            if(page == 1){      // 只有一页
                return "<li style='width:50px;font-size:12px;' class='cur'>1-" + number + "</li>";
            }else if(page == 2){
                return "<li style='width:50px;font-size:12px;' class='cur'>1-" + pageSize + "</li><li style='width:50px;font-size:12px;'>" + (pageSize+1) + "-" + number +"</li>";
            }else{
                var pageHtml = "";
                for(var i=0; i<page; i++){
                    if(i == 0){
                        pageHtml += "<li style='width:50px;font-size:12px;' class='cur'>1-" + pageSize + "</li>";
                    }else if(i == page - 1){
                        pageHtml += "<li style='width:50px;font-size:12px;'>" + ((page - 1) * pageSize + 1) + "-" + number +"</li>";
                    }else{
                        pageHtml += "<li style='width:50px;font-size:12px;'>" + (i * pageSize + 1) + "-" + ((i + 1) * pageSize) +"</li>";
                    }
                }
                return pageHtml;
            }
        },
        renderPageDetail:function (start,end) {
            var html = '<div class="tabCon" style="width:350px;height:100px;border:none;">';
            for(var i=start; i<=end; i++){
                html += '<a href="#" style="display: block;width: 40px;height: 20px;background: #e8e8e8;float: left;margin: 1px;text-align: center;">' + i +'</a>';
            }
            html += '</div>';
            return html;
        },
        tabCon:function () {
            if(page == 1){      // 只有一页
                return funcs.renderPageDetail(1, number);
            }else if(page == 2){
                return funcs.renderPageDetail(1, pageSize) + funcs.renderPageDetail((pageSize+1), number);
            }else{
                var pageDetailHtml = "";
                for(var i=0; i<page; i++){
                    if(i == 0){
                        pageDetailHtml += funcs.renderPageDetail(1, pageSize);
                    }else if(i == page - 1){
                        pageDetailHtml += funcs.renderPageDetail((page - 1) * pageSize + 1, number);
                    }else{
                        pageDetailHtml += funcs.renderPageDetail(i * pageSize + 1, (i + 1) * pageSize);
                    }
                }
                return pageDetailHtml;
            }
        }
    };

    var html = '    <div style="width:350px;margin: 0 auto;">\n' +
        '        <h1 class="titleH1 underNone clearfix">\n' +
        '            <span class="left underNone underLine" style="font-size: 15px;">剧集信息</span>\n' +
        '        </h1>\n' +
        '        <div class="jyTable">\n' +
        '            <div class="clearfix">\n' +
        '                <ul class="title title1 left">\n' +
                            funcs.pageHtml() +
        '                </ul>\n' +
        '            </div>\n' +
        '            <div class=\'zong\' style="width:350px;height:100px;">\n' +
        '                <div class="list list1">\n' +
                            funcs.tabCon() +
        '                </div>\n' +
        '            </div>\n' +
        '        </div>\n' +
        '    </div>\n';

    return html;
}

