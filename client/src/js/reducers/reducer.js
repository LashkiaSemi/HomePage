import { 
    FETCH_ACTIVITIES_SUCCESS, FETCH_ACTIVITIES_FAILURE,
    FETCH_EQUIPMENTS_SUCCESS, FETCH_EQUIPMENTS_FAILURE,
    FETCH_JOBS_SUCCESS, FETCH_JOBS_FAILURE,
    FETCH_LECTURES_SUCCESS, FETCH_LECTURES_FAILURE,
    FETCH_MEMBERS_SUCCESS, FETCH_MEMBERS_FAILURE,
    FETCH_MEMBER_SUCCESS, FETCH_MEMBER_FAILURE,
    FETCH_RESEARCHES_SUCCESS, FETCH_RESEARCHES_FAILURE,
    FETCH_SOCIETIES_SUCCESS, FETCH_SOCIETIES_FAILURE,
    SHOW_LOADING, HIDE_LOADING, LOGIN_SUCCESS, LOGIN_FAILURE, UPDATE_MEMBER_REQUEST, UPDATE_MEMBER_SUCCESS, FETCH_ACCOUNT_SUCCESS, UPDATE_ACCOUNT_SUCCESS, FETCH_ACCOUNT_FAILURE, LOGOUT_SUCCESS, LOGOUT_FAILURE } from '../constants/action-types'
import { combineReducers } from 'redux'

import { STRAGE_KET } from '../constants/config'

// TODO: reducerの分割した方がよくね？

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
function logged(state=[], action) {
    switch(action.type) {
        case LOGIN_SUCCESS:
            //TODO: 納得いかないw
            // localstrageにidをタンク
            localStorage.setItem(STRAGE_KET, action.payload.data.user_id)
            window.location.href = "http://localhost:3000/"
            return Object.assign([], state, state.concat({user_id:action.payload.data.user_id}))
        
        case LOGIN_FAILURE:
            console.log("reducer: login failure")
            return state
        
        case LOGOUT_SUCCESS:
            console.log("reducer: logout success")
            // localstarageを消せ！
            localStorage.removeItem(STRAGE_KET)
            // TODO: redirect
            window.location.href = "http://localhost:3000"
            return state
        
        case LOGOUT_FAILURE:
            console.log("reducer: logout failure")
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
        case UPDATE_MEMBER_SUCCESS:
            console.log("reducer: update success")
            return Object.assign({}, action.payload.data)
        default:
            return state
    }
}

// 複数人で同時にやったらstateが起き変わったりしない？
function account(state={}, action) {
    switch(action.type) {
        case FETCH_ACCOUNT_SUCCESS:
            return Object.assign({}, action.payload.data)
        case FETCH_ACCOUNT_FAILURE:
            // TODO: redirect
            return state
        case UPDATE_ACCOUNT_SUCCESS:
            console.log("reducer: update success")
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
    logged,
    account,
})

export default rootReducer