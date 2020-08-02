import sjcl from 'sjcl'

export function decryptPinataKey(cipherTextJson: string, password: string) {
    return sjcl.decrypt(password, cipherTextJson)
}
