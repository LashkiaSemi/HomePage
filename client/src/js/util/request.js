import axios from 'axios'
import { showLoading, hideLoading } from '../actions/action'
import store from '../store/store'

// apiにgetリクエストを送信
// stateのisLoadingはここでいじってます
export async function get(url) {
    store.dispatch(showLoading())
    return await axios.get(url).finally(() => store.dispatch(hideLoading()))
}