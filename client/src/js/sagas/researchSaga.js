import { takeEvery, call, put } from 'redux-saga/effects'
import { FETCH_RESEARCHES, LOADED_RESEARCHES, API_ERROR } from '../constants/action-types'
import { BASE_URL } from '../constants/config'
import { get } from '../util/request'

export function* watchResearches() {
    yield takeEvery(FETCH_RESEARCHES, fetchResearches)
}

function* fetchResearches() {
    try {
        const payload = yield call(getResearches)
        yield put({ type: LOADED_RESEARCHES, payload })
    } catch (e) {
        yield put({ type: API_ERROR, payload: e })
    }
}

function getResearches() {
    return get(BASE_URL + "/researches")
}