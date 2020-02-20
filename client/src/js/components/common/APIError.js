import React from 'react'
import ErrorList from './ErrorList'

/*
APIErrorList api側のエラーをリストっぽく表示する
props:
    apiError = apiのエラーレスポンス。{ error: {} }って感じ
               いまのところerrorに入ってるのはaxiosのレスポンスが丸々
               実装は reducer > error.js

*/
export const APIErrorList = (props) => {
    if(props.apiError === null || typeof props.apiError === 'undefined') {
        return <></>
    }
    var errorTemplate = (
        Object.keys(props.apiError).length > 0
            ? <ErrorList
                errors={[{ id: 500, content: "API接続エラー" },]} /> // TODO: idがきしょい
            : <></>
    )
    if (typeof props.apiError.error !== 'undefined') {
        errorTemplate = (
            Object.keys(props.apiError).length > 0
                ? <ErrorList
                    errors={[{ id: props.apiError.error.data.code, content: props.apiError.error.data.message },]} />
                : <></>
        )
    }
    return errorTemplate
}

