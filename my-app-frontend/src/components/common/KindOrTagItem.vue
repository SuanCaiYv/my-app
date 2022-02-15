<template>
    <div class="kind-or-tag">
        <div class="value">{{ name }}</div>
        <button class="delete" @click="del">✖️</button>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<string>("KindOrTagItem")

const props = defineProps({
    id: String,
    name: String,
    deleteFunc: Function,
})

const del = function () {
    // @ts-ignore
    props.deleteFunc(props.id)
}
</script>

<style scoped>
.kind-or-tag {
    min-width: 1px;
    height: 30px;
    border-radius: 16px;
    margin-left: 10px;
    margin-top: 10px;
    display: inline-block;
    background-color: darkgrey;
}

.value {
    min-width: 1px;
    height: 30px;
    display: inline-block;
    text-align: left;
    line-height: 30px;
    font-size: 1.4rem;
    vertical-align: bottom;
}

.delete {
    width: 30px;
    height: 30px;
    border: none;
    border-radius: 16px;
    padding: 0;
    display: inline-block;
    font-size: 1.4rem;
    vertical-align: bottom;
    background-color: darkgrey;
}

.delete:hover {
    background-color: lightgrey;
}

.delete:active {
    background-color: snow;
}
</style>