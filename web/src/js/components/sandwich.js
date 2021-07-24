import { createElementObj } from "../utils/create"


const Sandwich = () => {
    const span = createElementObj({
        tagName: 'span',
        classNames: 'navbar-toggler-icon',
    })

    const sandwich = createElementObj({
        tagName: 'button',
        classNames: 'navbar-toggler',
        attrs: [
            ['type', 'button'],
            ['data-bs-toggle', 'collapse'],
            ['data-bs-target', '#navbarSupportedContent'],
            ['aria-controls', "navbarSupportedContent"],
            ['aria-expanded', "false"],
            ['aria-label', "Toggle navigation"]],
        children: [span],
    })

    return {
        init: () => {

        },
        get: () => {
            return sandwich
        }

    }
}

export default Sandwich;