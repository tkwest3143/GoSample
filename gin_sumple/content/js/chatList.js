function btnPostClick() {
  form1 = document.getElementById("form");
  form1.action = "/chatPost";
  form1.submit;
}
function hankaku2Zenkaku(str) {
  return str.replace(/[Ａ-Ｚａ-ｚ０-９]/g, function (s) {
    return String.fromCharCode(s.charCodeAt(0) - 0xfee0);
  });
}
$(function () {
  $("#txt-search").blur(function () {
    // 変数の宣言
    let searchResult, targetText;
    // 検索結果を格納するための配列を用意
    searchResult = [];
    var searchtxt = $("#txt-search").val();
    if (searchtxt != "") {
      searchtxt = searchtxt.toLowerCase();
      $(".friend-list").each(function () {
        let isShow = false;
        $(this)
          .find(".about")
          .each(function () {
            targetText = $(this).children(".target-name").val();
            // 検索対象となるリストに入力された文字列が存在するかどうかを判断
            if (targetText.indexOf(searchtxt) !== -1) {
              // 存在する場合はそのリストのテキストを用意した配列に格納
              searchResult.push(targetText);
              isShow = true;
            }
          });

        if (!isShow) {
          $(this).hide();
        } else {
          $(this).show();
        }
      });
    } else {
      // 検索ボックスが空の場合
      $(".friend-list").each(function () {
        // 非表示のものを表示させる。
        if ($(this).css("display") !== "block") {
          $(this).show();
        }
      });
    }
  });

  $("#openModal").click(function () {
    $("#modalArea").fadeIn();
    $("#modalArea").removeClass(".hidden");
  });
  $("#closeModal , #modalBg").click(function () {
    $("#modalArea").fadeOut();
  });
  $("#btnRoomCreate").click(function () {
    var friendList = "";
    $("input[name=selectFriend]:checked").each(function () {
      var v = $(this).val();
      friendList = friendList + v + ",";
    });
    friendList = friendList.slice(0, -1);
    $("#selectFriendList").val(friendList);
    $("#formRoomCreate").submit();
  });
});
