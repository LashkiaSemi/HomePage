import {
    FETCH_TAGS_SUCCESS,
    CREATE_TAG_SUCCESS,
    UPDATE_TAG_SUCCESS,
    DELETE_TAG_SUCCESS,
} from '../constants/action-types'
import { httpRedirect } from '../util/common'
import { CLIENT_URL } from '../constants/config'

export function tags(state = [], action) {
    switch (action.type) {
        case FETCH_TAGS_SUCCESS:
            return Object.assign([], action.payload.data.tags)
        case CREATE_TAG_SUCCESS:
            // TODO: stateの更新
            httpRedirect(CLIENT_URL + "/admin/tags")
            return state
        case UPDATE_TAG_SUCCESS:
            // TODO: stateの更新
            httpRedirect(CLIENT_URL + "/admin/tags")
            return state
        case DELETE_TAG_SUCCESS:
            // TODO: stateの更新
            httpRedirect(CLIENT_URL + "/admin/tags")
            return state
        default:
            return state
    }
}
