import React from "react";
import ChatStore from "../data/chatData";

export default class Chat extends React.Component {
  constructor() {
    super();
    this.state = {
      chats: ChatStore.getAll(),
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
    ChatStore.createChat(text, "aaa");
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
              <span class="message-data-time">14:56</span>
            </div>
            <div class="message my-message"> {chat.text}</div>
          </li>
        );
      } else {
        //ログインユーザ以外が投稿したチャット
        return (
          <li class="clearfix">
            <div class="message-data other-data">
              <span class="message-data-time">14:56</span>
            </div>
            <div class="message other-message float-right"> {chat.text}</div>
          </li>
        );
      }
    });

    return (
      <div class="chat chat-history">
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
