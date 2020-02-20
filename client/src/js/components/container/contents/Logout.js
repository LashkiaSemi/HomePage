import React from 'react'
import { logoutRequest } from '../../../actions/action'
import { connect } from 'react-redux'

const mapDispatchToProps = (dispatch) => {
    return {
        dispatchRequest: () => dispatch(logoutRequest())
    }
}

// ConnectedLogout ログアウト時に一瞬遷移する画面
class ConnectedLogout extends React.Component {
    componentDidMount(){
        this.props.dispatchRequest()
    }
    render() {
        return (
            <div className="content">
                <p>logout</p>
            </div>
        )
    }
}

const Logout = connect(
    null,
    mapDispatchToProps
)(ConnectedLogout)

export default Logout