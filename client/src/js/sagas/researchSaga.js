import { takeEvery, call, put } from 'redux-saga/effects'
import { FETCH_RESEARCHES_REQUEST } from '../constants/action-types'
import { BASE_URL } from '../constants/config'
import * as Request from '../util/request'
import { fetchResearchesFailure, fetchResearchesSuccess } from '../actions/action'

export function* watchResearches() {
    yield takeEvery(FETCH_RESEARCHES_REQUEST, fetchResearches)
}

function* fetchResearches() {
    try {
        const payload = yield call(getResearches)
        yield put(fetchResearchesSuccess(payload))
    } catch (e) {
        yield put(fetchResearchesFailure(e))
    }
}

function getResearches() {
    return Request.get(BASE_URL + "/researches")
}