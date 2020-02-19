import {
    FETCH_RESEARCHES_SUCCESS,
    CREATE_RESEARCH_SUCCESS,
    UPDATE_RESEARCH_SUCCESS,
    DELETE_RESEARCH_SUCCESS,
} from '../constants/action-types'
import { httpRedirect } from '../util/common'
import { CLIENT_URL } from '../constants/config'

export function researches(state = [], action) {
    switch (action.type) {
        case FETCH_RESEARCHES_SUCCESS:
            return Object.assign([], action.payload.data.researches)
        case CREATE_RESEARCH_SUCCESS:
            // TODO: stateの更新
            httpRedirect(CLIENT_URL + "/admin/researches")
            return state
        case UPDATE_RESEARCH_SUCCESS:
            // TODO: stateの更新
            httpRedirect(CLIENT_URL + "/admin/researches")
            return state
        case DELETE_RESEARCH_SUCCESS:
            // TODO: stateの更新
            httpRedirect(CLIENT_URL + "/admin/researches")
            return state
        default:
            return state
    }
}