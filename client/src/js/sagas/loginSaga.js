import { call, put, takeEvery } from 'redux-saga/effects'
import * as Request from '../util/request'
import { loginSuccess, loginFailure } from '../actions/action'
import { LOGIN_REQUEST } from '../constants/action-types'
import { BASE_URL } from '../constants/config'

export function* watchLogin() {
    yield takeEvery(LOGIN_REQUEST, handleRequest)
}

function* handleRequest(action) {
    try {
        const payload = action.payload
        const response = yield call(loginRequest, payload.id, payload.password)
        console.log(response)
        yield put(loginSuccess(response))
    } catch(e) {
        yield put(loginFailure(e))
    }
}

function loginRequest(student_id, password) {
    return Request.post(BASE_URL+"/login", {student_id, password})
}
