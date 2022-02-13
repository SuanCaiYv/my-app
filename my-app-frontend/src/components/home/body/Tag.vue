<template>
    <div class="tag">
        <button class="button" :class="{buttonClicked: isActive}" @click="isActive = !isActive">{{value}}</button>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("Tag")

const props = defineProps({
    value: String
})

let isActive = ref<Boolean>(false)
</script>

<style scoped>
.tag {
    width: auto;
    height: 40px;
    margin-top: 10px;
    margin-left: 5px;
    margin-right: 5px;
    display: inline-block;
}

.button {
    width: 100%;
    height: 100%;
    border: 2px solid ;
    border-radius: 16px;
    line-height: 40px;
    font-size: 1.2rem;
    display: inline-block;
    background-color: #99dbff;
}

.buttonClicked {
    background-color: aqua;
}
</style>