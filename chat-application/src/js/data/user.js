import { EventEmitter } from "events";

class UserStore extends EventEmitter {
  constructor() {
    super();
    this.user = new Array();
    this.error = new Array();
    this.loginFlg = 0;
    this.userid;
  }
  setUser(username, userId) {
    this.user.push({
      userId,
      username,
    });
    this.userid = userId;
    this.loginFlg = 1;

    this.emit("change");
  }
  setError(err) {
    this.error.push({
      err,
    });

    this.emit("change");
  }

  getUser() {
    return this.user;
  }
  getLoginFlg() {
    return this.loginFlg;
  }
  getUserId() {
    return this.userid;
  }
  getError() {
    return this.error;
  }
}

const userStore = new UserStore();

export default userStore;
