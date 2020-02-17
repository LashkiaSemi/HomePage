import React from 'react'
import { Redirect } from 'react-router-dom'
import { STRAGE_KEY } from '../../constants/config'
import * as Crypto from '../../util/crypto'

const Admin = (props) => {
    const strageValue = localStorage.getItem(STRAGE_KEY)
    if (!strageValue) {
        return <Redirect to="/login" />
    }
    const loginInfo = Crypto.Decrypt(strageValue)
    return (
        loginInfo.indexOf("owner") > -1
        ? props.children
        : <Redirect to="/login" />
    )
}

export default Admin