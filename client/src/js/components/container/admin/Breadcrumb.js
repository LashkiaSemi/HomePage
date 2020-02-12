import React from 'react'
import { Link } from 'react-router-dom'

const BreadCrumb = (props) => {
    return (
        <div className="breadcrumb mb-30">
            {
                props.items.map(item=>(
                    <>
                        <span>/</span><Link to={`/admin/${item.path}`}>{item.label}</Link>
                    </>
                ))
            }
        </div>
    )
}

export default BreadCrumb