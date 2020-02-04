import React from 'react'
import { connect  } from 'react-redux'
import { updateAccountRequest } from '../../../actions/action'

const mapDispatchToProps = (dispatch) => {
    return {
        dispatchRequest: (body) => dispatch(updateAccountRequest({body}))
    }
}

class ConnectedPasswordEdit extends React.Component {
    constructor(props) {
        super(props)

        this.state = {
            oldPass: '',
            newPass: '',
            confirmPass: '',
        }

        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
    }


    handleChange(e){
        const field = e.target.name
        this.setState({
            [field]: e.target.value
        })
    }

    handleSubmit(e) {
        e.preventDefault()
        // confirmあたりの検証
        if(this.state.newPass !== this.state.confirmPass) {
            // TODO: 警告！
            console.log("確認が不一致")
            return
        }

        // datasetの作成
        const body = {
            old_password: this.state.oldPass,
            password: this.state.newPass
        }
        // stateの初期化
        this.setState({
            oldPass: '',
            newPass: '',
            confirmPass: '',
        })
        // dispatch
        // dispatchはするけど、古いあれが更新されてないとwarning出す必要が。
        this.props.dispatchRequest(body)
    }

    render() {
        return (
            <div className="content">
                <h1 className="content-title h1-block">パスワードの再設定</h1>
                <form className="form" onSubmit={this.handleSubmit}>
                    <div className="form-item">
                        <label className="input-label">元のパスワード</label>
                        <input type="password" name="oldPass" className="input-text" value={this.state.oldPass} onChange={this.handleChange} />
                    </div>

                    <div className="form-item">
                        <label className="input-label">新しいパスワード</label>
                        <input type="password" name="newPass" className="input-text mb-10" value={this.state.newPass} onChange={this.handleChange} placeholder="新規パスワード" />
                        <input type="password" name="confirmPass" className="input-text" value={this.state.confirmPass} onChange={this.handleChange} placeholder="確認用" />
                    </div>

                    <div className="form-item">
                        <input type="submit" className="btn btn-primary" value="保存" />
                    </div>

                </form>
            </div>
        )
    }
}

const PasswordEdit = connect(
    null,
    mapDispatchToProps
)(ConnectedPasswordEdit)

export default PasswordEdit