  function btnPostClick() {
    form1 = document.getElementById("form");
    form1.action = "/chatPost";
    form1.submit;
  }
  $(function () {
    $("#openModal").click(function () {
      $("#modalArea").fadeIn();
    });
    $("#closeModal , #modalBg").click(function () {
      $("#modalArea").fadeOut();
    });
    $("#btnRoomCreate").click(function () {
        var friendList="";
        $('input[name=selectFriend]:checked').each(function() {
            var v = $(this).val();
            friendList=friendList+v+","
          });
          friendList=friendList.slice(0,-1);
          $("#selectFriendList").val(friendList);
          $('#formRoomCreate').submit();
      });
  });