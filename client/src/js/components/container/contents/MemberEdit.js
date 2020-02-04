import React from 'react'
import { connect } from 'react-redux'
import { fetchAccountRequest, updateAccountRequest } from '../../../actions/action'

const mapStateToProps = (state) => {
    return {
        isLoading: state.isLoading,
        member: state.account
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        fetchRequest: () => dispatch(fetchAccountRequest()),
        dispatchRequest: (form) => dispatch(updateAccountRequest({body: form}))
    }
}

class ConnectedMemberEdit extends React.Component {
    constructor(props) {
        super(props)


        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
        this.initFromValue = this.initFromValue.bind(this)
    }

    componentDidMount() {
        this.props.fetchRequest()
    }

    componentDidUpdate(){
        if (!this.state && Object.keys(this.props.member).length) {
            this.initFromValue()
        }
    }

    initFromValue(){
        this.setState({
            name: this.props.member.name,
            studentID: this.props.member.student_id,
            grade: this.props.member.grade,
            department: this.props.member.department,
            comment: this.props.member.comment,
        })
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
            grade: parseInt(this.state.grade),
            department: this.state.department,
            comment: this.state.comment
        }
        this.props.dispatchRequest(formData)
    }


    render() {
        if(!this.state){
            return (
                <div className="content">
                    <label>Now Loading...</label>
                </div>
            )
        }
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
                        <select className="input-select" name="grade" value={this.state.grade} onChange={this.handleChange}>
                            <option value="2">学部2年</option>
                            <option value="3">学部3年</option>
                            <option value="4">学部4年</option>
                            <option value="5">修士1年</option>
                            <option value="6">修士2年</option>
                            <option value="7">修士1年</option>
                            <option value="8">修士2年</option>
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