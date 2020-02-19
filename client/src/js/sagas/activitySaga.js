import { takeEvery, call, put } from 'redux-saga/effects'
import { FETCH_ACTIVITIES_REQUEST, CREATE_ACTIVITY_REQUEST, UPDATE_ACTIVITY_REQUEST, DELETE_ACTIVITY_REQUEST } from '../constants/action-types'
import { API_URL } from '../constants/config'
import * as Request from '../util/request'
import { fetchActivitiesSuccess, fetchActivitiesFailure, createActivitySuccess, createActivityFailure, updateActivitySuccess, updateActivityFailure, deleteActivitySuccess, deleteActivityFailure } from '../actions/action'

// watcher
export function* watchActivities() {
    yield takeEvery(FETCH_ACTIVITIES_REQUEST, fetchActivities)
    yield takeEvery(CREATE_ACTIVITY_REQUEST, createActivity)
    yield takeEvery(UPDATE_ACTIVITY_REQUEST, updateActivity)
    yield takeEvery(DELETE_ACTIVITY_REQUEST, removeActivity)
}

// worker
function* fetchActivities() {
    try {
        const payload = yield call(getActivities)
        yield put(fetchActivitiesSuccess(payload))
    } catch(e) {
        yield put(fetchActivitiesFailure(e))
    }
}

function* createActivity(action) {
    try {
        const payload = yield call(postActivity, action.payload.body)
        yield put(createActivitySuccess(payload))
    } catch(e) {
        yield put(createActivityFailure(e))
    }
}

function* updateActivity(action) {
    try {
        const payload = yield call(putActivity, action.payload.id, action.payload.body)
        yield put(updateActivitySuccess(payload))
    } catch(e) {
        console.log(e)
        yield put(updateActivityFailure(e))
    }
}

function* removeActivity(action) {
    try {
        const payload = yield call(deleteActivity, action.payload.id)
        yield put(deleteActivitySuccess(payload))
    } catch(e) {
        yield put(deleteActivityFailure(e))
    }
}

// api call
function getActivities() {
    const options = { withCredentials: true }
    return Request.httpGet(API_URL+"/activities", options)
}

function postActivity(body) {
    const options = { withCredentials: true }
    return Request.httpPost(API_URL+"/activities", body, options)
}

function putActivity(id, body) {
    const options = { withCredentials: true }
    return Request.httpPut(API_URL+"/activities/"+id, body, options)
}

function deleteActivity(id) {
    const options = { withCredentials: true }
    return Request.httpDelete(API_URL+"/activities/"+id, options)
}