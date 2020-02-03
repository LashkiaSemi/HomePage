import { takeEvery, put, call } from 'redux-saga/effects'
import { UPDATE_MEMBER_REQUEST, FETCH_MEMBERS_REQUEST, FETCH_MEMBER_REQUEST } from '../constants/action-types'
import * as Request from '../util/request'
import { BASE_URL } from '../constants/config'
import { updateMemberFailure, updateMemberSuccess, 
    fetchMembersSuccess, fetchMembersFailure, 
    fetchMemberSuccess, fetchMemberFailure
    } from '../actions/action'

// members
export function* watchMembers() {
    yield takeEvery(FETCH_MEMBERS_REQUEST, fetchMembers)
}

function* fetchMembers() {
    try {
        const payload = yield call(getMembers)
        yield put(fetchMembersSuccess(payload))
    } catch(e) {
        yield put(fetchMembersFailure(e))
    }
}

function getMembers() {
    return Request.get(BASE_URL+"/users")
}

// member
export function* watchMember() {
    yield takeEvery(FETCH_MEMBER_REQUEST, fetchMember)
    yield takeEvery(UPDATE_MEMBER_REQUEST, updateMember)

}

function* fetchMember(action){
    try {
        const payload = yield call(getMember, action.payload.id)
        yield put(fetchMemberSuccess(payload))
    } catch(e) {
        yield put(fetchMemberFailure(e))
    }
}

function* updateMember(action) {
    try {
        const payload = yield call(putMember, action.payload.id, action.payload.body)
        yield put(updateMemberSuccess(payload))
    } catch(e) {
        yield put(updateMemberFailure(e))
    }
}

function getMember(id) {
    return Request.get(BASE_URL+"/users/"+id)
}

function putMember(id, body) {
    return Request.put(BASE_URL+"/users/"+id, body)
}