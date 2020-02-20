import React from 'react'
import { Link } from 'react-router-dom'



/*
BreadCrumb パンくずリスト
props:
    path  = admin以降のurl. ex) /activities
    label = リストに表示する文字
*/
const BreadCrumb = (props) => {
    return (
        <div className="breadcrumb mb-30">
            {
                props.items.map(item=>(
                    <span key={item.path}>
                        <span>/</span><Link to={`/admin${item.path}`}>{item.label}</Link>
                    </span>
                ))
            }
        </div>
    )
}

export default BreadCrumb