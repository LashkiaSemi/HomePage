import { takeEvery, put, call } from 'redux-saga/effects'
import { FETCH_MEMBERS, LOADED_MEMBERS, API_ERROR } from '../constants/action-types'
import { get } from '../util/request'
import { BASE_URL } from '../constants/config'

export function* watchMembers() {
    yield takeEvery(FETCH_MEMBERS, fetchMembers)
}

function* fetchMembers() {
    try {
        const payload = yield call(getMembers)
        yield put({type: LOADED_MEMBERS, payload})
    } catch(e) {
        yield put({type: API_ERROR, payload: e})
    }
}

function getMembers() {
    return get(BASE_URL+"/users")
}