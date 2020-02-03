import { takeEvery, call, put } from 'redux-saga/effects'
import { FETCH_LECTURES_REQUEST } from '../constants/action-types'
import { BASE_URL } from '../constants/config'
import * as Request from '../util/request'
import { fetchLecturesSuccess, fetchLecturesFailure } from '../actions/action'

export function* watchLectures() {
    yield takeEvery(FETCH_LECTURES_REQUEST, fetchLectures)
}

function* fetchLectures() {
    try {
        const payload = yield call(getLectures)
        yield put(fetchLecturesSuccess(payload))
    } catch (e) {
        yield put(fetchLecturesFailure(e))
    }
}

function getLectures() {
    return Request.get(BASE_URL + "/lectures")
}