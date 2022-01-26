<template>
    <div class="search">
        <input type="text" class="key" v-model="keyStr" @keydown.enter.down="searchKey(keyStr)" placeholder="键入并回车以检索"/>
    </div>
</template>

<script setup lang="ts">
import {inject, ref} from "vue"

const name = ref<String>("Search")

let keyStr = ref<String>("")

const searchKey = inject("searchKey")
</script>

<style scoped>
.search {
    width: 100%;
    height: 100%;
}

.key {
    width: 200px;
    height: 100%;
    /*border: 1px solid black;*/
    /*box-sizing: border-box;*/
    border: none;
    padding: 0;
    border-radius: 16px;
    display: inline-block;
    font-size: 1rem;
    background-color: whitesmoke;
}
</style>