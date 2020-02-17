import { takeEvery, call, put } from 'redux-saga/effects'
import { FETCH_JOBS_REQUEST, CREATE_JOB_REQUEST, UPDATE_JOB_REQUEST, DELETE_JOB_REQUEST } from '../constants/action-types'
import { API_URL } from '../constants/config'
import * as Request from '../util/request'
import { fetchJobsSuccess, fetchJobsFailure, createJobSuccess, createJobFailure, updateJobSuccess, updateJobFailure, deleteJobSuccess, deleteJobFailure } from '../actions/action'

// watch
export function* watchJobs() {
    yield takeEvery(FETCH_JOBS_REQUEST, fetchJobs)
    yield takeEvery(CREATE_JOB_REQUEST, createJob)
    yield takeEvery(UPDATE_JOB_REQUEST, updateJob)
    yield takeEvery(DELETE_JOB_REQUEST, removeJob)
}

// work
function* fetchJobs() {
    try {
        const payload = yield call(getJobs)
        yield put(fetchJobsSuccess(payload))
    } catch (e) {
        yield put(fetchJobsFailure(e))
    }
}

function* createJob(action) {
    try {
        const payload = yield call(postJob, action.payload.body)
        yield put(createJobSuccess(payload))
    } catch(e) {
        yield put(createJobFailure(e))
    }
}

function* updateJob(action) {
    try {
        const payload = yield call(putJob, action.payload.id, action.payload.body)
        yield put(updateJobSuccess(payload))
    } catch (e) {
        yield put(updateJobFailure(e))
    }
}

function* removeJob(action) {
    try {
        const payload = yield call(deleteJob, action.payload.id)
        yield put(deleteJobSuccess(payload))
    } catch (e) {
        yield put(deleteJobFailure(e))
    }
}

// api call
function getJobs() {
    const options = { withCredentials: true }
    return Request.httpGet(API_URL + "/jobs", options)
}

function postJob(body) {
    const options = { withCredentials: true }
    return Request.httpPost(API_URL + "/jobs", body, options)
}

function putJob(id, body) {
    const options = { withCredentials: true }
    return Request.httpPut(API_URL + "/jobs/"+id, body, options)
}

function deleteJob(id) {
    const options = { withCredentials: true }
    return Request.httpDelete(API_URL + "/jobs/"+id, options)
}
