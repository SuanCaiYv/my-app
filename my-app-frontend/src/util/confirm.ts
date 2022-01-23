import {createApp} from "vue";
import ConfirmComponent from "../components/Confirm.vue";

const confirmFunc = function (confirmMsg: string, cancel: Function, confirm: Function) {
    let divElement = document.createElement("div");
    const instance = createApp(ConfirmComponent, {
        divNode: divElement,
        msg: confirmMsg,
        cancelCallback: cancel,
        confirmCallback: confirm
    })
    instance.mount(divElement)
    // @ts-ignore
    document.getElementById("app").appendChild(divElement)
}

export {confirmFunc};