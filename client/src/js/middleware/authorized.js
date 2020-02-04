import { UPDATE_MEMBER_REQUEST } from '../constants/action-types'

export function authorizedMiddleware({dispatch}) {
    return function(next) {
        return function (action) {
            // TODO: check authorized
            if (action.type === UPDATE_MEMBER_REQUEST) {
                console.log("authorize")
            }
            return next(action)
        }
    }
}