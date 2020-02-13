import React from 'react'
import { connect } from 'react-redux'
import { Link } from 'react-router-dom'
import BreadCrumb from './Breadcrumb'
import Modal from '../../common/Modal'
import { fetchActivitiesRequest, deleteActivityRequest } from '../../../actions/action'
import { findActivityByID } from '../../../util/findItem'

const mapStateToProps = (state) => {
    return {
        activities: state.activities,
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        fetchRequest: () => dispatch(fetchActivitiesRequest()),
        deleteRequest: (id) => dispatch(deleteActivityRequest({id})),
    }
}

class ConnectedActivityList extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            isInitialized: false,
            displayModal: false,
            selectedItemID: 0,
        }
        this.switchModal = this.switchModal.bind(this)
        this.handleDelete = this.handleDelete.bind(this)
    }

    componentDidMount() {
        this.props.fetchRequest()
    }

    switchModal(e) {
        this.setState({
            displayModal: !this.state.displayModal,
            selectedItemID: e.target.dataset.id
        })
    }

    handleDelete(){
        console.log(this.state.selectedItemID)
        this.props.deleteRequest(this.state.selectedItemID)
        this.setState({
            displayModal: false,
            selectedItemID: 0
        })
        // TODO: reload
    }

    render() {
        return (
            <div className="content">
                <BreadCrumb items={[{ path: "/", label: "管理者サイト" },{path: "/activities", label: "活動記録"}]}/>
                <div className="table-admin-caption">
                    <label>活動記録</label>
                    <Link className="btn btn-info" to={"/admin/activities/new"}>新規作成</Link>
                </div>
                <table className="table-admin table-admin-stripe">
                    <thead>
                        <tr>
                            <th>タイトル</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        {
                            this.props.activities.map(act=>(
                                <ListRow
                                    key={act.id}
                                    activity={act}
                                    handleSwitch={this.switchModal}/>
                            ))
                        }
                    </tbody>
                </table>
                {
                    this.state.displayModal
                    ? <DeleteModal
                        id={this.state.selectedItemID}
                        activities={this.props.activities}
                        handleSwitch={this.switchModal}
                        handleDelete={this.handleDelete}/>
                    : <></>
                }
            </div>
        )
    }
}

const ListRow = (props) => {
    return (
        <tr>
            <td>{props.activity.activity}</td>
            <td>
                <Link className="btn btn-primary" to={`/admin/activities/${props.activity.id}/edit`}>編集</Link>
                <button className="btn btn-danger" data-id={props.activity.id} onClick={props.handleSwitch}>削除</button>
            </td>
        </tr>
    )
}

const DeleteModal = (props) => {
    const activity = findActivityByID(props.activities, props.id)
    const modalBody = (
        <>
            <p><b>{activity.activity}</b>を削除します。よろしいですか。</p>
            <div>
                <button className="btn btn-danger" onClick={props.handleDelete}>削除</button>
                <button className="btn btn-info" onClick={props.handleSwitch}>キャンセル</button>
            </div>
        </>
    )
    return (
        <Modal
            title={"削除確認"}
            body={modalBody}
            handleSwitch={props.handleSwitch}/>
    )
}

const AdminActivityList = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedActivityList)

export default AdminActivityList