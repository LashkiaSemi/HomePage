import React from 'react'
import { connect } from 'react-redux'
import { fetchLectures } from '../../../actions/action'

const mapStateToProps = state => {
    return {
        isLoading: state.isLoading,
        lectures: state.lectures
    }
}

class ConnectedLecture extends React.Component {
    componentDidMount() {
        this.props.fetchLectures()
    }

    render() {
        return (
            <div className="content">
                <h1 className="content-title h1-block">レクチャー</h1>
                <div className="content-header">
                    <h2 className="h2">レクチャー資料一覧</h2>
                    {/* TODO: ボタン出しっぱ！だめ！ */}
                    <button className="btn btn-success">資料アップロード</button>
                </div>
                <LectureTable lectures={this.props.lectures}/>
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
                </tr>
            </thead>
            <tbody>
                {
                    props.lectures.map((lec) => (
                        <LectureRow key={lec.id} lecture={lec}/>
                    ))
                }
            </tbody>
        </table>
    )
}

const LectureRow = (props) => {
    return (
        <tr>
            {/* TODO: この辺あってる気がしない */}
            <td>{props.lecture.name}</td>
            <td>{props.lecture.author}</td>
            <td>{props.lecture.comment}</td>
            <td>{props.lecture.date}</td>
        </tr>
    )
}

const Lecture = connect(
    mapStateToProps,
    { fetchLectures }
)(ConnectedLecture)

export default Lecture