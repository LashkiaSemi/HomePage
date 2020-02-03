import { takeEvery, put, call } from 'redux-saga/effects'
import { FETCH_MEMBERS, LOADED_MEMBERS, FETCH_MEMBER, LOADED_MEMBER, API_ERROR, UPDATE_MEMBER, UPDATE_MEMBER_REQUEST } from '../constants/action-types'
import * as Request from '../util/request'
import { BASE_URL } from '../constants/config'
import { apiError, updateMemberSuccess } from '../actions/action'

// members
export function* watchMembers() {
    yield takeEvery(FETCH_MEMBERS, fetchMembers)
}

function* fetchMembers() {
    try {
        const payload = yield call(getMembers)
        yield put({type: LOADED_MEMBERS, payload})
    } catch(e) {
        yield put({type: API_ERROR, payload: e})
    }
}

function getMembers() {
    return Request.get(BASE_URL+"/users")
}

// member
export function* watchMember() {
    yield takeEvery(FETCH_MEMBER, fetchMember)
    yield takeEvery(UPDATE_MEMBER_REQUEST, updateMember)

}

function* fetchMember(action){
    try {
        const payload = yield call(getMember, action.payload.id)
        yield put({type: LOADED_MEMBER, payload})
    } catch(e) {
        yield put({type: API_ERROR, payload: e})
    }
}

function* updateMember(action) {
    try {
        const payload = yield call(putMember, action.payload.id, action.payload.body)
        yield put(updateMemberSuccess(payload))
    } catch(e) {
        yield put(apiError(e))
    }
}

function getMember(id) {
    return Request.get(BASE_URL+"/users/"+id)
}

function putMember(id, body) {
    return Request.put(BASE_URL+"/users/"+id, body)
}