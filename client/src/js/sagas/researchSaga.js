import { takeEvery, call, put } from 'redux-saga/effects'
import { FETCH_RESEARCHES_REQUEST, CREATE_RESEARCH_REQUEST, UPDATE_RESEARCH_REQUEST, DELETE_RESEARCH_REQUEST } from '../constants/action-types'
import { API_URL } from '../constants/config'
import * as Request from '../util/request'
import { fetchResearchesFailure, fetchResearchesSuccess, createResearchSuccess, createResearchFailure, deleteResearchSuccess, deleteResearchRequest, updateResearchFailure, updateResearchSuccess, deleteResearchFailure } from '../actions/action'

// watcher
export function* watchResearches() {
    yield takeEvery(FETCH_RESEARCHES_REQUEST, fetchResearches)
    yield takeEvery(CREATE_RESEARCH_REQUEST, createResearch)
    yield takeEvery(UPDATE_RESEARCH_REQUEST, updateResearch)
    yield takeEvery(DELETE_RESEARCH_REQUEST, removeResearch)

}

// worker
function* fetchResearches() {
    try {
        const payload = yield call(getResearches)
        yield put(fetchResearchesSuccess(payload))
    } catch (e) {
        yield put(fetchResearchesFailure(e))
    }
}

function* createResearch(action) {
    try {
        const payload = yield call(postResearch, action.payload.body)
        yield put(createResearchSuccess(payload))
    } catch(e) {
        yield put(createResearchFailure(e))
    }
}

function* updateResearch(action) {
    try {
        const payload = yield call(putResearch, action.payload.id, action.payload.body)
        yield put(updateResearchSuccess(payload))
    } catch (e) {
        yield put(updateResearchFailure(e))
    }
}

function* removeResearch(action) {
    try {
        const payload = yield call(deleteResearch, action.payload.id)
        yield put(deleteResearchSuccess(payload))
    } catch (e) {
        yield put(deleteResearchFailure(e))
    }
}

// api call
function getResearches() {
    const options = { withCredentials: true }
    return Request.httpGet(API_URL + "/researches", options)
}

function postResearch(body) {
    var boundary = Math.random().toString().substr(2)
    const options = {
        withCredentials: true,
        headers: {
            "Content-Type": "multipart/form-data; boundary=------------------------" + boundary
        }
    }
    return Request.httpPost(API_URL+"/researches", body, options)
}

function putResearch(id, body) {
    var boundary = Math.random().toString().substr(2)
    const options = {
        withCredentials: true,
        headers: {
            "Content-Type": "multipart/form-data; boundary=------------------------" + boundary
        }
    }
    return Request.httpPut(API_URL+"/researches/"+id, body, options)
}

function deleteResearch(id) {
    const options = { withCredentials: true }
    return Request.httpDelete(API_URL+"/researches/"+id, options)
}