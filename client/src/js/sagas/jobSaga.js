import { takeEvery, call, put } from 'redux-saga/effects'
import { FETCH_JOBS_REQUEST } from '../constants/action-types'
import { API_URL } from '../constants/config'
import * as Request from '../util/request'
import { fetchJobsSuccess, fetchJobsFailure } from '../actions/action'

// watch
export function* watchJobs() {
    yield takeEvery(FETCH_JOBS_REQUEST, fetchJobs)
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

function getJobs() {
    return Request.get(API_URL + "/jobs")
}
