import React from 'react'
import { connect } from 'react-redux'
import { fetchLecturesRequest } from '../../../actions/action'
import { STRAGE_KEY } from '../../../constants/config'
import { Link } from 'react-router-dom'


const mapDispatchToProps = dispatch => {
    return {
        fetchRequest: () => dispatch(fetchLecturesRequest())
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
            isLogin: null
        }
    }
    
    componentDidMount() {
        this.props.fetchRequest()
        this.setState({
            isLogin: localStorage.getItem(STRAGE_KEY)
        })
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
                <LectureTable lectures={this.props.lectures} isLogin={this.state.isLogin}/>
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
                        <LectureRow key={lec.id} lecture={lec} isLogin={props.isLogin}/>
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
                    props.isLogin != props.lecture.user.id
                    ? <></>
                    : <>
                        <Link to={`/lectures/${props.lecture.id}/edit`} className="btn btn-info">編集</Link>
                        <button className="btn btn-danger">削除</button>
                    </>
                }
            </td>

        </tr>
    )
}

const Lecture = connect(
    mapStateToProps,
    mapDispatchToProps,
)(ConnectedLecture)

export default Lecture