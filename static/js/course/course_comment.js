function expand(currentNode) {
    if($(".comment_form").is(":hidden")){
        $(".comment_form").show();    //如果元素为隐藏,则将它显现
        $(currentNode).html("隐藏");
    }else{
        $(".comment_form").hide();     //如果元素为显现,则将其隐藏
        $(currentNode).html("展开");
    }
}

function reloadComment(parent_id, topic_id, topic_type) {
    $.ajax({
        url:"/comment/topicReply/filter",
        method:"post",
        data:{"parent_id":parent_id,"topic_id":topic_id,"topic_type":topic_type},
        success:function (result) {
            if(result.status == "SUCCESS"){
                // 替换默认分隔符
                $.views.settings.delimiters("[[", "]]");

                $("#topic_reply_area").html("");

                //获取模板
                var jsRenderTpl = $.templates('#topic_reply_template');
                //模板与数据结合
                var finalTpl = jsRenderTpl({"topic_replys":result.topic_replys});
                $('#topic_reply_area').html(finalTpl);
            }
        }
    });
}

function showSubReply(currentNode,parent_id) {
    // 获取 topic_id 和 topic_type
    var topic_id = $("input[name='topic_id'][type='hidden']").val();
    var topic_type = $("input[name='topic_type'][type='hidden']").val();

    $.ajax({
        url:"/comment/topicReply/filter",
        method:"post",
        data:{"parent_id":parent_id,"topic_id":topic_id,"topic_type":topic_type},
        success:function (result) {
            if(result.status == "SUCCESS" && result.topic_replys.length > 0){
                if($("#sub_topic_reply_" + parent_id).is(":hidden")){
                    $("#sub_topic_reply_" + parent_id).show();    //如果元素为隐藏,则将它显现
                    $(currentNode).html("收起回复");

                    // 替换默认分隔符
                    $.views.settings.delimiters("[[", "]]");

                    $("#sub_topic_reply_" + parent_id).html("");

                    //获取模板
                    var jsRenderTpl = $.templates('#topic_reply_template');
                    //模板与数据结合
                    var finalTpl = jsRenderTpl({"topic_replys":result.topic_replys});
                    $("#sub_topic_reply_" + parent_id).html(finalTpl);


                }else{
                    $("#sub_topic_reply_" + parent_id).hide();     //如果元素为显现,则将其隐藏
                    $(currentNode).html("查看所有回复(" + result.topic_replys.length + ")");
                }
            }
        }
    });
}

function addSubReply(currentNode, refer_user_name, parent_id) {
    // 移动到评论锚点
    window.location.href="#reply_anchor";
    $(".comment_form").show();

}

$(function () {
    // 隐藏域获取课程 id
    var course_id = $("input[name='course_id'][type='hidden']").val();
   // 渲染 topicTheme 信息
   $.ajax({
       url:"/comment/topicTheme/filter",
       method:"post",
       data:{"topic_id":course_id,"topic_type":"course_comment"},
       async:false,
       success:function (data) {
           if(data.status == "SUCCESS"){
               new Vue({
                   // 修改 vue 默认分隔符,解决冲突问题
                   delimiters: ['[[', ']]'],
                   el: '#topic_theme',
                   data: {
                       topic_theme: data.topic_theme
                   }
               });
           }
       }
   });

    // 获取 topic_id 和 topic_type
    var topic_id = $("input[name='topic_id'][type='hidden']").val();
    var topic_type = $("input[name='topic_type'][type='hidden']").val();

    $("#submit_comment").click(function () {
        // 获取父评论 id
        var parent_id = $("input[name='parent_id'][type='hidden']").val();
        // 获取评论内容
        var reply_content = $("textarea[name='reply_content']").val();

        // 获取被评论人员
        var refer_user_name = $("input[name='refer_user_name'][type='hidden']").val();

        $.ajax({
            url:"/comment/topicReply/add",
            method:"post",
            data:{"parent_id":parent_id,"reply_content":reply_content,"topic_id":topic_id,"topic_type":topic_type,"refer_user_name":refer_user_name},
            success:function (data) {
                if(data.status == "SUCCESS"){
                    reloadComment(parent_id, topic_id, topic_type);
                }
            }
        });
    });

    // 默认加载 parent_id = 0 根节点评论
    reloadComment(0, topic_id, topic_type);
});