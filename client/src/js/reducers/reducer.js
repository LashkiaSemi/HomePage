import { LOADED_JOBS, SHOW_LOADING, HIDE_LOADING } from '../constants/action-types'
import { combineReducers } from 'redux'

function isLoaded(state=false, action) {
    switch (action.type){
        case SHOW_LOADING:
            return true
        case HIDE_LOADING:
            return false
        default:
            return state
    }
}

function jobs(state=[], action) {
    switch(action.type) {
        case LOADED_JOBS:
            return state.concat(action.payload.data.jobs)
        default:
            return state
    }
}

const rootReducer = combineReducers({
    isLoaded,
    jobs,
})

export default rootReducer