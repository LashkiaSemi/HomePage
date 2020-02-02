import { FETCH_JOBS, FETCH_MEMBERS, SHOW_LOADING, HIDE_LOADING } from '../constants/action-types'

export const fetchMembers = () => { return {type: FETCH_MEMBERS} }

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