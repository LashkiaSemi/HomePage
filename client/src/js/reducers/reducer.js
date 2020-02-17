import { 
    FETCH_ACTIVITIES_SUCCESS, FETCH_ACTIVITIES_FAILURE,
    FETCH_EQUIPMENTS_SUCCESS, FETCH_EQUIPMENTS_FAILURE,
    FETCH_JOBS_SUCCESS, FETCH_JOBS_FAILURE,
    FETCH_LECTURES_SUCCESS, FETCH_LECTURES_FAILURE,
    FETCH_MEMBERS_SUCCESS, FETCH_MEMBERS_FAILURE,
    FETCH_MEMBER_SUCCESS, FETCH_MEMBER_FAILURE,
    FETCH_RESEARCHES_SUCCESS, FETCH_RESEARCHES_FAILURE,
    FETCH_SOCIETIES_SUCCESS, FETCH_SOCIETIES_FAILURE,
    SHOW_LOADING, HIDE_LOADING, LOGIN_SUCCESS, LOGIN_FAILURE, UPDATE_MEMBER_REQUEST, UPDATE_MEMBER_SUCCESS, FETCH_ACCOUNT_SUCCESS, UPDATE_ACCOUNT_SUCCESS, FETCH_ACCOUNT_FAILURE, LOGOUT_SUCCESS, LOGOUT_FAILURE, UPDATE_ACCOUNT_PASSWORD_SUCCESS, UPDATE_ACCOUNT_PASSWORD_FAILURE, CREATE_LECTURE_FAILURE, CREATE_LECTURE_SUCCESS, FETCH_LECTURE_SUCCESS, CREATE_ACTIVITY_SUCCESS, UPDATE_ACTIVITY_SUCCESS, DELETE_ACTIVITY_SUCCESS, UPDATE_LECTURE_SUCCESS, DELETE_LECTURE_SUCCESS, CREATE_SOCIETY_REQUEST, UPDATE_SOCIETY_FAILURE, UPDATE_SOCIETY_SUCCESS, CREATE_SOCIETY_SUCCESS, DELETE_SOCIETY_SUCCESS, CREATE_RESEARCH_SUCCESS, UPDATE_RESEARCH_FAILURE, UPDATE_RESEARCH_SUCCESS, DELETE_RESEARCH_REQUEST,
    DELETE_RESEARCH_SUCCESS, 
    CREATE_MEMBER_SUCCESS,
    DELETE_MEMBER_SUCCESS,
    CREATE_JOB_SUCCESS,
    UPDATE_JOB_SUCCESS,
    DELETE_JOB_SUCCESS,
    CREATE_EQUIPMENT_REQUEST,
    CREATE_EQUIPMENT_SUCCESS,
    UPDATE_EQUIPMENT_SUCCESS,
    DELETE_EQUIPMENT_REQUEST,
    DELETE_EQUIPMENT_SUCCESS,
    FETCH_TAGS_REQUEST,
    FETCH_TAGS_SUCCESS,
    CREATE_TAG_SUCCESS,
    UPDATE_TAG_SUCCESS,
    DELETE_TAG_SUCCESS} from '../constants/action-types'
import { combineReducers } from 'redux'

import { CLIENT_URL, STRAGE_KEY } from '../constants/config'

import * as Crypto from '../util/crypto'

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

function logged(state=[], action) {
    switch(action.type) {
        case LOGIN_SUCCESS:
            // localstrageに入れる値の暗号化
            const strageValue = Crypto.Encrypt(action.payload.data.user_id + action.payload.data.role)
            // localstrageにタンク
            localStorage.setItem(STRAGE_KEY, strageValue)
            window.location.href = CLIENT_URL
            return Object.assign([], state, state.concat({user_id:action.payload.data.user_id}))
        case LOGIN_FAILURE:
            console.log("reducer: login failure")
            return state
        
        case LOGOUT_SUCCESS:
            console.log("reducer: logout success")
            // localstarageを消せ！
            localStorage.removeItem(STRAGE_KEY)
            // memo: redirect
            window.location.href = CLIENT_URL
            return state
        
        case LOGOUT_FAILURE:
            console.log("reducer: logout failure")
            // TODO: これいる...?
            if (localStorage.getItem(STRAGE_KEY)){
                localStorage.removeItem(STRAGE_KEY)
            }
            window.location.href = CLIENT_URL
            return state
        default:
            return state
    }
}

function activities(state=[], action) {
    switch(action.type) {
        case FETCH_ACTIVITIES_SUCCESS:
            return Object.assign([], action.payload.data.activities)
        case CREATE_ACTIVITY_SUCCESS:
            console.log("reducer: create activiry success")
            return state
        case UPDATE_ACTIVITY_SUCCESS:
            console.log("reducer: update activity success")
            return state
        case DELETE_ACTIVITY_SUCCESS:
            console.log("reducer: delete activity success")
            return state
        default:
            return state
    }
}

function societies(state=[], action) {
    switch(action.type) {
        case FETCH_SOCIETIES_SUCCESS:
            return Object.assign([], action.payload.data.societies)
        case CREATE_SOCIETY_SUCCESS:
            console.log("reducer: create society success")
            return state
        case UPDATE_SOCIETY_SUCCESS:
            console.log("reducer: update society success")
            return state
        case DELETE_SOCIETY_SUCCESS:
            console.log("reducer: create society success")
            return state
        default:
            return state
    }
}

function researches(state=[], action) {
    switch(action.type) {
        case FETCH_RESEARCHES_SUCCESS:
            return Object.assign([], action.payload.data.researches)
        case CREATE_RESEARCH_SUCCESS:
            console.log("reducer: create research success")
            return state
        case UPDATE_RESEARCH_SUCCESS:
            console.log("reducer: update research success")
            return state
        case DELETE_RESEARCH_SUCCESS:
            console.log("reducer: delete research success")
            return state
        default:
            return state
    }
}

function members(state=[], action) {
    switch(action.type) {
        case FETCH_MEMBERS_SUCCESS:
            return Object.assign([], action.payload.data.users)
        case CREATE_MEMBER_SUCCESS:
            console.log("reducer: create member success")
            return state
        case UPDATE_MEMBER_SUCCESS:
            console.log("reducer: update member success")
            return state
        case DELETE_MEMBER_SUCCESS:
            console.log("reducer: delete member success")
            return state
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
        case UPDATE_ACCOUNT_PASSWORD_SUCCESS:
            console.log("reducer: password update success")
            return state
        case UPDATE_ACCOUNT_PASSWORD_FAILURE:
            console.log("reducer: password update fail")
            return state
        default:
            return state
    }
}

function jobs(state=[], action) {
    switch(action.type) {
        case FETCH_JOBS_SUCCESS:
            return Object.assign([], action.payload.data.jobs)
        case CREATE_JOB_SUCCESS:
            console.log("reducer: create jobs success")
            return state
        case UPDATE_JOB_SUCCESS:
            console.log("reducer: update jobs success")
            return state
        case DELETE_JOB_SUCCESS:
            console.log("reducer: delete jobs success")
            return state
        default:
            return state
    }
}

function equipments(state=[], action) {
    switch(action.type) {
        case FETCH_EQUIPMENTS_SUCCESS:
            return Object.assign([], action.payload.data.equipments)
        case CREATE_EQUIPMENT_SUCCESS:
            console.log("reducer: create equipment success")
            return state
        case UPDATE_EQUIPMENT_SUCCESS:
            console.log("reducer: update equipment success")
            return state
        case DELETE_EQUIPMENT_SUCCESS:
            console.log("reducer: delete equipment success")
            return state
        default:
            return state
    }
}

function lectures(state=[], action) {
    switch(action.type) {
        case FETCH_LECTURES_SUCCESS:
            return Object.assign([], action.payload.data.lectures)
        // case FETCH_LECTURE_SUCCESS:
        //     return Object.assign([], action.payload.data)
        case CREATE_LECTURE_SUCCESS:
            console.log("reducer: lecture create success")
            // TODO; redirect
            return state
        case UPDATE_LECTURE_SUCCESS:
            // TODO: redirect
            console.log("reducer: lecture update success")
            return state
        case DELETE_LECTURE_SUCCESS:
            // TODO: reload
            console.log("reducer: lecture delete success")
            return state
        default:
            return state
    }
}

function lecture(state={}, action) {
    switch (action.type) {

        case FETCH_LECTURE_SUCCESS:
            return Object.assign({}, action.payload.data)

        default:
            return state
    }
}

function tags(state = [], action) {
    switch(action.type) {
        case FETCH_TAGS_SUCCESS:
            return Object.assign([], action.payload.data.tags)
        case CREATE_TAG_SUCCESS:
            console.log("reducer: create tag success")
            return state
        case UPDATE_TAG_SUCCESS:
            console.log("reducer: update tag success")
            return state
        case DELETE_TAG_SUCCESS:
            console.log("reducer: delete tag success")
            return state
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
    lecture,
    tags,
    logged,
    account,
})

export default rootReducer