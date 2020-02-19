import {
    FETCH_JOBS_SUCCESS,
    CREATE_JOB_SUCCESS,
    UPDATE_JOB_SUCCESS,
    DELETE_JOB_SUCCESS,
} from '../constants/action-types'
import { httpRedirect } from '../util/common'
import { CLIENT_URL } from '../constants/config'

export function jobs(state = [], action) {
    switch (action.type) {
        case FETCH_JOBS_SUCCESS:
            return Object.assign([], action.payload.data.jobs)
        case CREATE_JOB_SUCCESS:
            // TODO: stateの更新
            httpRedirect(CLIENT_URL + "/admin/jobs")
            return state
        case UPDATE_JOB_SUCCESS:
            // TODO: stateの更新
            httpRedirect(CLIENT_URL + "/admin/jobs")
            return state
        case DELETE_JOB_SUCCESS:
            // TODO: stateの更新
            httpRedirect(CLIENT_URL + "/admin/jobs")
            return state
        default:
            return state
    }
}