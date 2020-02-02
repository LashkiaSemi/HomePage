import { takeEvery, call, put } from 'redux-saga/effects'
import { FETCH_LECTURES, LOADED_LECTURES, API_ERROR } from '../constants/action-types'
import { BASE_URL } from '../constants/config'
import { get } from '../util/request'

export function* watchLectures() {
    yield takeEvery(FETCH_LECTURES, fetchLectures)
}

function* fetchLectures() {
    try {
        const payload = yield call(getLectures)
        yield put({ type: LOADED_LECTURES, payload })
    } catch (e) {
        yield put({ type: API_ERROR, payload: e })
    }
}

function getLectures() {
    return get(BASE_URL + "/lectures")
}