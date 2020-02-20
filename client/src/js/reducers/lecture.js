import {
    FETCH_LECTURES_SUCCESS,
    FETCH_LECTURE_SUCCESS,
    CREATE_LECTURE_SUCCESS,
    UPDATE_LECTURE_SUCCESS,
    DELETE_LECTURE_SUCCESS,
    ADMIN_CREATE_LECTURE_SUCCESS,
    ADMIN_UPDATE_LECTURE_SUCCESS,
    ADMIN_DELETE_LECTURE_SUCCESS,
} from '../constants/action-types'
import { httpRedirect } from '../util/common'
import { CLIENT_URL } from '../constants/config'

export function lectures(state = [], action) {
    switch (action.type) {
        case FETCH_LECTURES_SUCCESS:
            return Object.assign([], action.payload.data.lectures)
        case CREATE_LECTURE_SUCCESS:
            // TODO; stateの更新
            httpRedirect(CLIENT_URL + "/lectures")
            return state
        case UPDATE_LECTURE_SUCCESS:
            // TODO; stateの更新
            httpRedirect(CLIENT_URL + "/lectures")
            return state
        case DELETE_LECTURE_SUCCESS:
            // TODO; stateの更新
            httpRedirect(CLIENT_URL + "/lectures")
            return state

        // admin
        case ADMIN_CREATE_LECTURE_SUCCESS:
            // TODO; stateの更新
            httpRedirect(CLIENT_URL + "/admin/lectures")
            return state
        case ADMIN_UPDATE_LECTURE_SUCCESS:
            // TODO; stateの更新
            httpRedirect(CLIENT_URL + "/admin/lectures")
            return state
        case ADMIN_DELETE_LECTURE_SUCCESS:
            // TODO; stateの更新
            httpRedirect(CLIENT_URL + "/admin/lectures")
            return state
        default:
            return state
    }
}

export function lecture(state = {}, action) {
    switch (action.type) {
        case FETCH_LECTURE_SUCCESS:
            return Object.assign({}, action.payload.data)

        default:
            return state
    }
}
