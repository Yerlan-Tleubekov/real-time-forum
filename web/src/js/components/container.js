import { createElementObj } from "../utils/create"
import { tagsOptions } from "../utils/options"

const options = {
    container: {
        tagName: 'div',
        attrs: [["class", "container-fluid"]]
    
    }
}

const Container = (...children) => {
    console.log('container children', children)
    const container = createElementObj({...options.container, children: [...children]})

    return {
        get: () => {
            return container
        }
    }
}

export default Container