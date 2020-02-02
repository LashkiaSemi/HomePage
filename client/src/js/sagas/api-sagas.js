import { fork } from 'redux-saga/effects'
import { watchJobs } from './api-jobs'
import { watchMembers } from './api-members'

// rootSaga
// sagaのwatcherをまとめておく
export default function* rootSaga() {
    yield fork(watchJobs)
    yield fork(watchMembers)

}