export class RespError {
  constructor(status, code, msg) {
    this.status = status;
    this.code = code;
    this.msg = msg;
  }
}

export const ErrorCode = Object.freeze({
  Undefined: 0,
  ServerError: 1,
  BadParams: 2,
});

export class User {
  constructor(mail, password) {
    this.mail = mail;
    this.password = password;
  }
}

export class Dataset {
  constructor(name, type, upload, description) {
    this.name = name;
    this.type = type;
    this.upload = upload;
    this.description = description;
  }
}

export class Prediction {
  constructor() {

  }
}
