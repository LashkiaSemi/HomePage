import {
    FETCH_ACTIVITIES_SUCCESS, 
    CREATE_ACTIVITY_SUCCESS,
    UPDATE_ACTIVITY_SUCCESS,
    DELETE_ACTIVITY_SUCCESS,
} from '../constants/action-types'
import { httpRedirect } from '../util/common'
import { CLIENT_URL } from '../constants/config'

// failの時の処理?
export function activities(state = [], action) {
    switch (action.type) {
        case FETCH_ACTIVITIES_SUCCESS:
            return Object.assign([], action.payload.data.activities)
        case CREATE_ACTIVITY_SUCCESS:
            httpRedirect(CLIENT_URL+"/admin/activities")
            return state
        case UPDATE_ACTIVITY_SUCCESS:
            // TODO: stateを更新？
            httpRedirect(CLIENT_URL + "/admin/activities")
            return state
        case DELETE_ACTIVITY_SUCCESS:
            // TODO: stateを更新？
            httpRedirect(CLIENT_URL + "/admin/activities")
            return state
        default:
            return state
    }
}