import './css/normalize.css'
import { createElementObj } from "./js/utils/create"
import { tagsOptions } from './js/utils/options'
import Header from './js/components/header'



function MainLayout() {
    const layout = document.getElementById("body")
    const header = createElementObj(tagsOptions.header)
    const main = createElementObj(tagsOptions.main)
    const footer = createElementObj(tagsOptions.footer)

    return {
        init: () => {
            layout.append(header, main, footer)
        
            const componentHeader = Header()
            componentHeader.init()
        }
    }
}

MainLayout().init()