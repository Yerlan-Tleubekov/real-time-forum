import './css/normalize.css'
import { createElementObj } from "./js/utils/create"
import { tagsOptions } from './js/utils/options'
import Header from './js/components/header'
import AuthModal from './js/pages/auth'



function MainLayout() {
    const layout = document.getElementById("body")
    const header = createElementObj(tagsOptions.header)
    const main = createElementObj(tagsOptions.main)
    const footer = createElementObj(tagsOptions.footer)

    return {
        init: () => {
            layout.append(header, main, footer)

            const authModal = AuthModal("Authorization")
            const componentHeader = Header()
            
            componentHeader.init()
            authModal.init()


        }
    }
}

MainLayout().init()