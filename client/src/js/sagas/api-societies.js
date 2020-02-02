import { takeEvery, call, put } from 'redux-saga/effects'
import { FETCH_SOCIETIES, LOADED_SOCIETIES, API_ERROR } from '../constants/action-types'
import { BASE_URL } from '../constants/config'
import { get } from '../util/request'

export function* watchSocieties() {
    yield takeEvery(FETCH_SOCIETIES, fetchSocieties)
}

function* fetchSocieties() {
    try {
        const payload = yield call(getSocieties)
        yield put({ type: LOADED_SOCIETIES, payload })
    } catch (e) {
        yield put({ type: API_ERROR, payload: e })
    }
}

function getSocieties() {
    return get(BASE_URL + "/societies")
}