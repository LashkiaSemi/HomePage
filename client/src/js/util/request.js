import axios from 'axios'
import { showLoading, hideLoading } from '../actions/action'
import store from '../store/store'

// apiにgetリクエストを送信
// stateのisLoadingはここでいじってます
export async function get(url) {
    const options = { withCredentials: true }
    store.dispatch(showLoading())
    return await axios.get(url, options).finally(() => store.dispatch(hideLoading()))
}

export async function post(url, body) {
    const options = {withCredentials: true}
    store.dispatch(showLoading())
    return await axios.post(url, body, options).finally(() => store.dispatch(hideLoading()))
}

export async function put(url, body) {
    const options = { withCredentials: true }
    store.dispatch(showLoading())
    return await axios.put(url, body,options)
}