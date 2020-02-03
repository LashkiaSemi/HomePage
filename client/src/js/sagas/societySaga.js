import { takeEvery, call, put } from 'redux-saga/effects'
import { FETCH_SOCIETIES, LOADED_SOCIETIES, API_ERROR, FETCH_SOCIETIES_REQUEST } from '../constants/action-types'
import { BASE_URL } from '../constants/config'
import * as Request from '../util/request'
import { fetchSocietiesSuccess, fetchSocietiesFailure } from '../actions/action'

export function* watchSocieties() {
    yield takeEvery(FETCH_SOCIETIES_REQUEST, fetchSocieties)
}

function* fetchSocieties() {
    try {
        const payload = yield call(getSocieties)
        yield put(fetchSocietiesSuccess(payload))
    } catch (e) {
        yield put(fetchSocietiesFailure(e))
    }
}

function getSocieties() {
    return Request.get(BASE_URL + "/societies")
}