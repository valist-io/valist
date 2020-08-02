"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.decryptPinataKey = void 0;
var sjcl_1 = __importDefault(require("sjcl"));
function decryptPinataKey(cipherTextJson, password) {
    return sjcl_1.default.decrypt(password, cipherTextJson);
}
exports.decryptPinataKey = decryptPinataKey;
