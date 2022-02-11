<template>
    <div class="search">
        <input type="text" class="key" v-model="keyStr" @keydown.enter.down="searchKey" placeholder="键入并回车以检索"/>
    </div>
</template>

<script setup lang="ts">
import {inject, ref} from "vue"

const name = ref<String>("Search")

let keyStr = ref<String>("")

const searchKey = function () {
    console.log(keyStr.value)
}
</script>

<style scoped>
.search {
    width: 100%;
    height: 100%;
    background-color: white;
}

.key {
    width: 192px;
    height: 40px;
    margin-top: 10px;
    margin-bottom: 10px;
    border: none;
    border-radius: 16px;
    display: inline-block;
    font-size: 1rem;
    font-weight: bold;
    background-color: whitesmoke;
    padding: 0 0 0 8px;
    outline: none;
}
</style>