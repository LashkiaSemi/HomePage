import { takeEvery, call, put } from 'redux-saga/effects'
import { FETCH_LECTURES_REQUEST, CREATE_LECTURE_REQUEST, FETCH_LECTURE_REQUEST } from '../constants/action-types'
import { API_URL } from '../constants/config'
import * as Request from '../util/request'
import { fetchLecturesSuccess, fetchLecturesFailure, createLectureSuccess, createLectureFailure, fetchLectureFailure, fetchLectureSuccess } from '../actions/action'

export function* watchLectures() {
    yield takeEvery(FETCH_LECTURES_REQUEST, fetchLectures)
    yield takeEvery(FETCH_LECTURE_REQUEST, fetchLecture)
    yield takeEvery(CREATE_LECTURE_REQUEST, createLecture)
}

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

function getLectures() {
    return Request.get(API_URL + "/lectures")
}

function getLecture(id) {
    return Request.get(API_URL+"/lectures/"+id)
}

function postLecture(body) {
    var boundary = Math.random().toString().substr(2)
    const options = { 
        withCredentials: true,
        headers: {
            "Content-Type": "multipart/form-data; boundary=------------------------" + boundary
        }
    }
    return Request.post(API_URL+"/lectures", body, options)
}