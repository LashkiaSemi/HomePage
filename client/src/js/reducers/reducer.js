import { ADD_USER } from '../constants/action-types'

const initialState = {
    sessions: [],
    users: [] // TODO: 後でけす。。。か？
}

function rootReducer(state = initialState, action) {
    switch(action.type) {
        case ADD_USER:
            return Object.assign({}, state, {
                users: state.users.concat(action.payload)
            })
        default:
            return state
    }
}

export default rootReducer