import { EventEmitter } from "events";

class ChatStore extends EventEmitter {
  constructor() {
    super();
    this.chats = [];
  }
  createChat(text, name, date, myBoteFlg) {
    const id = Date.now();

    this.chats.push({
      id,
      text,
      name,
      date,
      myBoteFlg,
    });

    this.emit("change");
  }

  getAll() {
    return this.chats;
  }
}

const chatStore = new ChatStore();

export default chatStore;
