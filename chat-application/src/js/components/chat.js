import React from "react";
import ChatStore from "../data/chatData";

export default class Chat extends React.Component {
  constructor() {
    super();
    this.state = {
      chats: ChatStore.getAll(),
    };
    var xhr = new XMLHttpRequest();

    xhr.open("GET", "/chatList");
    xhr.send();
    xhr.onreadystatechange = function () {
      if (xhr.readyState === 4 && xhr.status === 200) {
        //チャットデータをそれぞれのdataに格納し、画面表示用のチャットデータに格納する

        var data = JSON.parse(xhr.responseText);
        console.log(this.chats);
        var text = "";
        var name = "";
        var date = "";

        for (let i = 0; i < data.chatList.length; i++) {
          text = data.chatList[i].ChatText;
          name = data.chatList[i].Contributer;
          date = data.chatList[i].BoteDateDisp;
          ChatStore.createChat(text, name, date);
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
    document.getElementById("chat_text").value = "";
  }

  render() {
    const { chats } = this.state;
    const ChatComponents = chats.map((chat) => {
      //チャットの投稿の画面を作成する
      if (chat.name == "aaa") {
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
