import { 
    LOADED_JOBS, LOADED_MEMBERS,
    LOADED_ACTIVITIES, LOADED_SOCIETIES,
    LOADED_RESEARCHES, LOADED_EQUIPMENTS,
    LOADED_LECTURES,
    SHOW_LOADING, HIDE_LOADING } from '../constants/action-types'
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

function activities(state=[], action) {
    switch(action.type) {
        case LOADED_ACTIVITIES:
            return Object.assign([], action.payload.data.activities)
        default:
            return state
    }
}

function societies(state=[], action) {
    switch(action.type) {
        case LOADED_SOCIETIES:
            return Object.assign([], action.payload.data.societies)
        default:
            return state
    }
}

function researches(state=[], action) {
    switch(action.type) {
        case LOADED_RESEARCHES:
            return Object.assign([], action.payload.data.researches)
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

function equipments(state=[], action) {
    switch(action.type) {
        case LOADED_EQUIPMENTS:
            return Object.assign([], action.payload.data.equipments)
        default:
            return state
    }
}

function lectures(state=[], action) {
    switch(action.type) {
        case LOADED_LECTURES:
            return Object.assign([], action.payload.data.lectures)
        default:
            return state
    }
}

const rootReducer = combineReducers({
    isLoading,
    activities,
    societies,
    researches,
    members,
    jobs,
    equipments,
    lectures,
})

export default rootReducer