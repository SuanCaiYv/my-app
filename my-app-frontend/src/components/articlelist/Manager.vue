<template>
    <div class="manager">
        <button class="button" @click="router.push('/state_manager')">其他管理</button>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import {useRouter} from "vue-router";

const name = ref<string>("Manager")
const router = useRouter()
</script>

<style scoped>
.manager {
    width: 120px;
    height: 100%;
    grid-area: manager;
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