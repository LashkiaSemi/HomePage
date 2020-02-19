import {
    FETCH_SOCIETIES_SUCCESS,
    CREATE_SOCIETY_SUCCESS,
    UPDATE_SOCIETY_SUCCESS,
    DELETE_SOCIETY_SUCCESS,
} from '../constants/action-types'
import { httpRedirect } from '../util/common'
import { CLIENT_URL } from '../constants/config'

export function societies(state = [], action) {
    switch (action.type) {
        case FETCH_SOCIETIES_SUCCESS:
            return Object.assign([], action.payload.data.societies)
        case CREATE_SOCIETY_SUCCESS:
            // TODO: stateのこうしん
            httpRedirect(CLIENT_URL + "/admin/societies")
            return state
        case UPDATE_SOCIETY_SUCCESS:
            // TODO: stateの更新
            httpRedirect(CLIENT_URL + "/admin/societies")
            return state
        case DELETE_SOCIETY_SUCCESS:
            // TODO: stateの更新
            httpRedirect(CLIENT_URL + "/admin/societies")
            return state
        default:
            return state
    }
}