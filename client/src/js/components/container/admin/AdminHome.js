import React from 'react'
import { connect } from 'react-redux'

const mapStateToProps = (state) => {
    return {
        states: state
    }
}

class ConnectedAdminHome extends React.Component {
    constructor(props){
        super(props)
    }

    render() {
        // console.log(Object.keys(this.props.states))
        return (
            <div className="content">
                <table className="table table-stripe">
                    <caption>データ管理</caption>
                    <tbody>
                        {
                            TOCS.map(toc => (
                                <tr>
                                    <td>{toc.label}</td>
                                </tr>
                            ))
                        }
                    </tbody>
                </table>
            </div>
        )
    }
}

const AdminHome = connect(
    mapStateToProps
)(ConnectedAdminHome)

export default AdminHome

const TOCS = [
    {
        id: "activities",
        label: "活動記録",
    },
    {
        id: "societies",
        label: "学会発表",
    },
    {
        id: "researches",
        label: "卒業研究",
    },
    {
        id: "members",
        label: "メンバー",
    },
    {
        id: "jobs",
        label: "就職先",
    },
    {
        id: "equipments",
        label: "研究室備品",
    },
    {
        id: "lectures",
        label: "レクチャー",
    },
    
]