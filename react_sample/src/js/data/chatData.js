import { EventEmitter } from "events";

class ChatStore extends EventEmitter {
  constructor() {
    super();
    this.chats = [
      {
        id: 113464613,
        text: "Go Shopping",
        name: "aaa",
      },
      {
        id: 235684679,
        text: "Pay Water Bills",
        name: "bbb",
      },
    ];
  }
  createChat(text, name) {
    const id = Date.now();

    this.chats.push({
      id,
      text,
      name,
    });

    this.emit("change");
  }

  getAll() {
    return this.chats;
  }
}

const chatStore = new ChatStore();

export default chatStore;
