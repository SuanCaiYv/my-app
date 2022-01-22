<template>
    <div v-show="visible">
        <div class="alert">
            <p class="alert-message">{{ msg }}</p>
            <button class="alert-button" @click="handleClose">确定</button>
        </div>
        <div class="alert-mask" @click.self="handleClose"></div>
    </div>
</template>

<script setup lang="ts">
import {defineProps, getCurrentInstance, onMounted, ref, watch, watchEffect} from "vue"

const name = ref<String>("Alert")

const props = defineProps({
    val: Boolean,
    el: Node,
    callback: Function
})

onMounted(() => {
    visible.value = props.val
})

let visible = ref<boolean>(false)
let msg = ref<string>("")

const setVisible = function (visible0: boolean) {
    visible.value = visible0
}

const handleClose = function () {
    console.log(props.el)
    // @ts-ignore
    document.getElementById("app").removeChild(props.el)
    // @ts-ignore
    props.callback("aaa")
    visible.value = false
}
</script>

<style scoped>
.alert {
    position: absolute;
    width: 320px;
    left: 50%;
    transform: translate(-50%, 0);
    top: 20%;
    background: #fff;
    border-radius: 4px;
    padding: 24px;
    z-index: 1001;
}

.alert-message {
    font-size: 14px;
    line-height: 22px;
    color: #333;
    margin-bottom: 32px;
}

.alert-button {
    min-width: 80px;
    padding: 8px 24px;
    text-align: center;
    background: #0075de;
    border: 0;
    outline: 0;
    float: right;
    color: #fff;
    border-radius: 4px;
    cursor: pointer;
}

.alert-mask {
    position: absolute;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    background-color: #000;
    opacity: 0.5;
    z-index: 1000;
}
</style>