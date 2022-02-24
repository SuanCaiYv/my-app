<template>
    <div class="tag">
        <button class="button" :class="{buttonClicked: isActive}" @click="clicked">{{props.name}}</button>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<string>("Tag")

const props = defineProps({
    id: String,
    name: String,
    clickFunc: Function
})

const isActive = ref<boolean>(false)

const clicked = function () {
    isActive.value = !isActive.value
    // @ts-ignore
    props.clickFunc(props.id, isActive.value)
}
</script>

<style scoped>
.tag {
    width: auto;
    height: 30px;
    margin-top: 10px;
    margin-left: 5px;
    margin-right: 5px;
    display: inline-block;
}

.button {
    width: 100%;
    height: 100%;
    border: none;
    box-sizing: border-box;
    border-radius: 16px;
    padding: 0 8px;
    line-height: 26px;
    font-size: 1rem;
    font-weight: bolder;
    display: inline-block;
    background-color: white;
}

.buttonClicked {
    background-color: rgba(0,0,0,0.2);
}
</style>