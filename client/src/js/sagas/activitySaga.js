import { takeEvery, call, put } from 'redux-saga/effects'
import { FETCH_ACTIVITIES_REQUEST } from '../constants/action-types'
import { API_URL } from '../constants/config'
import * as Request from '../util/request'
import { fetchActivitiesSuccess, fetchActivitiesFailure } from '../actions/action'

export function* watchActivities() {
    yield takeEvery(FETCH_ACTIVITIES_REQUEST, fetchActivities)
}

function* fetchActivities() {
    try {
        const payload = yield call(getActivities)
        yield put(fetchActivitiesSuccess(payload))
    } catch(e) {
        yield put(fetchActivitiesFailure(e))
    }
}

function getActivities() {
    return Request.get(API_URL+"/activities")
}