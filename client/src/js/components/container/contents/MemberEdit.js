import React from 'react'
import { connect } from 'react-redux'
import { updateMemberRequest, fetchMember } from '../../../actions/action'

const mapStateToProps = (state) => {
    return {
        isLoading: state.isLoading,
        member: state.member
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        fetchRequest: (id) => dispatch(fetchMember({id: id})),
        dispatchRequest: (id, form) => dispatch(updateMemberRequest({id: id, body: form}))
    }
}

class ConnectedMemberEdit extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            name: "",
            studentID: "",
            grade: 2,
            department: "",
            comment: "",
        }
        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
    }

    componentDidMount() {
        this.props.fetchRequest(this.props.match.params.id)
    }

    handleChange(e) {
        const field = e.target.name
        this.setState({
            [field]: e.target.value
        })
    }

    handleSubmit(e) {
        e.preventDefault()
        const formData = {
            name: this.state.name,
            studentID: this.state.studentID,
            grade: this.state.grade,
            department: this.state.department,
            comment: this.state.comment
        }
        this.props.dispatchRequest(this.props.match.params.id, formData)
    }


    render() {
        return (
            <div className="content">
                <h1 className="content-title h1-block">アカウント情報編集</h1>
                <form className="form" onSubmit={this.handleSubmit}>
                    <div className="form-item">
                        <label className="input-label">名前</label>
                        <input type="text" className="input-text" name="name" value={this.state.name} onChange={this.handleChange} />
                    </div>

                    <div className="form-item">
                        <label className="input-label">学籍番号</label>
                        <input type="text" className="input-text" name="studentID" value={this.state.studentID} onChange={this.handleChange} />
                    </div>

                    <div className="form-item">
                        <label className="input-label">学年</label>
                        <select className="input-select" name="grade">
                            <option value="B2">学部2年</option>
                            <option value="B3">学部3年</option>
                            <option value="B4">学部4年</option>
                            <option value="M1">修士1年</option>
                            <option value="M2">修士2年</option>
                            <option value="D1">修士1年</option>
                            <option value="D2">修士2年</option>
                            <option value="0">卒業生</option>
                        </select>
                    </div>

                    <div className="form-item">
                        <label className="input-label">学部</label>
                        <input type="text" className="input-text" name="department" value={this.state.department} onChange={this.handleChange} />
                    </div>

                    <div className="form-item">
                        <label className="input-label">コメント</label>
                        <textarea className="input-textarea" name="comment" value={this.state.comment} onChange={this.handleChange}></textarea>
                    </div>

                    <div className="form-item">
                        <input type="submit" className="btn btn-primary" value="保存" />
                    </div>
                </form>
            </div>
        )
    }


}

const MemberEdit = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedMemberEdit)

export default MemberEdit