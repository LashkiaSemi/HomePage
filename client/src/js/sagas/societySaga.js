import { takeEvery, call, put } from 'redux-saga/effects'
import { FETCH_SOCIETIES, LOADED_SOCIETIES, API_ERROR, FETCH_SOCIETIES_REQUEST, CREATE_SOCIETY_REQUEST, UPDATE_SOCIETY_REQUEST, DELETE_SOCIETY_REQUEST } from '../constants/action-types'
import { API_URL } from '../constants/config'
import * as Request from '../util/request'
import { fetchSocietiesSuccess, fetchSocietiesFailure, createSocietyRequest, createSocietySuccess, createSocietyFailure, updateSocietyFailure, updateSocietySuccess, deleteSocietyFailure, deleteSocietySuccess } from '../actions/action'

// watcher
export function* watchSocieties() {
    yield takeEvery(FETCH_SOCIETIES_REQUEST, fetchSocieties)
    yield takeEvery(CREATE_SOCIETY_REQUEST, createSociety)
    yield takeEvery(UPDATE_SOCIETY_REQUEST, updateSociety)
    yield takeEvery(DELETE_SOCIETY_REQUEST, removeSociety)
}

// worker
function* fetchSocieties() {
    try {
        const payload = yield call(getSocieties)
        yield put(fetchSocietiesSuccess(payload))
    } catch (e) {
        yield put(fetchSocietiesFailure(e))
    }
}

function* createSociety(action) {
    try {
        const payload = yield call(postSociety, action.payload.body)
        yield put(createSocietySuccess(payload))
    } catch(e) {
        yield put(createSocietyFailure(e))
    }
}

function* updateSociety(action) {
    try {
        const payload = yield call(putSociety, action.payload.id, action.payload.body)
        yield put(updateSocietySuccess(payload))
    } catch(e) {
        yield put(updateSocietyFailure(e))
    }
}

function* removeSociety(action) {
    try {
        const payload = yield call(deleteSociety, action.payload.id)
        yield put(deleteSocietySuccess(payload))
    } catch(e) {
        yield put(deleteSocietyFailure(e))
    }
}

// api call
function getSocieties() {
    const options = { withCredentials: true }
    return Request.httpGet(API_URL + "/societies", options)
}

function postSociety(body) {
    const options = { withCredentials: true }
    return Request.httpPost(API_URL+"/societies", body, options)
}

function putSociety(id, body) {
    const options = { withCredentials: true }
    return Request.httpPut(API_URL + "/societies/" + id, body, options)
}

function deleteSociety(id) {
    const options = {withCredentials: true}
    return Request.httpDelete(API_URL+"/societies/"+id, options)
}