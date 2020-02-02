import { FETCH_JOBS, SHOW_LOADING, HIDE_LOADING } from '../constants/action-types'

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