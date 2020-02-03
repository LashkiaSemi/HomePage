import { 
    LOGIN_REQUEST, LOGIN_SUCCESS, LOGIN_FAILURE,
    FETCH_JOBS, FETCH_MEMBERS, FETCH_MEMBER,
    FETCH_ACTIVITIES, FETCH_SOCIETIES,
    FETCH_RESEARCHES, FETCH_EQUIPMENTS,
    FETCH_LECTURES,
    SHOW_LOADING, HIDE_LOADING, API_ERROR, UPDATE_MEMBER_REQUEST, UPDATE_MEMBER_SUCCESS, UPDATE_MEMBER_FAILURE,  } 
    from '../constants/action-types'

export const loginRequest = (payload) => { return {type: LOGIN_REQUEST, payload} }
export const loginSuccess = (payload) => { return {type: LOGIN_SUCCESS, payload} }
export const loginFailure = (payload) => { return {type: LOGIN_FAILURE, payload} }

export const updateMemberRequest = (payload) => { return {type: UPDATE_MEMBER_REQUEST, payload} }
export const updateMemberSuccess = (payload) => { return {type: UPDATE_MEMBER_SUCCESS, payload} }
export const updateMemberFailure = (payload) => { return {type: UPDATE_MEMBER_FAILURE, payload} }

export const fetchMembers = () => { return {type: FETCH_MEMBERS} }
export const fetchMember = (payload) => { return {type: FETCH_MEMBER, payload} }
export const fetchActivities = () => {return {type: FETCH_ACTIVITIES}}
export const fetchSocieties = () => { return {type: FETCH_SOCIETIES} }
export const fetchResearches = () => { return {type: FETCH_RESEARCHES} }
export const fetchEquipments = () => { return {type: FETCH_EQUIPMENTS} }
export const fetchLectures = () => { return {type: FETCH_LECTURES} }

export function fetchJobs() {
    return {
        type: FETCH_JOBS,
    }
}

export function showLoading() {
    return {
        type: SHOW_LOADING
    }
}

export function hideLoading(){
    return {
        type: HIDE_LOADING
    }
}

export const apiError = (payload) => { return {type: API_ERROR, payload} }