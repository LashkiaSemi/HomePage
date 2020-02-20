import { takeEvery, call, put } from 'redux-saga/effects'
import { FETCH_LECTURES_REQUEST, CREATE_LECTURE_REQUEST, FETCH_LECTURE_REQUEST, UPDATE_LECTURE_REQUEST, DELETE_LECTURE_REQUEST, ADMIN_CREATE_LECTURE_REQUEST, ADMIN_UPDATE_LECTURE_REQUEST, ADMIN_DELETE_LECTURE_REQUEST } from '../constants/action-types'
import { API_URL } from '../constants/config'
import * as Request from '../util/request'
import { fetchLecturesSuccess, 
         fetchLecturesFailure, 
         createLectureSuccess, 
         createLectureFailure, 
         fetchLectureFailure, 
         fetchLectureSuccess, 
         updateLectureSuccess, 
         updateLectureFailure, 
         deleteLectureSuccess, 
         deleteLectureFailure,
         adminCreateLectureSuccess,
         adminCreateLectureFailure,
         adminUpdateLectureSuccess,
         adminUpdateLectureFailure,
         adminDeleteLectureSuccess,
         adminDeleteLectureFailure } from '../actions/action'

// watcher
export function* watchLectures() {
    yield takeEvery(FETCH_LECTURES_REQUEST, fetchLectures)
    yield takeEvery(FETCH_LECTURE_REQUEST, fetchLecture)
    yield takeEvery(CREATE_LECTURE_REQUEST, createLecture)
    yield takeEvery(UPDATE_LECTURE_REQUEST, updateLecture)
    yield takeEvery(DELETE_LECTURE_REQUEST, removeLecture)

    // admin
    yield takeEvery(ADMIN_CREATE_LECTURE_REQUEST, adminCreateLecture)
    yield takeEvery(ADMIN_UPDATE_LECTURE_REQUEST, adminUpdateLecture)
    yield takeEvery(ADMIN_DELETE_LECTURE_REQUEST, adminRemoveLecture)
}

// worker
function* fetchLectures() {
    try {
        const payload = yield call(getLectures)
        yield put(fetchLecturesSuccess(payload))
    } catch (e) {
        yield put(fetchLecturesFailure(e))
    }
}

function* fetchLecture(action) {
    try {
        const payload = yield call(getLecture, action.payload.id)
        yield put(fetchLectureSuccess(payload))
    } catch(e) {
        yield put(fetchLectureFailure(e))
    }
}

function* createLecture(action) {
    try {
        const payload = yield call(postLecture, action.payload.body)
        yield put(createLectureSuccess(payload))
    } catch(e) {
        yield put(createLectureFailure(e))
    }
}

function* updateLecture(action) {
    try {
        const payload = yield call(putLecture, action.payload.id, action.payload.body)
        yield put(updateLectureSuccess(payload))
    }catch(e) {
        yield put(updateLectureFailure(e))
    }
}

function* removeLecture(action) {
    try {
        const payload = yield call(deleteLecture, action.payload.id)
        yield put(deleteLectureSuccess(payload))
    } catch(e) {
        yield put(deleteLectureFailure(e))
    }
}


function* adminCreateLecture(action) {
    try {
        const payload = yield call(postLecture, action.payload.body)
        yield put(adminCreateLectureSuccess(payload))
    } catch (e) {
        yield put(adminCreateLectureFailure(e))
    }
}

function* adminUpdateLecture(action) {
    try {
        const payload = yield call(putLecture, action.payload.id, action.payload.body)
        yield put(adminUpdateLectureSuccess(payload))
    } catch (e) {
        yield put(adminUpdateLectureFailure(e))
    }
}

function* adminRemoveLecture(action) {
    try {
        const payload = yield call(deleteLecture, action.payload.id)
        yield put(adminDeleteLectureSuccess(payload))
    } catch (e) {
        yield put(adminDeleteLectureFailure(e))
    }
}

// api call
function getLectures() {
    const options = { withCredentials: true }
    return Request.httpGet(API_URL + "/lectures", options)
}

function getLecture(id) {
    const options = { withCredentials: true }
    return Request.httpGet(API_URL+"/lectures/"+id, options)
}

function postLecture(body) {
    var boundary = Math.random().toString().substr(2)
    const options = { 
        withCredentials: true,
        headers: {
            "Content-Type": "multipart/form-data; boundary=------------------------" + boundary
        }
    }
    return Request.httpPost(API_URL+"/lectures", body, options)
}

function putLecture(id, body) {
    var boundary = Math.random().toString().substr(2)
    const options = {
        withCredentials: true,
        headers: {
            "Content-Type": "multipart/form-data; boundary=------------------------" + boundary
        }
    }
    return Request.httpPut(API_URL+"/lectures/"+id, body, options)
}

function deleteLecture(id) {
    const options = { withCredentials: true }
    return Request.httpDelete(API_URL+"/lectures/"+id, options)
}