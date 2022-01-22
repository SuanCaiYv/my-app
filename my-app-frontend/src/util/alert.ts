import {createApp} from 'vue'
import AlertComponent from "../components/Alert.vue"

const Alert = function (options: object, callback0: Function) {
    let divElement = document.createElement("div");
    const instance = createApp(AlertComponent, {
        val: true,
        el: divElement,
        callback: callback0
    })
    instance.mount(divElement)
    // @ts-ignore
    document.getElementById("app").appendChild(divElement)
}

export {Alert}