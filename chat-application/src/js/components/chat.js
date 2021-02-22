import React from "react";
import ChatStore from "../data/chatData";
import UserStore from "../data/user";

export default class Chat extends React.Component {
  constructor() {
    super();
    this.state = {
      chats: ChatStore.getAll(),
      user: UserStore.getUser(),
    };
    var xhr = new XMLHttpRequest();

    xhr.open("GET", "/chatList");
    xhr.send();
    const { user } = this.state;
    xhr.onreadystatechange = function () {
      if (xhr.readyState === 4 && xhr.status === 200) {
        //チャットデータをそれぞれのdataに格納し、画面表示用のチャットデータに格納する
        var data = JSON.parse(xhr.responseText);
        var text = "";
        var name = "";
        var date = "";
        var myBoteFlg = false;
        for (let i = 0; i < data.chatList.length; i++) {
          text = data.chatList[i].ChatText;
          name = data.chatList[i].Contributer;
          date = data.chatList[i].BoteDateDisp;
          console.log(data.userId + "=" + data.chatList[i].UserID);
          myBoteFlg = data.userId == data.chatList[i].UserID;
          ChatStore.createChat(text, name, date, myBoteFlg);
        }
      }
    };
  }

  /**
   * チャットのデータが変更された場合、チャットをすべて取得する
   */
  componentDidMount() {
    ChatStore.on("change", () => {
      this.setState({
        chats: ChatStore.getAll(),
      });
    });
  }
  //チャット投稿イベント
  chatSend() {
    const text = document.getElementById("chat_text").value;
    var xhr = new XMLHttpRequest();

    xhr.open("GET", "/chatList");
    xhr.send();

    xhr.onreadystatechange = function () {
      if (xhr.readyState === 4 && xhr.status === 200) {
        //チャットデータをそれぞれのdataに格納し、画面表示用のチャットデータに格納する
        var data = JSON.parse(xhr.responseText);
        var text = "";
        var name = "";
        var date = "";
        var myBoteFlg = false;
        for (let i = 0; i < data.chatList.length; i++) {
          text = data.chatList[i].ChatText;
          name = data.chatList[i].Contributer;
          date = data.chatList[i].BoteDateDisp;
          myBoteFlg = data.userId == data.chatList[i].Contributer;
          ChatStore.createChat(text, name, date, myBoteFlg);
        }
      }
    };
    document.getElementById("chat_text").value = "";
  }

  render() {
    const { chats } = this.state;
    const ChatComponents = chats.map((chat) => {
      //チャットの投稿の画面を作成する
      if (chat.myBoteFlg) {
        //ログインユーザが投稿したチャット
        return (
          <li class="clearfix">
            <div class="message-data my-data">
              <span class="message-data-time">{chat.date}</span>
            </div>
            <div class="message my-message"> {chat.text}</div>
          </li>
        );
      } else {
        //ログインユーザ以外が投稿したチャット
        return (
          <li class="clearfix">
            <div class="message-data other-data">
              <span class="message-data-time">{chat.date}</span>
            </div>
            <div class="message other-message float-right"> {chat.text}</div>
          </li>
        );
      }
    });

    return (
      <div class="chat chat-history asd">
        <div class="chat-view">
          <ul class="chat-view">{ChatComponents}</ul>
        </div>
        <div class="chat-message clearfix">
          <textarea
            name="chatText"
            id="chat_text"
            placeholder="メッセージを入力してください"
            rows="3"
          ></textarea>

          <button id="send" onClick={this.chatSend.bind(this)}>
            send
          </button>
        </div>
      </div>
    );
  }
}
