<template>
    <div class="search">
        <input type="text" class="key" v-model="keyStr"/>
        <button class="find" @click="searchKey(keyStr)">搜索</button>
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
    width: 140px;
    height: 100%;
    box-sizing: border-box;
    border: 1px solid black;
    border-radius: 8px;
    background-color: white;
}

.find {
    width: 60px;
    height: 100%;
    box-sizing: border-box;
    border: 1px solid black;
    border-radius: 8px;
    background-color: white;
}
</style>