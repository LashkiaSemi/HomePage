import {
    FETCH_EQUIPMENTS_SUCCESS,
    CREATE_EQUIPMENT_SUCCESS,
    UPDATE_EQUIPMENT_SUCCESS,
    DELETE_EQUIPMENT_SUCCESS,
} from '../constants/action-types'
import { httpRedirect } from '../util/common'
import { CLIENT_URL } from '../constants/config'

export function equipments(state = [], action) {
    switch (action.type) {
        case FETCH_EQUIPMENTS_SUCCESS:
            return Object.assign([], action.payload.data.equipments)
        case CREATE_EQUIPMENT_SUCCESS:
            // TODO: stateの更新
            httpRedirect(CLIENT_URL + "/admin/equipments")
            return state
        case UPDATE_EQUIPMENT_SUCCESS:
            // TODO: stateの更新
            httpRedirect(CLIENT_URL + "/admin/equipments")
            return state
        case DELETE_EQUIPMENT_SUCCESS:
            // TODO: stateの更新
            httpRedirect(CLIENT_URL + "/admin/equipments")
            return state
        default:
            return state
    }
}