import React from 'react'
import { Link } from 'react-router-dom'
import { fetchMemberRequest } from '../../../actions/action'
import { connect } from 'react-redux'
import { STRAGE_KET } from '../../../constants/config'

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
            isEdit: localStorage.getItem(STRAGE_KET) === this.props.match.params.id
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
                            <label className="label-block">学年：{GRADE[this.props.member.grade]}</label>
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

// todo: switchcaseなfunctionでもいいかも
const GRADE = {
    2: "学部2年",
    3: "学部3年",
    4: "学部4年",
    5: "修士1年",
    6: "修士2年",
    7: "博士1年",
    8: "博士2年",
    0: "卒業生"
}

export default MemberDetail