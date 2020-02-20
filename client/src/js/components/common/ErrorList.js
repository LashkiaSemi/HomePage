import React from 'react'

/*
ErrorList 汎用のエラーリスト
props:
    errors = リストに表示するエラーを配列で
        [{ id     : 一意の値
           content: 表示内容 }, ...]
*/
const ErrorList = (props) => {
    return (
        <div className="error-list mb-20">
            {
                props.errors.map(err => (
                    <li key={err.id}>{err.content}</li>
                ))
            }
        </div>
    )
}

export default ErrorList