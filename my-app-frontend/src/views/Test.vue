<template>
    <div class="test">
        <div style="width: 100px; height: 100px">
            <img src="../../public/icons/100.svg" style="width: 100px; height: 100px">
        </div>
    </div>
</template>

<script setup lang="ts">
import {reactive, ref, watch} from "vue"

const name = ref<String>("Test")

const str = ref<string>('')
const arr = reactive<Array<string>>([])

watch(arr, () => {
    console.log(arr.length)
})

const update = function () {
    arr.push(str.value)
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