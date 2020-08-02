"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.setWeb3 = void 0;
var web3_1 = __importDefault(require("web3"));
function setWeb3(provider) {
    var web3 = new web3_1.default(provider);
    return web3;
}
exports.setWeb3 = setWeb3;
// const resolveWeb3 = (resolve: any) => {
//     window.web3 = new 
//     const alreadyInjected = typeof web3 !== 'undefined' // i.e. Mist/Metamask
//     const localProvider = `http://localhost:9545`
//     if (alreadyInjected) {
//         console.log(`Injected web3 detected.`)
//         web3 = new Web3(web3.currentProvider)
//     } else {
//         console.log(`No web3 instance injected, using Local web3.`)
//         const provider = new Web3.providers.HttpProvider(localProvider)
//         web3 = new Web3(provider)
//     }
//     resolve(web3)
// }
// export default () =>
//     new Promise((resolve) => {
//         // Wait for loading completion to avoid race conditions with web3 injection timing.
//         window.addEventListener(`load`, () => {
//         resolveWeb3(resolve)
//     })
//     // If document has loaded already, try to get Web3 immediately.
//     if (document.readyState === `complete`) {
//         resolveWeb3(resolve)
//     }
// })
