import sjcl from 'sjcl'

export function encryptPinataKey(pinataKey: string, password: string) {
    let parameters = { "iter" : 1000,};
    let rp = {};

    sjcl.misc.cachedPbkdf2(password, parameters);
    return sjcl.encrypt(password, pinataKey, parameters, rp);
}