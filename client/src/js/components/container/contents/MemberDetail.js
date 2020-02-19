import React from 'react'
import { Link } from 'react-router-dom'
import { fetchMembersRequest } from '../../../actions/action'
import { connect } from 'react-redux'
import { STRAGE_KEY } from '../../../constants/config'
import * as Crypto from '../../../util/crypto'
import { findItemByID } from '../../../util/findItem'
import { APIErrorList } from '../../common/APIError'

const mapDispatchToProps = dispatch => {
    return {
        fetchRequest: () => dispatch(fetchMembersRequest())
    }
}

const mapStateToProps = state => {
    return {
        isLoading: state.isLoading,
        members: state.members,
        apiError: state.apiError
    }
}

class ConnectedMemberDetail extends React.Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            isEdit: false,
            isInitialized: false,
            member: {},
        }
    }

    componentDidMount(){
        this.props.fetchRequest()
        if(localStorage.getItem(STRAGE_KEY) !== null) {
            this.setState({
                isEdit: (Crypto.Decrypt(localStorage.getItem(STRAGE_KEY)).indexOf(this.props.match.params.id) > -1),
            })
        }
    }

    componentDidUpdate(){
        if (this.state.isInitialized) {
            return
        }
        if (this.props.members.length > 0) {
            this.setState({
                member: findItemByID(this.props.members, this.state.id),
                isInitialized: true
            })
        }
    }

    render() {
        return (
            <div className="content">
                <APIErrorList
                    apiError={this.props.apiError}/>
                <MemberCard 
                    member={this.state.member}
                    isEdit={this.state.isEdit}/>
            </div>
        )
    }
}

const MemberCard = (props) => {
    return (
        <div className="card mt-50">
            <div className="card-content">
                <div className="card-item">
                    <h2 className="h2">{props.member.name}</h2>
                    <label className="label-block col-gray">{props.member.student_id}</label>
                </div>
                <div className="card-item">
                    <label className="label-block">学年：{numToGrade(props.member.grade)}</label>
                </div>
                <div className="card-item">
                    <label className="label-block">学部：{props.member.department}</label>
                </div>
                <div className="card-item">
                    <label className="label-block">コメント：{props.member.comment}</label>
                </div>
                {
                    !props.isEdit
                        ? <></>
                        : <div className="card-item">
                            <Link to={`/account/edit`} className="btn btn-success">編集する</Link>
                            <Link to={`/account/edit_pass`} className="btn btn-info">パスワードの変更</Link>
                        </div>

                }
            </div>
        </div>
    )
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