import axios from 'axios'
import { showLoading, hideLoading } from '../actions/action'
import store from '../store/store'

// TODO: optionのカスタムできた方がいいかも
// TODO: deleteが予約後の関係でhttpDeleteになったので、他もそれに倣った方がいいかも

// apiにgetリクエストを送信
// stateのisLoadingはここでいじってます
export async function get(url) {
    const options = { withCredentials: true }
    store.dispatch(showLoading())
    return await axios.get(url, options).finally(() => store.dispatch(hideLoading()))
}

// urlに向けてpostリクエスト
export async function post(url, body, options) {
    store.dispatch(showLoading())
    return await axios.post(url, body, options).finally(() => store.dispatch(hideLoading()))
}

// urlに向けてput
export async function put(url, body) {
    const options = { withCredentials: true }
    store.dispatch(showLoading())
    return await axios.put(url, body,options)
}

// urlに向けてdelete。めっちゃ不恰好やん。
export async function httpDelete(url) {
    const options = { withCredentials: true }
    store.dispatch(showLoading())
    return await axios.delete(url, options)
}