import React from 'react'

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