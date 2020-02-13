import React from 'react'
import { connect } from 'react-redux'
import { fetchLecturesRequest, deleteLectureRequest } from '../../../actions/action'
import { STRAGE_KEY } from '../../../constants/config'
import { Link } from 'react-router-dom'
import * as Crypto from '../../../util/crypto'

const mapDispatchToProps = dispatch => {
    return {
        fetchRequest: () => dispatch(fetchLecturesRequest()),
        dispatchDeleteRequest: (id) => dispatch(deleteLectureRequest({id: id}))
    }
}

const mapStateToProps = state => {
    return {
        isLoading: state.isLoading,
        lectures: state.lectures
    }
}

class ConnectedLecture extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            isLogin: "",
            displayModal: false, // deleteの確認画面を開くアレ
        }

        this.switchModal = this.switchModal.bind(this)
        this.handleDelete = this.handleDelete.bind(this)
    }
    
    componentDidMount() {
        this.props.fetchRequest()
        this.setState({
            isLogin: Crypto.Decrypt(localStorage.getItem(STRAGE_KEY))
        })
    }

    switchModal(){
        this.setState({
            displayModal: !this.state.displayModal
        })
    }

    handleDelete(e) {
        this.props.dispatchDeleteRequest(e.target.dataset.id)
    }

    render() {
        return (
            <div className="content">
                <h1 className="content-title h1-block">レクチャー</h1>
                <div className="content-header">
                    <h2 className="h2">レクチャー資料一覧</h2>
                    {
                        this.state.isLogin
                            ? <Link className="btn btn-success" to="/lectures/new">資料アップロード</Link>
                            : <></>
                    }
                </div>
                <LectureTable 
                    lectures={this.props.lectures} 
                    isLogin={this.state.isLogin}
                    handleSwitch={this.switchModal}
                    handleDelete={this.handleDelete}
                    displayModal={this.state.displayModal}
                    />
            </div>
        )
    }
}

const LectureTable = (props) => {
    return (
        <table className="table-stripe">
            <thead>
                <tr>
                    <th>名前</th>
                    <th>投稿者</th>
                    <th>コメント</th>
                    <th>投稿日</th>
                    <th></th>
                </tr>
            </thead>
            <tbody>
                {
                    props.lectures.map((lec) => (
                        <LectureRow 
                            key={lec.id} 
                            lecture={lec} 
                            isLogin={props.isLogin}
                            handleSwitch={props.handleSwitch}
                            handleDelete={props.handleDelete}
                            displayModal={props.displayModal}/>
                    ))
                }
            </tbody>
        </table>
    )
}

const LectureRow = (props) => {
    return (
        <tr>
            <td>{props.lecture.title}</td>
            <td>{props.lecture.user.name}</td>
            <td>{props.lecture.comment}</td>
            <td>{props.lecture.updated_at}</td>
            {/* TODO: download script */}
            <td>
                <button className="btn btn-primary">Download</button>
                {
                    props.isLogin.indexOf(props.lecture.user.id) < 0 
                    ? <></>
                    : <>
                        <Link to={`/lectures/${props.lecture.id}/edit`} className="btn btn-info">編集</Link>
                        <button className="btn btn-danger" onClick={props.handleSwitch}>削除</button>
                    </>
                }
            </td>
            {
                props.displayModal
                    ? <DeleteModal
                        handleSwitch={props.handleSwitch}
                        handleDelete={props.handleDelete}
                        lecture={props.lecture}/>
                    : <></>
            }
        </tr>
    )
}

// TODO: commonあたりに移行させてもいいのでは？
const DeleteModal = (props) => {
    return (
        <div className="modal">
            <div className="modal-overlay">
                <div className="modal-content">
                    <div className="modal-title">
                        <label>削除確認</label>
                    </div>
                    <div className="modal-body">
                        <p><b>{props.lecture.title}</b> を削除します。よろしいですか。</p>
                        <div>
                            <button className="btn btn-danger" data-id={props.lecture.id} onClick={props.handleDelete}>削除</button>
                            <button className="btn btn-info" onClick={props.handleSwitch}>戻る</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

const Lecture = connect(
    mapStateToProps,
    mapDispatchToProps,
)(ConnectedLecture)

export default Lecture