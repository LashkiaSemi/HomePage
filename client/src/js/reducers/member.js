import {
    FETCH_MEMBERS_SUCCESS,
    CREATE_MEMBER_SUCCESS,
    UPDATE_MEMBER_SUCCESS,
    DELETE_MEMBER_SUCCESS,
    FETCH_MEMBER_SUCCESS,
} from '../constants/action-types'
import { httpRedirect } from '../util/common'
import { CLIENT_URL } from '../constants/config'

export function members(state = [], action) {
    switch (action.type) {
        case FETCH_MEMBERS_SUCCESS:
            return Object.assign([], action.payload.data.users)
        case CREATE_MEMBER_SUCCESS:
            // TODO: stateの更新
            httpRedirect(CLIENT_URL + "/admin/members")
            return state
        case UPDATE_MEMBER_SUCCESS:
            // TODO: stateの更新
            httpRedirect(CLIENT_URL + "/admin/members")
            return state
        case DELETE_MEMBER_SUCCESS:
            // TODO: stateの更新
            httpRedirect(CLIENT_URL + "/admin/members")
            return state
        default:
            return state
    }
}

// TODO: membersに統合してしまってもいいかもしらん
export function member(state = {}, action) {
    switch (action.type) {
        case FETCH_MEMBER_SUCCESS:
            return Object.assign({}, action.payload.data)
        case UPDATE_MEMBER_SUCCESS:
            // TODO: stateの更新
            // TODO: redirect -> /admin/members
            console.log("reducer: update success")
            return Object.assign({}, action.payload.data)
        default:
            return state
    }
}