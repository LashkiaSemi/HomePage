import { SHOW_LOADING, HIDE_LOADING } from '../constants/action-types'

export function isLoading(state = false, action) {
    switch (action.type) {
        case SHOW_LOADING:
            return true
        case HIDE_LOADING:
            return false
        default:
            return state
    }
}