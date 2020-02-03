import React from 'react'
import { Link } from 'react-router-dom'
import { fetchMemberRequest } from '../../../actions/action'
import { connect } from 'react-redux'

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
    componentDidMount(){
        const id = this.props.match.params.id
        this.props.fetchRequest({id})
    }

    render() {
        // console.log(this.props.member)
        return (
            <div className="content">
                <div className="card mt-50">
                    <div className="card-content">
                        <div className="card-item">
                            <h2 className="h2">{this.props.member.name}</h2>
                            <label className="label-block col-gray">{this.props.member.student_id}</label>
                        </div>
                        <div className="card-item">
                            <label className="label-block">学年：{this.props.member.grade ? this.props.member.grade + "年" : "卒業生"}</label>
                        </div>
                        <div className="card-item">
                            <label className="label-block">学部：{this.props.member.department}</label>
                        </div>
                        <div className="card-item">
                            <label className="label-block">コメント：{this.props.member.comment}</label>
                        </div>
                        {/* 本人だった場合 */}
                        <div className="card-item">
                            <Link to={`/members/${this.props.member.id}/edit`} className="btn btn-success">編集する</Link>
                            <Link to={`/members/${this.props.member.id}/edit_pass`} className="btn btn-info">パスワードの変更</Link>
                        </div>
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

export default MemberDetail