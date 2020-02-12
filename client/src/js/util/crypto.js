import CryptoJS from 'crypto-js'
import { CRYPTO_KEY } from '../constants/config'

const key = CRYPTO_KEY

export const Encrypt = (word) => {
    return CryptoJS.AES.encrypt(word, key).toString()
}

export const Decrypt = (word) => {
    return CryptoJS.AES.decrypt(word, key).toString(CryptoJS.enc.Utf8)
}

// TODO: 使い方的な
// var pwd = "ことば"
// var mm = Encrypt(pwd)
// console.log(mm)

// var jm = Decrypt(mm)
// console.log(jm)