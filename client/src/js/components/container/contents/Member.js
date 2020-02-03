import React from 'react'
import { Link } from 'react-router-dom'
import { connect } from 'react-redux'
import { fetchMembersRequest } from '../../../actions/action'

const mapDispatchToProps = dispatch => {
    return {
        fetchRequest: () => dispatch(fetchMembersRequest())
    }
}

const mapStateToProps = state => {
    return {
        isLoading: state.isLoading,
        members: state.members
    }
}

class ConnectedMember extends React.Component {
    componentDidMount(){
        this.props.fetchRequest()
    }

    render() {
        // TODO: isLoading
        console.log(this.props.members)
        return (
            <div className="content">
                <h1 className="content-title h1-block">メンバー</h1>
                <MemberGrade members={this.props.members}/>
                <MemberGraduate members={this.props.members}/>
            </div>
        )
    }
}

const MemberGrade = (props) => {
    const members = []
    // 学年わけ
    // TODO: 修士、博士を入れるなら9まで
    for(var i=2; i < 5; i++){
        members.push({
            grade: i,
            members: props.members.filter((member) => member.grade === i)
        })
    }
    return(
        <div className="flex-block">
            {
                members.map((memberList) => (
                    <MemberList key={memberList.grade} members={memberList} />
                ))
            }
        </div>
    )
}

const MemberList = (props) => {
    return (
        <div className="list items-3">
            {/* TODO: 修士以上が滅んでる */}
            <h3 className="list-title h3">{props.members.grade}年</h3> 
            <ul>
                {
                    props.members.members.map((member)=>(
                        <MemberRow member={member} key={member.id}/>
                    ))
                }
            </ul>
        </div>
    )
}

const MemberGraduate = (props) => {
    const members = props.members.filter((member) => member.grade === 0)
    return (
        <div className="list">
            <h3 className="list-title h3">卒業生</h3>
            <ul>
                {
                    members.map((member)=>(
                        <MemberRow key={member.id} member={member}/>
                    ))
                }
            </ul>
        </div>
    )
}

const MemberRow = (props) => {
    return (
        <li>
            <Link to={`/members/${props.member.id}`} className="list-item">{props.member.name}</Link>
        </li>
    )
}

const Member = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedMember)

export default Member