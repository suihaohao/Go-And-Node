const text = "123456text"
const sha1 = require("sha1");

const sha1Sum = sha1(text);
console.log(sha1Sum)
