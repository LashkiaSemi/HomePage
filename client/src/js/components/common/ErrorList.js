import React from 'react'

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