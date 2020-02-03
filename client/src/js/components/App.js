import React from 'react'
import { BrowserRouter } from 'react-router-dom'

import '../../css/homepage.css'
import Header from './header/Header'
import Container from './container/Container'
import Footer from './footer/Footer'

class App extends React.Component {

    render (){
        return (
            <>
                <BrowserRouter>
                    <Header />
                    <Container />
                    <Footer />
                </BrowserRouter>
            </>
        )
    }
}
export default App