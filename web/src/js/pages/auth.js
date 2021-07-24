import { createElementObj } from "../utils/create";
import { Modal } from "bootstrap";

function AuthModal() {

    Modal
    const modal = createElementObj({
        tagName: "div",
        classNames: "auth-modal",
    })

    return modal
}


export default AuthModal
