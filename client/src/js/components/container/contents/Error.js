import React from 'react'
import { connect } from 'react-redux'



const ConnectedError404 = () => {
    return(
        <div className="content">
            <h1 className="content-title h1-block">Page Not Found</h1>
            <p>お探しのページは見つかりませんでした</p>
        </div>
    )
}

const Error404 = connect(
    null,
)(ConnectedError404)

export default Error404