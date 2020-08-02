"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.encryptPinataKey = void 0;
var sjcl_1 = __importDefault(require("sjcl"));
function encryptPinataKey(pinataKey, password) {
    var parameters = { "iter": 1000, };
    var rp = {};
    sjcl_1.default.misc.cachedPbkdf2(password, parameters);
    return sjcl_1.default.encrypt(password, pinataKey, parameters, rp);
}
exports.encryptPinataKey = encryptPinataKey;
