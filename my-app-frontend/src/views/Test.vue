<template>
    <div>
        <button class="test" @click="send">测试</button>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import alertFunc from "../util/alert";
import {confirmFunc} from "../util/confirm";
import {useRouter} from "vue-router";

const name = ref<String>("Test")

const router = useRouter()

const send = function () {
    router.push("/home")
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
.test:hover {
    background-color: #9dbbff;
}
</style>