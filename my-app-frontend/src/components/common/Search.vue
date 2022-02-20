<template>
    <div class="search">
        <input type="text" class="key" v-model="keyStr" @keydown.enter.down="searchKey" placeholder="键入并回车以检索"/>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<string>("Search")

const props = defineProps({
    searchFunc: Function
})

const keyStr = ref<string>("")

const searchKey = function () {
    // @ts-ignore
    props.searchFunc(keyStr.value)
}
</script>

<style scoped>
.search {
    width: 100%;
    height: 100%;
    background-color: white;
}

.key {
    width: calc(100% - 10px);
    height: 40px;
    margin-top: 10px;
    padding: 0 0 0 10px;
    border: none;
    border-radius: 16px;
    display: inline-block;
    font-size: 1rem;
    font-weight: bold;
    background-color: rgba(0,0,0,0.05);
    outline: none;
}
</style>