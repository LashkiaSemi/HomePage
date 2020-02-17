import { takeEvery, call, put } from 'redux-saga/effects'
import { FETCH_EQUIPMENTS_REQUEST, CREATE_EQUIPMENT_REQUEST, UPDATE_EQUIPMENT_REQUEST, DELETE_EQUIPMENT_REQUEST } from '../constants/action-types'
import { API_URL } from '../constants/config'
import * as Request from '../util/request'
import { fetchEquipmentsSuccess, fetchEquipmentsFailure, updateEquipmentSuccess, createEquipmentSuccess, createEquipmentFailure, updateEquipmentFailure, deleteEquipmentRequest, deleteEquipmentSuccess, deleteEquipmentFailure } from '../actions/action'

// watcher
export function* watchEquipments() {
    yield takeEvery(FETCH_EQUIPMENTS_REQUEST, fetchEquipments)
    yield takeEvery(CREATE_EQUIPMENT_REQUEST, createEquipment)
    yield takeEvery(UPDATE_EQUIPMENT_REQUEST, updateEquipment)
    yield takeEvery(DELETE_EQUIPMENT_REQUEST, removeEquipment)
}

// worker
function* fetchEquipments() {
    try {
        const payload = yield call(getEquipments)
        yield put(fetchEquipmentsSuccess(payload))
    } catch (e) {
        yield put(fetchEquipmentsFailure(e))
    }
}

function* createEquipment(action) {
    try {
        const payload = yield call(postEquipment, action.payload.body)
        yield put(createEquipmentSuccess(payload))
    } catch(e) {
        yield put(createEquipmentFailure(e))
    }
}

function* updateEquipment(action) {
    try {
        const payload = yield call(putEquipment, action.payload.id, action.payload.body)
        yield put(updateEquipmentSuccess(payload))
    } catch(e) {
        yield put(updateEquipmentFailure(e))
    }
}

function* removeEquipment(action) {
    try {
        const payload = yield call(deleteEquipment, action.payload.id)
        yield put(deleteEquipmentSuccess(payload))
    } catch(e) {
        yield put(deleteEquipmentFailure(e))
    }
}

// api call
function getEquipments() {
    const options = { withCredentials: true }
    return Request.httpGet(API_URL + "/equipments", options)
}

function postEquipment(body) {
    const options = { withCredentials: true }
    return Request.httpPost(API_URL+"/equipments", options)
}

function putEquipment(id, body) {
    const options = { withCredentials: true }
    return Request.httpPut(API_URL+"/equipments/"+id, body, options)
}

function deleteEquipment(id) {
    const options = { withCredentials: true }
    return Request.httpDelete(API_URL+"/equipments/"+id, options)
}