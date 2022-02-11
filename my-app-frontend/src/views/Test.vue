<template>
    <div>
        <input type="text" v-model="str">
        <div>{{s}}</div>
    </div>
</template>

<script setup lang="ts">
import {ref, watch} from "vue"
import {useRouter} from "vue-router";
import {marked} from "marked";

const name = ref<String>("Test")

const router = useRouter()

const str = ref<string>('')
const s = ref<string>('')

watch(str, (n, o) => {
    s.value = marked.parse(n)
})

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