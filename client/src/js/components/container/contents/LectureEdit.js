import React from 'react'
import { connect } from 'react-redux'
import { createLectureRequest, fetchLectureRequest, updateLectureRequest } from '../../../actions/action'
import ErrorList from '../../common/ErrorList'
import { STRAGE_KEY } from '../../../constants/config'
import * as Crypto from '../../../util/crypto'

const mapDispatchToProps = (dispatch) => {
    return {
        fetchRequest: (lecID) => dispatch(fetchLectureRequest({id: lecID})),
        dispatchCreateRequest: form => dispatch(createLectureRequest({body: form})),
        dispatchUpdateRequest: (id, form) => dispatch(updateLectureRequest({id: id, body: form})),
    }
}

const mapStateToProps = (state) => {
    return {
        isLoading: state.isLoading,
        lecture: state.lecture
    }
}

class ConnectedLectureEdit extends React.Component {
    constructor(props){
        super(props)
        this.state = {
            id: props.match.params.id,
            init: false,
            title: '',
            comment: '',
            isPublic: true,
            errors: [],
        }

        this.fileInput = React.createRef()

        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
        this.handleSwitchCheck = this.handleSwitchCheck.bind(this)
    }

    componentDidMount(){
        if(typeof this.state.id === 'undefined') {
            this.setState({
                init: true,
            })
            return
        }
        this.props.fetchRequest(this.state.id)
    }

    componentDidUpdate() {
        // 初期化
        if (!this.state.init && Object.keys(this.props.lecture).length) {
            if (this.props.lecture.id == this.state.id) {
                this.setState({
                    title: this.props.lecture.title,
                    comment: this.props.lecture.comment,
                    init: true,
                })
            }
        }
    }

    handleChange(e) {
        const field = e.target.name
        this.setState({
            [field]: e.target.value
        })
    }

    handleSwitchCheck(e) {
        const field = e.target.name
        this.setState({
            [field]: !this.state.isPublic
        })
    }

    async handleSubmit(e) {
        e.preventDefault()
        // TODO: 空値チェックはutilに移行
        var errors = []
        if(this.state.title === "") {
            errors.push({id: "titleEmpty", content: "タイトルは必須です"})
        }
        if (!this.state.id) {
            if (typeof this.fileInput.current.files[0] === 'undefined') {
                errors.push({ id: "fileEmpty", content: "ファイルは必須です" })
            }
        }

        if(errors.length > 0) {
            this.setState({
                errors
            })
            return
        }

        const body = {
            title: this.state.title,
            comment: this.state.comment,
            is_public: this.state.isPublic,
            user_id: parseInt(Crypto.Decrypt(localStorage.getItem(STRAGE_KEY)))
        }

        const formData = new FormData()
        formData.append("body", JSON.stringify(body))
        formData.append("file", this.fileInput.current.files[0])

        if(!this.state.id) {
            this.props.dispatchCreateRequest(formData)
        } else {
            this.props.dispatchUpdateRequest(this.state.id, formData)
        }
        // TODO: stateの初期化？
    }

    render(){
        return (
            <div className="content">
                <h1 className="content-title h1-block">レクチャー資料アップロード</h1>
                <form className="form" encType="multipart/form-data" onSubmit={this.handleSubmit}>
                    {
                        this.state.errors.length
                            ? <ErrorList errors={this.state.errors} />
                            : <></>
                    }
                    <div className="form-item">
                        <label className="input-label">タイトル</label>
                        <input type="text" className="input-text" name="title" value={this.state.title} onChange={this.handleChange} placeholder="例：xx年度 xx言語資料1"/>
                    </div>

                    <div className="form-item">
                        <label className="input-label">コメント</label>
                        <input type="text" className="input-text" name="comment" value={this.state.comment} onChange={this.handleChange} />
                    </div>

                    <div className="form-item">
                        <label className="input-label">ファイル</label>
                        <input type="file" ref={this.fileInput}/>
                    </div>

                    <div className="form-item">
                        <label className="mr-10">公開 / 非公開</label>
                        <input type="checkbox" checked={this.state.isPublic} name="isPublic" onChange={this.handleSwitchCheck} />
                    </div>

                    <div className="form-item">
                        <input type="submit" className="btn btn-primary" value="保存" />
                    </div>
                </form>
            </div>
        )
    }
}

const LectureEdit = connect(
    mapStateToProps,
    mapDispatchToProps,
)(ConnectedLectureEdit)

export default LectureEdit