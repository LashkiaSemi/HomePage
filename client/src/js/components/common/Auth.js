import React from 'react'
import { Redirect } from 'react-router-dom'
import { STRAGE_KEY } from '../../constants/config'

const Auth = (props) => {
    const isLogin = localStorage.getItem(STRAGE_KEY)
    return (
        isLogin
        ? props.children
        : <Redirect to="/login" />
    )
}

export default Auth