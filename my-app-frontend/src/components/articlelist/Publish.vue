<template>
    <div class="publish">
        <button class="button">发布文章</button>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("Publish")
</script>

<style scoped>
.publish {
    width: 120px;
    height: 100%;
    grid-area: publish;
    margin-right: 20px;
}

.button {
    width: 100%;
    height: 40px;
    margin-top: 10px;
    margin-bottom: 10px;
    padding: 0;
    border: 2px solid lightgray;
    border-radius: 20px;
    font-size: 1rem;
    font-weight: bolder;
    background-color: white;
}

.button:hover {
    background-color: lightgray;
}

.button:active {
    background-color: gainsboro;
}
</style>