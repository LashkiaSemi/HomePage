import { createStore, applyMiddleware } from 'redux'
import rootReducer from '../reducers/reducer'
import createSagaMiddleware from 'redux-saga'
import rootSaga from '../sagas/index'

// init saga
const sagaMiddleware = createSagaMiddleware()

// TODO: 第二引数に、initStateをつけられるらしい
// 初期化するならここ？
const store = createStore(
    rootReducer,
    applyMiddleware(sagaMiddleware)
)

sagaMiddleware.run(rootSaga)

export default store