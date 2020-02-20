import React from 'react'

/*
Modal モーダルのベース
props:
    title          = モーダルのヘッダに描かれるタイトル
    body           = モーダルの本文。JSXかHTMLかな
    handleSwitch() = モーダルの表示切り替え。呼び出し元で定義する必要がある
*/
const Modal = (props) => {
    return (
        <div className="modal">
            <div className="modal-overlay">
                <div className="modal-content">
                    <div className="modal-title">
                        <label>{props.title}</label>
                        <button className="btn modal-close" onClick={props.handleSwitch}>✖️</button>
                    </div>
                    <div className="modal-body">
                        {props.body}
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Modal