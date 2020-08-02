"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.getPinataKey = void 0;
var sjcl_1 = __importDefault(require("sjcl"));
function getPinataKey(password) {
    var password = "password";
    var text = "my secret text";
    var parameters = { "iter": 1000 };
    var rp = {};
    var cipherTextJson = {};
    sjcl_1.default.misc.cachedPbkdf2(password, parameters);
    cipherTextJson = sjcl_1.default.encrypt(password, text, parameters, rp);
    console.log(cipherTextJson);
    var decryptedText = sjcl_1.default.decrypt(password, cipherTextJson);
    console.log(decryptedText);
}
exports.getPinataKey = getPinataKey;
