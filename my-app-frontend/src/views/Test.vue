<template>
    <div class="test">
        <Publish></Publish>
    </div>
</template>

<script setup lang="ts">
import {ref, watch} from "vue"
import {useRouter} from "vue-router";
import {marked} from "marked";
import Publish from "../components/editor/Publish.vue"

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
.test {
    width: 100%;
    height: 100%;
}

.a {
    width: 100%;
    height: 100%;
    grid-area: a;
}

.b {
    width: 100%;
    height: 100%;
    grid-area: b;
    grid-template-areas: "left right";
    grid-template-columns: 1fr 1fr;
    grid-auto-rows: 10px;
}

.left {
    grid-area: left;
    background: red;
    width: 200px;
    overflow: scroll;
    word-break: break-all;
    float: left;
    height: 100%;
}

.right {
    grid-area: right;
    background: green;
    margin-left: 200px;
    height: 100%;
    word-break: break-all;
    overflow: scroll;
}
</style>