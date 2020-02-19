import React from 'react'
import { connect } from 'react-redux'
import AdminEdit from '../AdminEdit'
import BreadCrumb from '../../../common/Breadcrumb'
import { fetchEquipmentsRequest, createEquipmentRequest, updateEquipmentRequest, fetchTagsRequest } from '../../../../actions/action'
import { APIErrorList } from '../../../common/APIError'

const mapStateToProps = state => {
    return {
        equipments: state.equipments,
        tags: state.tags,
        apiError: state.apiError
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        fetchRequest: () => dispatch(fetchEquipmentsRequest()),
        fetchTagsRequest: () => dispatch(fetchTagsRequest()),
        createRequest: (body) => dispatch(createEquipmentRequest({ body })),
        updateRequest: (id, body) => dispatch(updateEquipmentRequest({ id, body }))
    }
}

class ConnectedEquipmentEdit extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            values: {
                name: "",
                stock: 0,
                note: "",
                tag_id: 0
            },
            fields: [
                { label: "備品名", type: "text", name: "name", required: true },
                { label: "数量", type: "number", name: "stock", requestType: "int" },
                { label: "備考", type: "text", name: "note" },
            ],
            isInitialized: false,
        }
    }

    componentDidMount() {
        this.props.fetchTagsRequest()
    }

    componentDidUpdate() {
        if (this.state.isInitialized) {
            return
        }
        if (Object.keys(this.props.tags).length > 0) {
            var options = []
            this.props.tags.map(tag => {
                options.push({ label: tag.name, value: tag.id})
            })

            this.setState({
                fields: this.state.fields.concat({ label: "タグ", type: "select", name: "tag_id", requestType: "int", required: true, options },),
                isInitialized: true
            })
        }
    }

    render() {
        return (
            <div className="content">
                <BreadCrumb items={[{ path: "/", label: "管理者サイト" }, { path: "/equipments", label: "研究室備品" }]} />
                <APIErrorList
                    apiError={this.props.apiError}/>
                {
                    !this.state.isInitialized
                    ? <></>
                    : <AdminEdit
                        items={this.props.equipments}
                        itemID={this.props.match.params.id}
                        fields={this.state.fields}
                        values={this.state.values}
                        fetchRequest={this.props.fetchRequest}
                        createRequest={this.props.createRequest}
                        updateRequest={this.props.updateRequest} />
                }
            </div>
        )
    }
}

const AdminEquipmentEdit = connect(
    mapStateToProps,
    mapDispatchToProps
)(ConnectedEquipmentEdit)

export default AdminEquipmentEdit
