import { call, put, takeEvery } from 'redux-saga/effects'
import * as Request from '../util/request'
import { loginSuccess, loginFailure, logoutSuccess, logoutFailure } from '../actions/action'
import { LOGIN_REQUEST, LOGOUT_REQUEST } from '../constants/action-types'
import { BASE_URL } from '../constants/config'

export function* watchLogin() {
    yield takeEvery(LOGIN_REQUEST, handleLogin)
    yield takeEvery(LOGOUT_REQUEST, handleLogout)
}

function* handleLogin(action) {
    try {
        const payload = action.payload
        const response = yield call(loginRequest, payload.id, payload.password)
        yield put(loginSuccess(response))
    } catch(e) {
        yield put(loginFailure(e))
    }
}

function loginRequest(student_id, password) {
    return Request.post(BASE_URL+"/login", {student_id, password})
}

function* handleLogout() {
    try {
        yield call(logoutRequest)
        yield put(logoutSuccess())
    } catch(e) {
        yield put(logoutFailure(e))
    }
}

function logoutRequest() {
    return Request.httpDelete(BASE_URL+"/logout")
}