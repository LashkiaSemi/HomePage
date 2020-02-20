import React from 'react'
import { loginRequest } from '../../../actions/action'
import { connect } from 'react-redux'
import { APIErrorList } from '../../common/APIError'

const mapStateToProps = state => {
    return {
        isLoading: state.isLoading,
        login: state.login,
        apiError: state.apiError,
    }
}

const matDispatchToProps = dispatch => {
    return {
        dispatchRequest: (id, password) => dispatch(loginRequest({id, password}))
    }
}

class ConnectedLogin extends React.Component {
    render() {
        return (
            <div className="content">
                <h1 className="content-title h1-block">ログイン</h1>
                <LoginForm
                    dispatchRequest={this.props.dispatchRequest}
                    apiError={this.props.apiError} />
            </div>
        )
    }
}

/*
LoginFrom ログインするためのフォーム
props:
    apiError           = apiで起きたエラー
    deispatchRequest() = ログインリクエスト
*/
class LoginForm extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            studentID: '',
            password: '',
        }
        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
    }

    handleChange(e) {
        const field = e.target.name
        this.setState({
            [field]: e.target.value
        })
    }

    handleSubmit(e) {
        e.preventDefault()
        this.props.dispatchRequest(this.state.studentID, this.state.password)
    }

    componentDidMount(){
        // 必要なのか...
        this.setState({
            studentID: '',
            password: ''
        })
    }

    render() {
        return (
            <form className="form">
                <APIErrorList
                    apiError={this.props.apiError}/>

                <h3 className="input-label">学籍番号</h3>
                <input type="text" className="input-text" name="studentID" onChange={this.handleChange} value={this.state.studentID} />
                <h3 className="input-label">パスワード</h3>
                <input type="password" className="input-text" name="password" onChange={this.handleChange} value={this.state.password} />

                <input type="submit" value="ログイン" className="btn btn-primary mt-20" onClick={this.handleSubmit} />
            </form>
        )
    }
}


const Login = connect(
    mapStateToProps,
    matDispatchToProps,
)(ConnectedLogin)

export default Login