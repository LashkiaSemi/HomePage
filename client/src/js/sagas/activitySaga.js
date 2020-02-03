import { takeEvery, call, put } from 'redux-saga/effects'
import { FETCH_ACTIVITIES, LOADED_ACTIVITIES, API_ERROR } from '../constants/action-types'
import { BASE_URL } from '../constants/config'
import { get } from '../util/request'

export function* watchActivities() {
    yield takeEvery(FETCH_ACTIVITIES, fetchActivities)
}

function* fetchActivities() {
    try {
        const payload = yield call(getActivities)
        yield put({type: LOADED_ACTIVITIES, payload})
    } catch(e) {
        yield put({type: API_ERROR, payload: e})
    }
}

function getActivities() {
    return get(BASE_URL+"/activities")
}