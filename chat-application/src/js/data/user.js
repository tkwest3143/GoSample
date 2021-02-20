import { EventEmitter } from "events";

class UserStore extends EventEmitter {
  constructor() {
    super();
    this.user = [];
  }
  setUsername(username) {
    this.user.push({
      username,
    });

    this.emit("change");
  }

  getUsername() {
    return this.user;
  }
}

const userStore = new UserStore();

export default userStore;
