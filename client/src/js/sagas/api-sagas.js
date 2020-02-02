import { fork } from 'redux-saga/effects'
import { watchJobs } from './api-jobs'

export default function* rootSaga() {
    yield fork(watchJobs)
}