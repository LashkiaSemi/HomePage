import React from 'react'
import { API_URL } from '../../../constants/config'

const Download = (props) => {
    const link = "http://localhost:8000" + "/logfile.log"
    return (
        <div className="content">
            <a href={link} download="download-logfile.log">Download</a>
        </div>
    )
}

export default Download
