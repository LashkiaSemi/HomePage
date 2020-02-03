import { 
    FETCH_ACTIVITIES_SUCCESS, FETCH_ACTIVITIES_FAILURE,
    FETCH_EQUIPMENTS_SUCCESS, FETCH_EQUIPMENTS_FAILURE,
    FETCH_JOBS_SUCCESS, FETCH_JOBS_FAILURE,
    FETCH_LECTURES_SUCCESS, FETCH_LECTURES_FAILURE,
    FETCH_MEMBERS_SUCCESS, FETCH_MEMBERS_FAILURE,
    FETCH_MEMBER_SUCCESS, FETCH_MEMBER_FAILURE,
    FETCH_RESEARCHES_SUCCESS, FETCH_RESEARCHES_FAILURE,
    FETCH_SOCIETIES_SUCCESS, FETCH_SOCIETIES_FAILURE,
    SHOW_LOADING, HIDE_LOADING, LOGIN_SUCCESS, LOGIN_FAILURE, UPDATE_MEMBER_REQUEST } from '../constants/action-types'
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

// TODO: なんとかしろ
function login(state=[], action) {
    switch(action.type) {
        case LOGIN_SUCCESS:
            // sessionを...
            // TODO: 同じsessionがupdateされない。というかこの辺の管理どうしろと？
            return Object.assign([], state, state.concat({user_id:action.payload.data.user_id, }))
        case LOGIN_FAILURE:
            console.log("reducer: login failure")
            return state
        default:
            return state
    }
}

function activities(state=[], action) {
    switch(action.type) {
        case FETCH_ACTIVITIES_SUCCESS:
            return Object.assign([], action.payload.data.activities)
        default:
            return state
    }
}

function societies(state=[], action) {
    switch(action.type) {
        case FETCH_SOCIETIES_SUCCESS:
            return Object.assign([], action.payload.data.societies)
        default:
            return state
    }
}

function researches(state=[], action) {
    switch(action.type) {
        case FETCH_RESEARCHES_SUCCESS:
            return Object.assign([], action.payload.data.researches)
        default:
            return state
    }
}

function members(state=[], action) {
    switch(action.type) {
        case FETCH_MEMBERS_SUCCESS:
            return Object.assign([], action.payload.data.users)
        default:
            return state
    }
}

// membersに統合してしまってもいいかもしらん
function member(state={}, action) {
    switch(action.type) {
        case FETCH_MEMBER_SUCCESS:
            return Object.assign({}, action.payload.data)
        case UPDATE_MEMBER_REQUEST:
            return Object.assign({}, action.payload.data)
        default:
            return state
    }
}

function jobs(state=[], action) {
    switch(action.type) {
        case FETCH_JOBS_SUCCESS:
            return Object.assign([], action.payload.data.jobs)
        default:
            return state
    }
}

function equipments(state=[], action) {
    switch(action.type) {
        case FETCH_EQUIPMENTS_SUCCESS:
            return Object.assign([], action.payload.data.equipments)
        default:
            return state
    }
}

function lectures(state=[], action) {
    switch(action.type) {
        case FETCH_LECTURES_SUCCESS:
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
    member,
    jobs,
    equipments,
    lectures,
    login,
})

export default rootReducer