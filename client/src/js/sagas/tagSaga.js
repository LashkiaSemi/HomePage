import { takeEvery, call, put } from 'redux-saga/effects'
import { FETCH_TAGS_REQUEST, CREATE_TAG_REQUEST, UPDATE_TAG_REQUEST, DELETE_TAG_REQUEST } from '../constants/action-types'
import { API_URL } from '../constants/config'
import * as Request from '../util/request'
import { fetchTagsSuccess, fetchTagsFailure, createTagSuccess, createTagFailure, updateTagSuccess, updateTagFailure, deleteTagSuccess, deleteTagFailure } from '../actions/action'

// watcher
export function* watchTags() {
    yield takeEvery(FETCH_TAGS_REQUEST, fetchTags)
    yield takeEvery(CREATE_TAG_REQUEST, createTag)
    yield takeEvery(UPDATE_TAG_REQUEST, updateTag)
    yield takeEvery(DELETE_TAG_REQUEST, removeTag)
}

// worker
function* fetchTags() {
    try {
        const payload = yield call(getTags)
        yield put(fetchTagsSuccess(payload))
    } catch (e) {
        yield put(fetchTagsFailure(e))
    }
}

function* createTag(action) {
    try {
        const payload = yield call(postTag, action.payload.body)
        yield put(createTagSuccess(payload))
    } catch(e) {
        yield put(createTagFailure(e))
    }
}

function* updateTag(action) {
    try {
        const payload = yield call(putTag, action.payload.id, action.payload.body)
        yield put(updateTagSuccess(payload))
    } catch(e) {
        yield put(updateTagFailure(e))
    }
}

function* removeTag(action) {
    try {
        const payload = yield call(deleteTag, action.payload.id)
        yield put(deleteTagSuccess(payload))
    } catch(e) {
        yield put(deleteTagFailure(e))
    }
}

// api call
function getTags() {
    const options = { withCredentials: true }
    return Request.httpGet(API_URL + "/tags", options)
}

function postTag(body) {
    const options = { withCredentials: true }
    return Request.httpPost(API_URL + "/tags", body, options)
}

function putTag(id, body) {
    const options = { withCredentials: true }
    return Request.httpPut(API_URL + "/tags/" + id, body, options)
}

function deleteTag(id) {
    const options = { withCredentials: true }
    return Request.httpDelete(API_URL + "/tags/" + id, options)
}