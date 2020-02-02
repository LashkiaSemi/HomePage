import { takeEvery, call, put } from 'redux-saga/effects'
import { FETCH_JOBS, LOADED_JOBS, API_ERROR } from '../constants/action-types'
import { BASE_URL } from '../constants/config'
import { get } from './request'

// watch
export function* watchJobs() {
    yield takeEvery(FETCH_JOBS, fetchJobs)
}

// work
function* fetchJobs() {
    try {
        const payload = yield call(getJobs) // TODO: callでおk?
        yield put({ type: LOADED_JOBS, payload })
    } catch (e) {
        yield put({ type: API_ERROR, payload: e })
    }
}

function getJobs() {
    return get(BASE_URL + "/jobs")
}
