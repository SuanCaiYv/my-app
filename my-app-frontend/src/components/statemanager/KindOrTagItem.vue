<template>
    <div class="kind-or-tag">
        <div class="value">啊嗷嗷</div>
        <button class="delete">X</button>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("KindOrTagItem")
</script>

<style scoped>
.kind-or-tag {
    min-width: 1px;
    height: 30px;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    margin-left: 10px;
    margin-top: 10px;
    display: inline-block;
    background-color: skyblue;
}

.value {
    min-width: 1px;
    height: 30px;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    display: inline-block;
    text-align: left;
    line-height: 30px;
    font-size: 1.4rem;
    vertical-align: bottom;
}

.delete {
    width: 30px;
    height: 30px;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    border: none;
    padding: 0;
    display: inline-block;
    font-size: 1.4rem;
    vertical-align: bottom;
    background-color: white;
}
</style>