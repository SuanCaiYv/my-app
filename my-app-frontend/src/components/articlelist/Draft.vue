<template>
    <div class="draft">
        <button class="button">草稿管理</button>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("Draft")
</script>

<style scoped>
.draft {
    width: 120px;
    height: 100%;
    grid-area: draft;
    margin-right: 20px;
}

.button {
    width: 100%;
    height: 40px;
    margin-top: 10px;
    margin-bottom: 10px;
    padding: 0;
    border: none;
    border-radius: 20px;
    font-size: 1rem;
    font-weight: bolder;
    background-color: rgba(0,0,0,0.1);
}

.button:hover {
    background-color: rgba(0,0,0,0.15);
}

.button:active {
    background-color: rgba(0,0,0,0.2);
}
</style>