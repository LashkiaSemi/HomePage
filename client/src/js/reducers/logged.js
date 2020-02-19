import { LOGIN_SUCCESS, LOGIN_FAILURE, LOGOUT_SUCCESS, LOGOUT_FAILURE } from '../constants/action-types'
import { CLIENT_URL, STRAGE_KEY } from '../constants/config'
import * as Crypto from '../util/crypto'
import { httpRedirect } from '../util/common'

export function logged(state = [], action) {
    switch (action.type) {
        case LOGIN_SUCCESS:
            // localstrageに入れる値の暗号化
            const strageValue = Crypto.Encrypt(action.payload.data.user_id + action.payload.data.role)
            // localstrageにタンク
            localStorage.setItem(STRAGE_KEY, strageValue)
            httpRedirect(CLIENT_URL)
            return Object.assign([], state, state.concat({ user_id: action.payload.data.user_id }))
        
        case LOGOUT_SUCCESS:
            // localstarageを消せ！
            localStorage.removeItem(STRAGE_KEY)
            httpRedirect(CLIENT_URL)
            return state

        case LOGOUT_FAILURE:
            // TODO: これいる...?
            if (localStorage.getItem(STRAGE_KEY)) {
                localStorage.removeItem(STRAGE_KEY)
            }
            httpRedirect(CLIENT_URL)
            return state
        default:
            return state
    }
}