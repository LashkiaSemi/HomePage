import React from 'react'
import { Link } from 'react-router-dom'
import { STRAGE_KEY } from '../../constants/config'

// Header
// 頭のナビゲーションバーです
class Header extends React.Component {
    constructor(props) {
        super(props)
    }
    
    render() {
        return (
            <div className="header">
                <Navigation />
            </div>
        )
    }
}

// Navigation
// ナビゲーションバー
const Navigation = (props) => (
    <nav>
        <Logo />
        <Menu />
    </nav>
)

// Logo
// らしきあぜみって書いておく場所
const Logo = (props) => (
    <div className="logo">
        <Link to="/">Lashkia研究室</Link>
    </div>
)

// Menu
const Menu = (props) => {
    // localStrageに値がなかったらnullになります
    const isLogin = localStorage.getItem(STRAGE_KEY)
    return (
        <ul className="menu">
            {
                MENU.map((menu) => (
                    isLogin
                    ? menu.isLogin
                        ? <MenuRow key={menu.id} menu={menu} />
                        : <></>
                    : menu.isNotLogin
                        ? <MenuRow key={menu.id} menu={menu} />
                        : <></>
                ))
            }
        </ul>
    )
}

// MenuRow
// 項目一個
const MenuRow = (props) => {
    return (
        <li><Link to={`/${props.menu.id}`}>{props.menu.display}</Link></li>
    )
}

// Static datas
const MENU = [
    {
        id: "activities",
        display: "活動記録",
        isLogin: true,
        isNotLogin: true,
    },
    {
        id: "societies",
        display: "学会発表",
        isLogin: true,
        isNotLogin: true,
    },
    {
        id: "researches",
        display: "卒業研究",
        isLogin: true,
        isNotLogin: true,
    },
    {
        id: "jobs",
        display: "就職先",
        isLogin: true,
        isNotLogin: true,
    },
    {
        id: "members",
        display: "メンバー",
        isLogin: true,
        isNotLogin: true,
    },
    {
        id: "links",
        display: "外部リンク",
        isLogin: true,
        isNotLogin: true,
    },
    {
        id: "equipments",
        display: "研究室備品",
        isLogin: true,
        isNotLogin: false,
    },
    {
        id: "lectures",
        display: "レクチャー",
        isLogin: true,
        isNotLogin: false,
    },
    {
        id: "login",
        display: "ログイン",
        isLogin: false,
        isNotLogin: true,
    },
    {
        id: "logout",
        display: "ログアウト",
        isLogin: true,
        isNotLogin: false,
    },
]

export default Header

