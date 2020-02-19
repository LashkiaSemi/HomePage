import React from 'react'
import { Link } from 'react-router-dom'
import { findItemByID, findCaptionByDataType } from '../../../util/findItem'
import Modal from '../../common/Modal'
import { APIErrorList } from '../../common/APIError'

class AdminList extends React.Component {
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

    handleDelete() {
        this.props.deleteRequest(this.state.selectedItemID)
        this.setState({
            displayModal: false,
            selectedItemID: 0
        })
        // TODO: reload
    }

    render(){
        
        return (
            <>
                <div className="table-admin-caption">
                    <label>{this.props.caption}</label>
                    <Link className="btn btn-info" to={`/admin/${this.props.path}/new`}>新規作成</Link>
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
                            this.props.items.map(item => (
                                <ListRow
                                    key={item.id}
                                    item={item}
                                    path={this.props.path}
                                    handleSwitch={this.switchModal} />
                            ))
                        }
                    </tbody>
                </table>
                {
                    this.state.displayModal
                        ? <DeleteModal
                            id={this.state.selectedItemID}
                            items={this.props.items}
                            path={this.props.path}
                            handleSwitch={this.switchModal}
                            handleDelete={this.handleDelete} />
                        : <></>
                }
            </>
        )
    }
}

const ListRow = (props) => {
    const title = findCaptionByDataType(props.path, props.item)
    return (
        <tr>
            <td>{title}</td>
            <td className="al-right">
                <Link className="btn btn-primary" to={`/admin/${props.path}/${props.item.id}/edit`}>編集</Link>
                <button className="btn btn-danger" data-id={props.item.id} onClick={props.handleSwitch}>削除</button>
            </td>
        </tr>
    )
}

const DeleteModal = (props) => {
    const item = findItemByID(props.items, props.id)
    const title = findCaptionByDataType(props.path, item)
    const modalBody = (
        <>
            <p><b>{title}</b>を削除します。よろしいですか。</p>
            <div className="al-right mt-20">
                <button className="btn btn-danger" onClick={props.handleDelete}>削除</button>
                <button className="btn btn-info" onClick={props.handleSwitch}>キャンセル</button>
            </div>
        </>
    )
    return (
        <Modal
            title={"削除確認"}
            body={modalBody}
            handleSwitch={props.handleSwitch} />
    )
}

export default AdminList