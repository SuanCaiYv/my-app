<template>
    <div>
        <button @click="send">测试</button>
        aaa
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import {Alert} from "../util/alert"

const name = ref<String>("Test")

let str = ref<String>("")

const send = function () {
    Alert({}, function (msg: String) {
        console.log(msg)
    })
}
// <input type="file" name="file" @change="upload0($event)">
// let file = ""
//
// const upload0 = function (e: Event) {
//     file = e.target.files[0]
// }
//
// const upload = function (e: Event) {
//     e.preventDefault()
//     let param = new FormData()
//     param.append("file", file)
//     httpClient.upload("/t/upload", {}, param, function (resp: Response) {
//         console.log(resp)
//     })
// }
</script>

<style scoped>

</style>