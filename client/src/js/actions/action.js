import { 
    FETCH_JOBS, FETCH_MEMBERS,
    FETCH_ACTIVITIES, FETCH_SOCIETIES,
    FETCH_RESEARCHES, FETCH_EQUIPMENTS,
    FETCH_LECTURES,
    SHOW_LOADING, HIDE_LOADING } 
    from '../constants/action-types'

export const fetchMembers = () => { return {type: FETCH_MEMBERS} }

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