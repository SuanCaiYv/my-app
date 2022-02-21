<template>
    <div class="confirm">
        <p class="confirm-message">{{ msg }}</p>
        <button class="cancel-button" @click="cancel">取消</button>
        <button class="confirm-button" @click="confirm">确定</button>
    </div>
    <div class="confirm-mask"></div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("Confirm")

const props = defineProps({
    divNode: Node,
    msg: String,
    cancelCallback: Function,
    confirmCallback: Function,
})

const cancel = function () {
    // @ts-ignore
    document.getElementById("app").removeChild(props.divNode)
    // @ts-ignore
    props.cancelCallback()
}

const confirm = function () {
    // @ts-ignore
    document.getElementById("app").removeChild(props.divNode)
    // @ts-ignore
    props.confirmCallback()
}
</script>

<style scoped>
.confirm {
    position: absolute;
    width: 320px;
    left: 50%;
    transform: translate(-50%, 0);
    top: 20%;
    background: white;
    border-radius: 24px;
    padding: 24px;
    font-weight: bolder;
    z-index: 1001;
}

.confirm-message {
    font-size: 14px;
    line-height: 22px;
    color: black;
    margin-bottom: 32px;
}

.cancel-button {
    min-width: 80px;
    padding: 8px 24px;
    font-weight: bolder;
    text-align: center;
    background: dodgerblue;
    border: 0;
    outline: 0;
    float: left;
    margin-left: 50px;
    color: white;
    border-radius: 18px;
    cursor: pointer;
}

.confirm-button {
    min-width: 80px;
    padding: 8px 24px;
    font-weight: bolder;
    text-align: center;
    background: dodgerblue;
    border: 0;
    outline: 0;
    float: right;
    margin-right: 50px;
    color: white;
    border-radius: 18px;
    cursor: pointer;
}

.confirm-mask {
    position: absolute;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    background-color: black;
    opacity: 0.5;
    z-index: 1000;
}
</style>