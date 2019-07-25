const text = "123456text"
const sha1 = require("sha1");

const sha1Sum = sha1(text);
console.log(sha1Sum)



//加盐哈希
const crypto = require('crypto');
const salt = "e9a51074898fb36d";
const hash = crypto.createHmac('sha256', salt);
hash.update(text);
const hmacResult = hash.digest('hex');

console.log(hmacResult)
