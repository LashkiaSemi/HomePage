// import React from 'react'

// // 使っていない気がする...
// // TODO: adminEdit: フォームを自動生成。tbodyまで守備範囲
// export function generateInputField() {
//     const fields = []
//     FIELDS.map(field => {
//         const cName = "input-admin-" + field.type
//         fields.push(
//             <tr className="form-admin-item">
//                 <td><label className="input-admin-label">{field.label}</label></td>
//                 <td><input type={field.type} className={cName} name={field.name} value={field.value} onChange={field.onChange} /></td>
//             </tr>
//         )
//     })
//     return fields
// }

// function onChange(){
//     console.log("onChange!")
// }

// const FIELDS = [
//     {
//         label: "活動内容",
//         type: "text",
//         name: "activity",
//         value: "sample value",
//         onChange: onChange
//     },
//     {
//         label: "日付",
//         type: "date",
//         name: "date",
//         value: "",
//         onChange: onChange
//     },
// ]


// // TODO: adminEdit: フォームの初期化