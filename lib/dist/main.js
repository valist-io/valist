"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Web3 = void 0;
var getWeb3_1 = require("./getWeb3");
function Web3(provider) {
    return getWeb3_1.setWeb3(provider);
}
exports.Web3 = Web3;
// export default function valist(provider: any) {
//     return setWeb3(provider);
// }
