import React from 'react'
import { Link } from 'react-router-dom'
import { fetchMemberRequest } from '../../../actions/action'
import { connect } from 'react-redux'
import { STRAGE_KEY } from '../../../constants/config'

const mapDispatchToProps = dispatch => {
    return {
        fetchRequest: (id) => dispatch(fetchMemberRequest({id}))
    }
}

const mapStateToProps = state => {
    return {
        isLoading: state.isLoading,
        member: state.member
    }
}

class ConnectedMemberDetail extends React.Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            isEdit: localStorage.getItem(STRAGE_KEY) === this.props.match.params.id
        }
    }

    componentDidMount(){
        this.props.fetchRequest(this.state.id)
    }

    render() {
        return (
            <div className="content">
                <div className="card mt-50">
                    <div className="card-content">
                        <div className="card-item">
                            <h2 className="h2">{this.props.member.name}</h2>
                            <label className="label-block col-gray">{this.props.member.student_id}</label>
                        </div>
                        <div className="card-item">
                            <label className="label-block">学年：{numToGrade(this.props.member.grade)}</label>
                        </div>
                        <div className="card-item">
                            <label className="label-block">学部：{this.props.member.department}</label>
                        </div>
                        <div className="card-item">
                            <label className="label-block">コメント：{this.props.member.comment}</label>
                        </div>
                        {
                            !this.state.isEdit
                            ? <></>
                            : <div className="card-item">
                                <Link to={`/account/edit`} className="btn btn-success">編集する</Link>
                                <Link to={`/account/edit_pass`} className="btn btn-info">パスワードの変更</Link>
                            </div>
                            
                        }
                    </div>
                </div>
            </div>
        )
    }
}

const MemberDetail = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedMemberDetail)

const numToGrade = (num) => {
    switch(num) {
        case 2:
            return "学部2年"
        case 3:
            return "学部3年"
        case 4:
            return "学部4年"
        case 5:
            return "修士1年"
        case 6:
            return "修士2年"
        case 7:
            return "博士1年"
        case 8:
            return "博士2年"
        case 0:
            return "卒業生"
        default:
            return ""
    }
}


export default MemberDetail