import { LOADED_JOBS, LOADED_MEMBERS, SHOW_LOADING, HIDE_LOADING } from '../constants/action-types'
import { combineReducers } from 'redux'

function isLoading(state=false, action) {
    switch (action.type){
        case SHOW_LOADING:
            return true
        case HIDE_LOADING:
            return false
        default:
            return state
    }
}

function members(state=[], action) {
    switch(action.type) {
        case LOADED_MEMBERS:
            return Object.assign([], action.payload.data.users)
        default:
            return state
    }
}

function jobs(state=[], action) {
    switch(action.type) {
        case LOADED_JOBS:
            return Object.assign([], action.payload.data.jobs)
        default:
            return state
    }
}

const rootReducer = combineReducers({
    isLoading,
    members,
    jobs,
})

export default rootReducer