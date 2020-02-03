import { fork } from 'redux-saga/effects'

import { watchJobs } from './jobSaga'
import { watchMembers, watchMember } from './memberSaga'
import { watchActivities } from './activitySaga'
import { watchSocieties } from './societySaga'
import { watchResearches } from './researchSaga'
import { watchEquipments } from './equipmentSaga'
import { watchLectures } from './lectureSaga'
import { watchLogin } from './loginSaga'

// rootSaga
// sagaのwatcherをまとめておく
export default function* rootSaga() {
    yield fork(watchJobs)
    yield fork(watchMembers)
    yield fork(watchMember)
    yield fork(watchActivities)
    yield fork(watchSocieties)
    yield fork(watchResearches)
    yield fork(watchEquipments)
    yield fork(watchLectures)
    yield fork(watchLogin)
}