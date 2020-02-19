import {
    FETCH_ACCOUNT_SUCCESS, FETCH_ACCOUNT_FAILURE,
    UPDATE_ACCOUNT_SUCCESS, UPDATE_ACCOUNT_FAILURE,
    UPDATE_ACCOUNT_PASSWORD_SUCCESS, UPDATE_ACCOUNT_PASSWORD_FAILURE,
} from '../constants/action-types'
import { httpRedirect } from '../util/common'
import { CLIENT_URL } from '../constants/config'

// 複数人で同時にやったらstateが起き変わったりしない？
export function account(state = {}, action) {
    switch (action.type) {
        case FETCH_ACCOUNT_SUCCESS:
            return Object.assign({}, action.payload.data)
        case FETCH_ACCOUNT_FAILURE:
            return state
        case UPDATE_ACCOUNT_SUCCESS:
            // TODO: stateの更新
            httpRedirect(CLIENT_URL + "/members/" + action.payload.data.id)
            return state
        case UPDATE_ACCOUNT_PASSWORD_SUCCESS:
            httpRedirect(CLIENT_URL + "/members/" + action.payload.data.id)
            return state
        case UPDATE_ACCOUNT_PASSWORD_FAILURE:
            console.log("reducer: password update fail")
            return state
        default:
            return state
    }
}