import { takeEvery, call, put } from 'redux-saga/effects'
import { FETCH_EQUIPMENTS, LOADED_EQUIPMENTS, API_ERROR } from '../constants/action-types'
import { BASE_URL } from '../constants/config'
import { get } from '../util/request'

export function* watchEquipments() {
    yield takeEvery(FETCH_EQUIPMENTS, fetchEquipments)
}

function* fetchEquipments() {
    try {
        const payload = yield call(getEquipments)
        yield put({ type: LOADED_EQUIPMENTS, payload })
    } catch (e) {
        yield put({ type: API_ERROR, payload: e })
    }
}

function getEquipments() {
    return get(BASE_URL + "/equipments")
}