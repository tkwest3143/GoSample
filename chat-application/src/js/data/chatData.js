import { EventEmitter } from "events";

class ChatStore extends EventEmitter {
  constructor() {
    super();
    this.chats = [];
  }
  createChat(text, name, date) {
    const id = Date.now();

    this.chats.push({
      id,
      text,
      name,
      date,
    });

    this.emit("change");
  }

  getAll() {
    return this.chats;
  }
}

const chatStore = new ChatStore();

export default chatStore;
