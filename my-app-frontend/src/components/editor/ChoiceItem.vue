<template>
    <div class="choice-item">
        <div class="value">{{props.value}}</div>
        <button class="delete" @click="props.deleteFunc(id)">✖️</button>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("ChoiceItem")

const props = defineProps({
    value: String,
    id: String,
    deleteFunc: Function,
})
</script>

<style scoped>
.choice-item {
    min-width: 1px;
    height: 30px;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    border-radius: 16px;
    margin-left: 10px;
    margin-top: 7px;
    margin-bottom: 7px;
    display: inline-block;
    background-color: darkgrey;
}

.value {
    min-width: 1px;
    height: 30px;
    /*border: 1px solid silver;*/
    /*box-sizing: border-box;*/
    padding: 0 0 0 8px;
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