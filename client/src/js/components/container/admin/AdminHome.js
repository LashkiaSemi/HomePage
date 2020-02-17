import React from 'react'
import { connect } from 'react-redux'
import { Link } from 'react-router-dom'

import BreadCrumb from '../../common/Breadcrumb'

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
        return (
            <div className="content">
                <BreadCrumb items={[{path: "/", label: "管理者サイト"}]}/>

                <div className="list-admin">
                    <label className="list-admin-title">データ管理</label>
                    <ul>
                        {
                            TOCS.map(toc => (
                                <li key={toc.id}><Link to={`/admin/${toc.id}`} className="list-admin-item">{toc.label}</Link></li>
                            ))
                        }
                    </ul>
                </div>
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
    {
        id: "tags",
        label: "タグ"
    }
    
]