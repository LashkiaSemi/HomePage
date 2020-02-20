import CryptoJS from 'crypto-js'
import { CRYPTO_KEY } from '../constants/config'

const key = CRYPTO_KEY

// Encrypt 暗号化
export const Encrypt = (word) => {
    return CryptoJS.AES.encrypt(word, key).toString()
}

// Decrypt 復号
export const Decrypt = (word) => {
    return CryptoJS.AES.decrypt(word, key).toString(CryptoJS.enc.Utf8)
}

// note: 使い方的な
// var pwd = "ことば"
// var mm = Encrypt(pwd)
// console.log(mm)

// var jm = Decrypt(mm)
// console.log(jm)